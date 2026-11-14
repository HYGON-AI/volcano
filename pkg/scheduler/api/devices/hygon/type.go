// Copyright (c) 2026 Hygon Information Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package vhcu

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
)

const (
	// DeviceName used to indicate this device
	DeviceName = "hami.io/mutex.lock"

	HCUInUseType            = "hygon.com/use-hcutype"
	HCUNoUseType            = "hygon.com/nouse-hcutype"
	HCUInUseUUID            = "hygon.com/use-gpuuuid"
	HCUNoUseUUID            = "hygon.com/nouse-gpuuuid"
	AssignedTimeAnnotations = "hami.io/vgpu-time"
	AssignedNodeAnnotations = "hami.io/vgpu-node"
	BindTimeAnnotations     = "hami.io/bind-time"
	DeviceBindPhase         = "hami.io/bind-phase"

	HygonHCUDevice = "HCU"

	// binpack means the lower device memory remained after this allocation, the better
	binpackPolicy = "binpack"
	// spread means better put this task into an idle GPU card than a shared GPU card
	spreadPolicy = "spread"

	binpackMultiplier = 100
	spreadMultiplier  = 100
)

var (
	HygonVHCUEnable        bool
	NodeLockEnable         bool
	SchedulePolicyArgument string

	// AssignedIDsToAllocateAnnotations is the pod annotation written by the
	// scheduler to tell the device plugin which devices the pod is assigned
	// (hami.io/<prefix>-devices-to-allocate).
	// Rebuilt by ApplyResourcePrefix for *cu compatibility.
	AssignedIDsToAllocateAnnotations = "hami.io/hcu-devices-to-allocate"

	// AssignedIDsAllocatedAnnotations is the pod annotation that records the
	// devices allocated to the pod (hami.io/<prefix>-devices-allocated).
	// Rebuilt by ApplyResourcePrefix for *cu compatibility.
	AssignedIDsAllocatedAnnotations = "hami.io/hcu-devices-allocated"
)

// PodRequestsHCU reports whether the pod requests any Hygon vHCU resource
// according to the active resource prefix. Nil-safe by design (package-level).
func PodRequestsHCU(pod *v1.Pod) bool {
	return HygonVHCUEnable && checkVHCUResourcesInPod(pod)
}

// ApplyResourcePrefix rebuilds the HAMi pod annotation keys for the active
// product-family prefix (hcu|*cu|dcu|…), keeping them consistent with the HCU
// device plugin. Must be called once at startup, right after config.InitResourcePrefix.
func ApplyResourcePrefix(prefix string) {
	AssignedIDsToAllocateAnnotations = fmt.Sprintf("hami.io/%s-devices-to-allocate", prefix)
	AssignedIDsAllocatedAnnotations = fmt.Sprintf("hami.io/%s-devices-allocated", prefix)
}

type ContainerDevice struct {
	UUID string
	// device type, like NVIDIA, MLU
	Type      string
	Usedmem   uint
	Usedcores uint
}

type ContainerDevices []ContainerDevice

type Link struct {
	SrcDvInd    int    `json:"SrcDvInd"`
	DstDvInd    int    `json:"DstDvInd"`
	RemoteBdfID int    `json:"RemoteBdfID"`
	LinkType    string `json:"LinkType"`
	Weight      int    `json:"Weight"`
	Hops        int    `json:"Hops"`
}

type Topology struct {
	DeviceCount int      `json:"DeviceCount"`
	Matrix      [][]Link `json:"Matrix"`
}
