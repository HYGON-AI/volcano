// Copyright (c) 2026 Hygon Information Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package vhcu

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"

	"volcano.sh/volcano/pkg/scheduler/api/devices"
	deviceconfig "volcano.sh/volcano/pkg/scheduler/api/devices/config"
	"volcano.sh/volcano/pkg/scheduler/plugins/util/nodelock"
)

type HCUUsage struct {
	UsedMem  uint
	UsedCore uint
}

// HCUDevice include HCU id, memory and the pods that are sharing it.
type HCUDevice struct {
	// GPU ID
	ID int
	// Node this HCU Device belongs
	Node string
	// HCU Unique ID
	UUID string
	// The resource usage by pods that are sharing this HCU
	PodMap map[string]*HCUUsage
	// memory per card
	Memory uint
	// max sharing number
	Number uint
	// max core number
	Cores uint
	// type of this number
	Type string
	// Health condition of this HCU
	Health bool
	// number of allocated
	UsedNum uint
	// number of device memory allocated
	UsedMem uint
	// number of core used
	UsedCore uint
}

type HCUDevices struct {
	Name string
	// We cache score in filter step according to schedulePolicy, to avoid recalculating in score
	Score float64

	Device map[int]*HCUDevice
}

func NewHCUDevices(name string, node *v1.Node) *HCUDevices {
	if node == nil {
		return nil
	}
	annos, ok := node.Annotations[deviceconfig.HygonVHCURegister]
	if !ok {
		return nil
	}
	handshake, handshakeOK := node.Annotations[deviceconfig.HygonVHCUHandshake]
	nodedevices := decodeNodeDevices(name, annos)
	if (nodedevices == nil) || len(nodedevices.Device) == 0 {
		return nil
	}

	for _, val := range nodedevices.Device {
		klog.V(3).InfoS("Hygon Device registered name", "name", nodedevices.Name, "val", *val)
	}

	// We have to handshake here in order to avoid time-inconsistency between scheduler and nodes.
	// Older device plugin versions do not maintain the handshake annotation; in that case the
	// register annotation is used directly so HCU scheduling still works.
	if !handshakeOK {
		klog.V(4).Infof("node %s has no HCU handshake annotation, using register annotation directly", node.Name)
		return nodedevices
	}
	if strings.Contains(handshake, "Requesting") {
		parts := strings.Split(handshake, "_")
		if len(parts) > 1 {
			formertime, err := time.Parse("2006.01.02 15:04:05", parts[1])
			if err == nil && time.Now().After(formertime.Add(time.Second*60)) {
				// The device plugin did not answer the handshake in 60s (e.g. an
				// older plugin without handshake support). Keep the devices and
				// re-request instead of dropping them.
				klog.V(3).Infof("node %v device handshake %s not answered in 60s, keeping devices and re-requesting", node.Name, handshake)
				tmppat := make(map[string]string)
				tmppat[deviceconfig.HygonVHCUHandshake] = "Requesting_" + time.Now().Format("2006.01.02 15:04:05")
				patchNodeAnnotations(node, tmppat)
			}
		}
	} else if strings.Contains(handshake, "Deleted") {
		return nil
	} else {
		tmppat := make(map[string]string)
		tmppat[deviceconfig.HygonVHCUHandshake] = "Requesting_" + time.Now().Format("2006.01.02 15:04:05")
		patchNodeAnnotations(node, tmppat)
	}
	return nodedevices
}

func (ds *HCUDevices) ScoreNode(pod *v1.Pod, schedulePolicy string) float64 {
	if ds == nil {
		return 0
	}
	/* TODO: we need a base score to be campatable with preemption, it means a node without evicting a task has
	   a higher score than those needs to evict a task */

	// Use cached stored in filter state in order to avoid recalculating.
	return ds.Score
}

func (ds *HCUDevices) GetIgnoredDevices() []string {
	return []string{}
}

func (ds *HCUDevices) AddQueueResource(pod *v1.Pod) map[string]float64 {
	if ds == nil {
		return map[string]float64{}
	}
	klog.V(5).InfoS("AddQueueResource", "Name", pod.Name)
	res := map[string]float64{}
	ids, ok := pod.Annotations[AssignedIDsAllocatedAnnotations]
	if !ok {
		klog.Errorf("pod %s has no annotation %s", pod.Name, AssignedIDsAllocatedAnnotations)
		return res
	}
	podDev := decodePodDevices(ids)
	for _, val := range podDev {
		for _, deviceused := range val {
			for _, gsdevice := range ds.Device {
				if strings.Contains(deviceused.UUID, gsdevice.UUID) {
					res[getConfig().ResourceMemoryName] += float64(deviceused.Usedmem * 1000)
					res[getConfig().ResourceCoreName] += float64(deviceused.Usedcores * 1000)
				}
			}
		}
	}
	klog.V(4).InfoS("AddQueueResource", "Name=", pod.Name, "res=", res)
	return res
}

// AddResource adds the pod to GPU pool if it is assigned
func (ds *HCUDevices) AddResource(pod *v1.Pod) {
	if ds == nil {
		return
	}

	ds.addResource(pod.Annotations, pod)
}

func (ds *HCUDevices) addResource(annotations map[string]string, pod *v1.Pod) {
	if _, ok := annotations[AssignedIDsAllocatedAnnotations]; !ok {
		klog.Errorf("pod %s has no annotation %s", pod.Name, AssignedIDsAllocatedAnnotations)
		return
	}

	ds.addToPodMap(annotations, pod)
}

func (ds *HCUDevices) addToPodMap(annotations map[string]string, pod *v1.Pod) {
	ids, ok := annotations[AssignedIDsAllocatedAnnotations]
	if !ok {
		klog.Errorf("pod %s has no annotation %s", pod.Name, AssignedIDsAllocatedAnnotations)
		return
	}
	podDev := decodePodDevices(ids)
	for _, val := range podDev {
		for _, deviceused := range val {
			for idx, gsdevice := range ds.Device {
				if strings.Contains(deviceused.UUID, gsdevice.UUID) {
					podUID := string(pod.UID)
					_, ok := gsdevice.PodMap[podUID]
					if !ok {
						ds.Device[idx].PodMap[podUID] = &HCUUsage{
							UsedMem:  0,
							UsedCore: 0,
						}
					}
					ds.Device[idx].UsedMem += deviceused.Usedmem
					ds.Device[idx].UsedCore += deviceused.Usedcores
					ds.Device[idx].PodMap[podUID].UsedMem += deviceused.Usedmem
					ds.Device[idx].PodMap[podUID].UsedCore += deviceused.Usedcores
				}
			}
		}
	}
}

// SubResource is a no-op for the in-memory ledger: the vHCU usage ledger is
// rebuilt each session from pod annotations by AddResource, so releasing a
// pod only requires the session rebuild (same behavior as the original vHCU
// implementation, which only updated device metrics here).
func (ds *HCUDevices) SubResource(pod *v1.Pod) {
	if ds == nil {
		return
	}
}

// DeepCopy returns a deep copy of HCUDevices for use in dry-run simulation.
func (ds *HCUDevices) DeepCopy() interface{} {
	if ds == nil {
		return nil
	}
	cp := &HCUDevices{
		Name:   ds.Name,
		Score:  ds.Score,
		Device: make(map[int]*HCUDevice, len(ds.Device)),
	}
	for id, dev := range ds.Device {
		if dev == nil {
			continue
		}
		newDev := &HCUDevice{
			ID:       dev.ID,
			Node:     dev.Node,
			UUID:     dev.UUID,
			Memory:   dev.Memory,
			Number:   dev.Number,
			Cores:    dev.Cores,
			Type:     dev.Type,
			Health:   dev.Health,
			UsedNum:  dev.UsedNum,
			UsedMem:  dev.UsedMem,
			UsedCore: dev.UsedCore,
			PodMap:   make(map[string]*HCUUsage, len(dev.PodMap)),
		}
		for uid, usage := range dev.PodMap {
			u := *usage
			newDev.PodMap[uid] = &u
		}
		cp.Device[id] = newDev
	}
	return cp
}

func (ds *HCUDevices) HasDeviceRequest(pod *v1.Pod) bool {
	return PodRequestsHCU(pod)
}

// GetStatus is used for debug and monitor.
func (ds *HCUDevices) GetStatus() string {
	return ""
}

func (ds *HCUDevices) Release(kubeClient kubernetes.Interface, pod *v1.Pod) (*devices.DeviceReservation, error) {
	if ds == nil || pod == nil || pod.Annotations == nil {
		return nil, nil
	}
	// Release is required for rollback paths (e.g. UnPipeline) where NodeInfo
	// does not invoke subResource for Pipelined tasks.
	ds.SubResource(pod)

	if pod.Annotations[DeviceBindPhase] == "success" {
		return nil, nil
	}

	keys := []string{
		AssignedNodeAnnotations,
		AssignedIDsToAllocateAnnotations,
		AssignedIDsAllocatedAnnotations,
		AssignedTimeAnnotations,
		BindTimeAnnotations,
		DeviceBindPhase,
	}
	if err := devices.RemovePodAnnotations(kubeClient, pod, keys); err != nil {
		return nil, err
	}
	for _, k := range keys {
		delete(pod.Annotations, k)
	}
	return &devices.DeviceReservation{
		DeviceType:  DeviceName,
		Annotations: devices.AnnotationKeyMap(keys),
	}, nil
}

func (ds *HCUDevices) FilterNode(pod *v1.Pod, schedulePolicy string) (int, string, error) {
	if HygonVHCUEnable {
		klog.V(4).Infoln("hami-vhcu DeviceSharing starts filtering pods", pod.Name)
		fit, _, score, err := checkNodeHCUSharingPredicateAndScore(pod, ds, true, schedulePolicy)
		klog.V(5).Infof("schedulePolicy:%s , FilterNode-ds: %+v , score: %+v", schedulePolicy, ds, score)

		if err != nil || !fit {
			klog.ErrorS(err, "Failed to fitler node to vhcu task", "pod", pod.Name)
			return devices.Unschedulable, "hami-vhcu DeviceSharing error", err
		}
		ds.Score = score
		klog.V(4).Infoln("hami-vhcu DeviceSharing successfully filters pods")
	}
	return devices.Success, "", nil
}

func (ds *HCUDevices) Allocate(kubeClient kubernetes.Interface, pod *v1.Pod) (*devices.DeviceReservation, error) {
	if HygonVHCUEnable {
		// Idempotency: the task pod may already carry the allocation
		// annotations from a previous cycle or an earlier bind attempt.
		if _, ok := pod.Annotations[AssignedIDsAllocatedAnnotations]; ok {
			klog.V(4).Infof("Pod %s already allocated (in-memory), skip", pod.Name)
			return nil, nil
		}
		updatedPod, err := kubeClient.CoreV1().Pods(pod.Namespace).Get(context.Background(), pod.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		_, ok := updatedPod.Annotations[AssignedIDsAllocatedAnnotations]
		if ok {
			klog.V(4).Infof("Pod %s already allocated, skip", pod.Name)
			return nil, nil
		}
		klog.V(4).Infoln("hami-vhcu DeviceSharing:Into AllocateToPod", pod.Name)
		fit, allocatedDevs, score, err := checkNodeHCUSharingPredicateAndScore(pod, ds, false, SchedulePolicyArgument)
		klog.V(5).Infof("schedulePolicy:%s , Allocate-ds: %+v , devices: %+v , score: %+v", SchedulePolicyArgument, ds, allocatedDevs, score)
		if err != nil || !fit {
			klog.ErrorS(err, "Failed to allocate vhcu task", "pod", pod.Name)
			return nil, err
		}
		lockValue := nodelock.GenerateNodeLockKeyByPod(pod)
		if NodeLockEnable {
			klog.InfoS("Node lock enabled", "node", ds.Name)
			nodelock.UseClient(kubeClient)
			err = nodelock.LockHCUNode(ds.Name, DeviceName, lockValue)
			if err != nil {
				return nil, errors.Errorf("node %s locked for %s hami-vhcu lockname %s", ds.Name, pod.Name, err.Error())
			}
		}

		annotations := make(map[string]string)
		annotations[AssignedNodeAnnotations] = ds.Name
		annotations[AssignedTimeAnnotations] = strconv.FormatInt(time.Now().Unix(), 10)
		podDevicesAnnotations := encodePodDevices(allocatedDevs)
		annotations[AssignedIDsToAllocateAnnotations] = podDevicesAnnotations
		annotations[AssignedIDsAllocatedAnnotations] = podDevicesAnnotations

		annotations[DeviceBindPhase] = "allocating"
		annotations[BindTimeAnnotations] = strconv.FormatInt(time.Now().Unix(), 10)
		// To avoid that the pod allocated info updating latency, add it first
		ds.addToPodMap(annotations, pod)
		if NodeLockEnable {
			nodelock.ReleaseHCUNodeLock(ds.Name, DeviceName, lockValue)
		}
		klog.V(3).Infoln("DeviceSharing:Allocate Success")
		// The annotations are carried to the bind request via the reservation
		// contract (predicates merges them into TaskInfo.PodAnnotations).
		return &devices.DeviceReservation{
			DeviceType:  DeviceName,
			Annotations: annotations,
		}, nil
	}
	return nil, nil
}

func (ds *HCUDevices) TryAddPod(device *HCUDevice, reqmem uint, reqcores uint) (bool, string) {

	if device.UsedNum+1 <= device.Number && device.UsedCore+reqcores <= device.Cores && device.UsedMem+reqmem <= device.Memory {
		device.UsedNum++
		device.UsedMem += reqmem
		device.UsedCore += reqcores
		return true, device.UUID
	}
	return false, device.UUID
}
