// Copyright (c) 2026 Hygon Information Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package vhcu

const (
	// DeviceName used to indicate this device
	DeviceName = "hami.io/mutex.lock"

	HCUInUseType                     = "hygon.com/use-hcutype"
	HCUNoUseType                     = "hygon.com/nouse-hcutype"
	HCUInUseUUID                     = "hygon.com/use-gpuuuid"
	HCUNoUseUUID                     = "hygon.com/nouse-gpuuuid"
	AssignedTimeAnnotations          = "hami.io/vgpu-time"
	AssignedIDsToAllocateAnnotations = "hami.io/hcu-devices-to-allocate"
	AssignedIDsAllocatedAnnotations  = "hami.io/hcu-devices-allocated"
	AssignedNodeAnnotations          = "hami.io/vgpu-node"
	BindTimeAnnotations              = "hami.io/bind-time"
	DeviceBindPhase                  = "hami.io/bind-phase"

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
)

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
