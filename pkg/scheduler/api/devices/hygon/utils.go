// Copyright (c) 2026 Hygon Information Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package vhcu

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8stypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"

	"volcano.sh/volcano/pkg/scheduler/api/devices"
	"volcano.sh/volcano/pkg/scheduler/api/devices/config"
)

func decodeNodeDevices(name, str string) *HCUDevices {
	if !strings.Contains(str, ":") {
		return nil
	}
	tmp := strings.Split(str, ":")
	retval := &HCUDevices{
		Name:   name,
		Device: make(map[int]*HCUDevice),
		Score:  float64(0),
	}
	for index, val := range tmp {
		if strings.Contains(val, ",") {
			items := strings.Split(val, ",")
			if len(items) < 7 {
				klog.Error("wrong Node HCU info: ", val)
				return nil
			}

			count, _ := strconv.Atoi(items[1])
			devmem, _ := strconv.Atoi(items[2])
			devcores, _ := strconv.Atoi(items[3])
			health, _ := strconv.ParseBool(items[6])
			i := HCUDevice{
				ID:     index,
				Node:   name,
				UUID:   items[0],
				Number: uint(count),
				Cores:  uint(devcores),
				Memory: uint(devmem),
				Type:   items[4],
				PodMap: make(map[string]*HCUUsage),
				Health: health,
			}

			retval.Device[index] = &i
		}
	}
	return retval
}

func encodeContainerDevices(cd []ContainerDevice) string {
	tmp := ""
	for _, val := range cd {
		tmp += val.UUID + "," + val.Type + "," + strconv.Itoa(int(val.Usedmem)) + "," + strconv.Itoa(int(val.Usedcores)) + ":"
	}
	klog.V(4).Infoln("Encoded container Devices=", tmp)
	return tmp
	//return strings.Join(cd, ",")
}

func encodePodDevices(pd []ContainerDevices) string {
	var ss []string
	for _, cd := range pd {
		ss = append(ss, encodeContainerDevices(cd))
	}
	return strings.Join(ss, ";")
}

func decodeContainerDevices(str string) ContainerDevices {
	if len(str) == 0 {
		return ContainerDevices{}
	}
	cd := strings.Split(str, ":")
	contdev := ContainerDevices{}
	tmpdev := ContainerDevice{}
	if len(str) == 0 {
		return contdev
	}
	for _, val := range cd {
		if strings.Contains(val, ",") {
			tmpstr := strings.Split(val, ",")
			tmpdev.UUID = tmpstr[0]
			tmpdev.Type = tmpstr[1]
			devmem, _ := strconv.ParseInt(tmpstr[2], 10, 32)
			tmpdev.Usedmem = uint(devmem)
			devcores, _ := strconv.ParseInt(tmpstr[3], 10, 32)
			tmpdev.Usedcores = uint(devcores)
			contdev = append(contdev, tmpdev)
		}
	}
	return contdev
}

func decodePodDevices(str string) []ContainerDevices {
	if len(str) == 0 {
		return []ContainerDevices{}
	}
	var pd []ContainerDevices
	for _, s := range strings.Split(str, ";") {
		cd := decodeContainerDevices(s)
		pd = append(pd, cd)
	}
	return pd
}

func checkVHCUResourcesInPod(pod *v1.Pod) bool {
	for _, container := range pod.Spec.Containers {
		_, ok := container.Resources.Limits[v1.ResourceName(getConfig().ResourceMemoryName)]
		if ok {
			return true
		}
		_, ok = container.Resources.Limits[v1.ResourceName(getConfig().ResourceCountName)]
		if ok {
			return true
		}
	}
	return false
}

func resourcereqs(pod *v1.Pod) []devices.ContainerDeviceRequest {
	countName := getConfig().ResourceCountName
	memoryName := getConfig().ResourceMemoryName
	percentageName := getConfig().ResourceMemoryPercentageName
	coreName := getConfig().ResourceCoreName
	return devices.ExtractResourceRequest(pod, "HCU", countName, memoryName, percentageName, coreName)
}

func checkHCUtype(annos map[string]string, cardtype string) bool {
	inuse, ok := annos[HCUInUseType]
	if ok {
		if !strings.Contains(inuse, ",") {
			if strings.Contains(strings.ToUpper(cardtype), strings.ToUpper(inuse)) {
				return true
			}
		} else {
			for _, val := range strings.Split(inuse, ",") {
				if strings.Contains(strings.ToUpper(cardtype), strings.ToUpper(val)) {
					return true
				}
			}
		}
		return false
	}
	nouse, ok := annos[HCUNoUseType]
	if ok {
		if !strings.Contains(nouse, ",") {
			if strings.Contains(strings.ToUpper(cardtype), strings.ToUpper(nouse)) {
				return false
			}
		} else {
			for _, val := range strings.Split(nouse, ",") {
				if strings.Contains(strings.ToUpper(cardtype), strings.ToUpper(val)) {
					return false
				}
			}
		}
		return true
	}
	return true
}

func checkHCUUUID(annos map[string]string, cardUUID string) bool {
	inuse, ok := annos[HCUInUseUUID]
	if ok {
		if !strings.Contains(inuse, ",") {
			if strings.Contains(strings.ToUpper(cardUUID), strings.ToUpper(inuse)) {
				return true
			}
		} else {
			for _, val := range strings.Split(inuse, ",") {
				if strings.Contains(strings.ToUpper(cardUUID), strings.ToUpper(val)) {
					return true
				}
			}
		}
		return false
	}
	nouse, ok := annos[HCUNoUseUUID]
	if ok {
		if !strings.Contains(nouse, ",") {
			if strings.Contains(strings.ToUpper(cardUUID), strings.ToUpper(nouse)) {
				return false
			}
		} else {
			for _, val := range strings.Split(nouse, ",") {
				if strings.Contains(strings.ToUpper(cardUUID), strings.ToUpper(val)) {
					return false
				}
			}
		}
		return true
	}
	return true
}

func checkType(annos map[string]string, d HCUDevice, n devices.ContainerDeviceRequest) bool {
	//General type check, NVIDIA->NVIDIA MLU->MLU
	if !strings.Contains(d.Type, n.Type) {
		return false
	}
	if n.Type == HygonHCUDevice {
		return checkHCUtype(annos, d.Type)
	}
	klog.Errorf("Unrecognized device %v", n.Type)
	return false
}

// getGPUDeviceSnapShot is not a strict deep copy, the pointer item is same with origin.
func getHCUDeviceSnapShot(snap *HCUDevices) *HCUDevices {
	ret := HCUDevices{
		Name:   snap.Name,
		Device: make(map[int]*HCUDevice),
		Score:  float64(0),
	}
	for index, val := range snap.Device {
		if val != nil {
			ret.Device[index] = &HCUDevice{
				ID:       val.ID,
				Node:     val.Node,
				UUID:     val.UUID,
				PodMap:   val.PodMap,
				Memory:   val.Memory,
				Number:   val.Number,
				Type:     val.Type,
				Health:   val.Health,
				UsedNum:  val.UsedNum,
				UsedMem:  val.UsedMem,
				UsedCore: val.UsedCore,
				Cores:    val.Cores,
			}
			klog.V(4).Infoln("getHCUDeviceSnapShot:", ret.Device[index].UsedMem, val.UsedMem, ret.Device[index].UsedCore, val.UsedCore)
		}
	}
	return &ret
}

// checkNodeHCUSharingPredicate checks if a pod with vhcu requirement can be scheduled on a node.
func checkNodeHCUSharingPredicateAndScore(pod *v1.Pod, dssnap *HCUDevices, replicate bool, schedulePolicy string) (bool, []ContainerDevices, float64, error) {
	// no gpu sharing request
	score := float64(0)
	if !checkVHCUResourcesInPod(pod) {
		return true, []ContainerDevices{}, 0, nil
	}

	ctrReq := resourcereqs(pod)
	if len(ctrReq) == 0 {
		return true, []ContainerDevices{}, 0, nil
	}

	var ds *HCUDevices
	if replicate {
		ds = getHCUDeviceSnapShot(dssnap)
	} else {
		ds = dssnap
	}

	// Select best card for exclusive request
	topologyInfo, err := GetTopologyInfo(dssnap.Name)
	if err != nil && topologyInfo == nil {
		klog.V(3).Infof("No HCU topology configmap found %v", err)
	}

	ctrdevs := []ContainerDevices{}
	for _, request := range ctrReq {
		devs := []ContainerDevice{}
		var allocatedIdx, preAllocatedIdx []int
		var filterTopologyInfo *Topology

		if int(request.Nums) > len(ds.Device) {
			return false, []ContainerDevices{}, 0, fmt.Errorf("no enough hcu cards on node %s", ds.Name)
		}
		klog.V(3).InfoS("Allocating device for container", "request", request)

		if _, ok := pod.Annotations[HCUInUseUUID]; ok {
			goto loop
		}

		// Deal with exclusive request
		if request.MemPercentagereq == 100 && request.Coresreq == 100 && topologyInfo != nil {
			for i := len(ds.Device) - 1; i >= 0; i-- {
				if ds.Device[i].UsedCore > 0 || ds.Device[i].UsedMem > 0 {
					allocatedIdx = append(allocatedIdx, ds.Device[i].ID)
				}
			}
			filterTopologyInfo = FilterTopologyByAllocated(topologyInfo, allocatedIdx)
			preAllocatedIdx = SelectBestDevices(filterTopologyInfo, int(request.Nums))
			if len(preAllocatedIdx) == 0 {
				return false, []ContainerDevices{}, 0, fmt.Errorf("no enough hcu cards on node %s for topology select", ds.Name)
			}

			for _, i := range preAllocatedIdx {
				if !checkType(pod.Annotations, *ds.Device[i], request) {
					klog.Errorln("failed checktype", ds.Device[i].Type, request.Type)
					return false, []ContainerDevices{}, 0, fmt.Errorf("failed checktype %s vs %s", ds.Device[i].Type, request.Type)
				}

				_, uuid := ds.TryAddPod(ds.Device[i], uint(float64(ds.Device[i].Memory)*float64(request.MemPercentagereq)/100.0), uint(request.Coresreq))
				klog.V(3).Info("fitted uuid: ", uuid)
				devs = append(devs, ContainerDevice{
					UUID:      uuid,
					Type:      request.Type,
					Usedmem:   uint(float64(ds.Device[i].Memory) * float64(request.MemPercentagereq) / 100.0),
					Usedcores: uint(request.Coresreq),
				})
				score += HCUScore(schedulePolicy, ds.Device[i])
			}
			ctrdevs = append(ctrdevs, devs)
			continue
		}

	loop:
		var deviceIndices []int
		for i := 0; i < len(ds.Device); i++ {
			deviceIndices = append(deviceIndices, i)
		}
		if schedulePolicy == binpackPolicy {
			sort.Slice(deviceIndices, func(i, j int) bool {
				return float64(ds.Device[deviceIndices[i]].UsedNum)/float64(ds.Device[deviceIndices[i]].Number)+
					float64(ds.Device[deviceIndices[i]].UsedCore)/float64(ds.Device[deviceIndices[i]].Cores)+
					float64(ds.Device[deviceIndices[i]].UsedMem)/float64(ds.Device[deviceIndices[i]].Memory) >
					float64(ds.Device[deviceIndices[j]].UsedNum)/float64(ds.Device[deviceIndices[j]].Number)+
						float64(ds.Device[deviceIndices[j]].UsedMem)/float64(ds.Device[deviceIndices[j]].Memory)+
						float64(ds.Device[deviceIndices[j]].UsedCore)/float64(ds.Device[deviceIndices[j]].Cores)
			})
		}
		for _, i := range deviceIndices {
			klog.V(3).InfoS("Scoring pod request", "memReq", request.Memreq, "memPercentageReq", request.MemPercentagereq, "coresReq", request.Coresreq, "Nums", request.Nums, "Index", i, "ID", ds.Device[i].ID)
			klog.V(3).InfoS("Current Device", "Index", i, "TotalMemory", ds.Device[i].Memory, "UsedMemory", ds.Device[i].UsedMem, "UsedCores", ds.Device[i].UsedCore, "replicate", replicate)
			if ds.Device[i].Number <= uint(ds.Device[i].UsedNum) {
				continue
			}
			memreqForCard := uint(0)
			// if we have mempercentage request, we ignore the mem request for every cards
			if request.MemPercentagereq != 101 {
				memreqForCard = uint(float64(ds.Device[i].Memory) * float64(request.MemPercentagereq) / 100.0)
			} else {
				memreqForCard = uint(request.Memreq * int32(getConfig().DeviceMemoryScaling))
			}
			if int(ds.Device[i].Memory)-int(ds.Device[i].UsedMem) < int(memreqForCard) {
				continue
			}
			if ds.Device[i].UsedCore+uint(request.Coresreq) > 100 {
				continue
			}
			// Coresreq=100 indicates it want this card exclusively
			if request.Coresreq == 100 && ds.Device[i].UsedNum > 0 {
				continue
			}
			// You can't allocate core=0 job to an already full GPU
			if ds.Device[i].UsedCore == 100 && request.Coresreq == 0 {
				continue
			}
			if !checkType(pod.Annotations, *ds.Device[i], request) {
				klog.Errorln("failed checktype", ds.Device[i].Type, request.Type)
				continue
			}
			if !checkHCUUUID(pod.Annotations, ds.Device[i].UUID) {
				klog.Errorln("failed check UUID", ds.Device[i].UUID)
				continue
			}

			fit, uuid := ds.TryAddPod(ds.Device[i], memreqForCard, uint(request.Coresreq))

			if !fit {
				klog.V(3).Info(ds.Device[i].ID, " not fit")
				continue
			}
			//total += gs.Devices[i].Count
			//free += node.Devices[i].Count - node.Devices[i].Used
			if request.Nums > 0 {
				request.Nums--
				if !replicate {
					klog.V(3).Info("fitted uuid: ", uuid)
				}
				devs = append(devs, ContainerDevice{
					UUID:      uuid,
					Type:      request.Type,
					Usedmem:   memreqForCard,
					Usedcores: uint(request.Coresreq),
				})
				if replicate {
					score = score + HCUScore(schedulePolicy, ds.Device[i])
				}
			}
			if request.Nums == 0 {
				break
			}
		}
		if request.Nums > 0 {
			return false, []ContainerDevices{}, 0, fmt.Errorf("not enough hcu fitted on this node")
		}
		ctrdevs = append(ctrdevs, devs)
	}
	return true, ctrdevs, score, nil
}

func HCUScore(schedulePolicy string, device *HCUDevice) float64 {
	var score float64
	numScore := float32(device.UsedNum) / float32(device.Number)
	coreScore := float32(device.UsedCore) / float32(device.Cores)
	memScore := float32(device.UsedMem) / float32(device.Memory)
	switch schedulePolicy {
	case binpackPolicy:
		score = float64(binpackMultiplier * (numScore + coreScore + memScore))
	case spreadPolicy:
		if device.UsedNum == 1 {
			score = spreadMultiplier
		}
	default:
		score = float64(0)
	}
	return score
}

func patchPodAnnotations(kubeClient kubernetes.Interface, pod *v1.Pod, annotations map[string]string) error {
	type patchMetadata struct {
		Annotations map[string]string `json:"annotations,omitempty"`
	}
	type patchPod struct {
		Metadata patchMetadata `json:"metadata"`
		//Spec     patchSpec     `json:"spec,omitempty"`
	}

	p := patchPod{}
	p.Metadata.Annotations = annotations

	bytes, err := json.Marshal(p)
	if err != nil {
		return err
	}
	_, err = kubeClient.CoreV1().Pods(pod.Namespace).
		Patch(context.Background(), pod.Name, k8stypes.StrategicMergePatchType, bytes, metav1.PatchOptions{})
	if err != nil {
		klog.Errorf("patch pod %v failed, %v", pod.Name, err)
	}

	return err
}

func patchNodeAnnotations(node *v1.Node, annotations map[string]string) error {
	type patchMetadata struct {
		Annotations map[string]string `json:"annotations,omitempty"`
	}
	type patchNode struct {
		Metadata patchMetadata `json:"metadata"`
		//Spec     patchSpec     `json:"spec,omitempty"`
	}

	p := patchNode{}
	p.Metadata.Annotations = annotations

	bytes, err := json.Marshal(p)
	if err != nil {
		return err
	}
	_, err = devices.GetClient().CoreV1().Nodes().
		Patch(context.Background(), node.Name, k8stypes.StrategicMergePatchType, bytes, metav1.PatchOptions{})
	if err != nil {
		klog.Errorf("patch node %v failed, %v", node.Name, err)
	}
	return err
}

func getConfig() config.HygonConfig {
	if config.GetConfig() != nil {
		return config.GetConfig().HygonConfig
	}
	return config.GetDefaultDevicesConfig().HygonConfig
}

func ParseTopologyInfoConfigMap(cm *v1.ConfigMap) (map[string]*Topology, error) {
	result := make(map[string]*Topology)

	for nodeName, topologyJSON := range cm.Data {
		var topo Topology
		if err := json.Unmarshal([]byte(topologyJSON), &topo); err != nil {
			return nil, fmt.Errorf("node %s parse failed: %w", nodeName, err)
		}
		result[nodeName] = &topo
	}
	return result, nil
}

func GetTopologyInfo(nodeName string) (*Topology, error) {
	cm, err := devices.GetClient().CoreV1().ConfigMaps("kube-system").
		Get(context.Background(), "hcu-topology-info", metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	topos, err := ParseTopologyInfoConfigMap(cm)
	if err != nil {
		return nil, err
	}

	if topo, ok := topos[nodeName]; ok {
		return topo, nil
	}
	return nil, nil
}

func FilterTopologyByAllocated(topo *Topology, allocatedIDs []int) *Topology {
	if topo == nil || len(topo.Matrix) == 0 {
		return topo
	}

	// Build the allocated set
	alloc := make(map[int]struct{}, len(allocatedIDs))
	for _, id := range allocatedIDs {
		alloc[id] = struct{}{}
	}

	// Find the row indices of the matrix that need to be retained.
	keepIdx := make([]int, 0, len(topo.Matrix))
	for i := 0; i < len(topo.Matrix); i++ {
		devID := topo.Matrix[i][i].SrcDvInd
		if _, ok := alloc[devID]; !ok {
			keepIdx = append(keepIdx, i)
		}
	}

	// Construct a new matrix
	newMatrix := make([][]Link, len(keepIdx))
	for i, oldI := range keepIdx {
		newMatrix[i] = make([]Link, len(keepIdx))
		for j, oldJ := range keepIdx {
			newMatrix[i][j] = topo.Matrix[oldI][oldJ]
		}
	}

	return &Topology{
		Matrix: newMatrix,
	}
}

func SelectBestDevices(topo *Topology, k int) []int {
	n := len(topo.Matrix)
	if k <= 0 || k > n {
		return nil
	}

	bestScore := -1 << 30
	var bestIdx []int

	var dfs func(start int, chosen []int, score int)
	dfs = func(start int, chosen []int, score int) {
		if len(chosen) == k {
			if score > bestScore {
				bestScore = score
				bestIdx = append([]int{}, chosen...)
			}
			return
		}

		if n-start < k-len(chosen) {
			return
		}

		for i := start; i < n; i++ {
			add := 0
			for _, j := range chosen {
				add += topo.Matrix[j][i].Weight
			}
			dfs(i+1, append(chosen, i), score+add)
		}
	}

	dfs(0, nil, 0)

	// Convert to the real device index (SrcDvInd)
	result := make([]int, 0, k)
	for _, idx := range bestIdx {
		result = append(result, topo.Matrix[idx][idx].SrcDvInd)
	}
	return result
}
