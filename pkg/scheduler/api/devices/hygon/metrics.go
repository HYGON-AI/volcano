// Copyright (c) 2026 Hygon Information Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package vhcu

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto" // auto-registry collectors in default registry
)

const (
	// VolcanoSubSystemName - subsystem name in prometheus used by volcano
	VolcanoSubSystemName = "volcano"

	// OnSessionOpen label
	OnSessionOpen = "OnSessionOpen"

	// OnSessionClose label
	OnSessionClose = "OnSessionClose"
)

var (
	VHCUDevicesSharedNumber = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: VolcanoSubSystemName,
			Name:      "vhcu_device_shared_number",
			Help:      "The number of VHCU tasks sharing this card",
		},
		[]string{"devID", "NodeName"},
	)
	VHCUDevicesAllocatedMemory = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: VolcanoSubSystemName,
			Name:      "vhcu_device_allocated_memory",
			Help:      "The number of VHCU memory allocated in this card",
		},
		[]string{"devID", "NodeName"},
	)
	VHCUDevicesAllocatedCores = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: VolcanoSubSystemName,
			Name:      "vhcu_device_allocated_cores",
			Help:      "The percentage of gpu compute cores allocated in this card",
		},
		[]string{"devID", "NodeName"},
	)
	VHCUDevicesMemoryTotal = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: VolcanoSubSystemName,
			Name:      "vhcu_device_memory_limit",
			Help:      "The number of total device memory in this card",
		},
		[]string{"devID", "NodeName"},
	)
	VHCUPodMemoryAllocated = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: VolcanoSubSystemName,
			Name:      "vhcu_device_memory_allocation_for_a_certain_pod",
			Help:      "The VHCU device memory allocated for a certain pod",
		},
		[]string{"devID", "NodeName", "podName"},
	)
	VHCUPodCoreAllocated = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: VolcanoSubSystemName,
			Name:      "vhcu_device_core_allocation_for_a_certain_pod",
			Help:      "The VHCU device core allocated for a certain pod",
		},
		[]string{"devID", "NodeName", "podName"},
	)
)

func (ds *HCUDevices) GetStatus() string {
	return ""
}

func ResetDeviceMetrics(UUID string, nodeName string, memory float64) {
	VHCUDevicesMemoryTotal.WithLabelValues(UUID, nodeName).Set(memory)
	VHCUDevicesSharedNumber.WithLabelValues(UUID, nodeName).Set(0)
	VHCUDevicesAllocatedCores.WithLabelValues(UUID, nodeName).Set(0)
	VHCUDevicesAllocatedMemory.WithLabelValues(UUID, nodeName).Set(0)

	VHCUPodMemoryAllocated.DeletePartialMatch(prometheus.Labels{"devID": UUID})
	VHCUPodCoreAllocated.DeletePartialMatch(prometheus.Labels{"devID": UUID})
}
func (ds *HCUDevices) AddPodMetrics(index int, podUID, podName string) {
	UUID := ds.Device[index].UUID
	NodeName := ds.Device[index].Node
	usage := ds.Device[index].PodMap[podUID]
	if usage == nil {
		VHCUDevicesSharedNumber.WithLabelValues(UUID, NodeName).Set(float64(ds.Device[index].UsedNum))
		VHCUDevicesAllocatedCores.WithLabelValues(UUID, NodeName).Set(float64(ds.Device[index].UsedCore))
		VHCUDevicesAllocatedMemory.WithLabelValues(UUID, NodeName).Set(float64(ds.Device[index].UsedMem))
		return
	}
	VHCUPodMemoryAllocated.WithLabelValues(UUID, NodeName, podName).Set(float64(usage.UsedMem))
	VHCUPodCoreAllocated.WithLabelValues(UUID, NodeName, podName).Set(float64(usage.UsedCore))
	VHCUDevicesSharedNumber.WithLabelValues(UUID, NodeName).Inc()
	VHCUDevicesAllocatedCores.WithLabelValues(UUID, NodeName).Set(float64(ds.Device[index].UsedCore))
	VHCUDevicesAllocatedMemory.WithLabelValues(UUID, NodeName).Set(float64(ds.Device[index].UsedMem))
}

func (ds *HCUDevices) SubPodMetrics(index int, podUID, podName string) {
	UUID := ds.Device[index].UUID
	NodeName := ds.Device[index].Node
	usage := ds.Device[index].PodMap[podUID]
	if usage == nil {
		VHCUPodMemoryAllocated.DeleteLabelValues(UUID, NodeName, podName)
		VHCUPodCoreAllocated.DeleteLabelValues(UUID, NodeName, podName)
		VHCUDevicesSharedNumber.WithLabelValues(UUID, NodeName).Set(float64(ds.Device[index].UsedNum))
		VHCUDevicesAllocatedCores.WithLabelValues(UUID, NodeName).Set(float64(ds.Device[index].UsedCore))
		VHCUDevicesAllocatedMemory.WithLabelValues(UUID, NodeName).Set(float64(ds.Device[index].UsedMem))
		return
	}
	VHCUPodMemoryAllocated.WithLabelValues(UUID, NodeName, podName).Set(float64(usage.UsedMem))
	VHCUPodCoreAllocated.WithLabelValues(UUID, NodeName, podName).Set(float64(usage.UsedCore))
	if usage.UsedMem == 0 {
		delete(ds.Device[index].PodMap, podUID)
		VHCUPodMemoryAllocated.DeleteLabelValues(UUID, NodeName, podName)
		VHCUPodCoreAllocated.DeleteLabelValues(UUID, NodeName, podName)
	}
	VHCUDevicesSharedNumber.WithLabelValues(UUID, NodeName).Dec()
	VHCUDevicesAllocatedCores.WithLabelValues(UUID, NodeName).Set(float64(ds.Device[index].UsedCore))
	VHCUDevicesAllocatedMemory.WithLabelValues(UUID, NodeName).Set(float64(ds.Device[index].UsedMem))
}
