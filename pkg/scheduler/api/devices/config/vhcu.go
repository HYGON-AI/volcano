// Copyright (c) 2026 Hygon Information Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package config

const (
	// HygonVHCUMemory extended gpu memory
	HygonVHCUMemory = "hygon.com/hcumem"
	// HygonVHCUCores indicates utilization percentage of VHCU
	HygonVHCUCores = "hygon.com/hcucores"
	// HygonVHCUNumber virtual GPU card number
	HygonVHCUNumber = "hygon.com/hcunum"
	// HygonVHCURegister virtual gpu information registered from device-plugin to scheduler
	HygonVHCURegister = "hami.io/node-hcu-register"
	// HygonVHCUHandshake for VHCU
	HygonVHCUHandshake = "hami.io/node-handshake-hcu"
)

// HygonConfig is used for Hygon-VHCU
type HygonConfig struct {
	// ResourceCountName is the name of GPU count
	ResourceCountName string `yaml:"resourceCountName"`
	// ResourceMemoryName is the name of GPU device memory
	ResourceMemoryName string `yaml:"resourceMemoryName"`
	// ResourceCoreName is the name of GPU core
	ResourceCoreName string `yaml:"resourceCoreName"`
	// ResourceMemoryPercentageName is the name of GPU device memory
	ResourceMemoryPercentageName string `yaml:"resourceMemoryPercentageName"`
	// ResourcePriority is the name of GPU priority
	ResourcePriority string `yaml:"resourcePriorityName"`
	// OverwriteEnv is whether we overwrite 'NVIDIA_VISIBLE_DEVICES' to 'none' for non-gpu tasks
	OverwriteEnv bool `yaml:"overwriteEnv"`
	// DefaultMemory is the number of device memory if not specified
	DefaultMemory uint `yaml:"defaultMemory"`
	// DefaultCores is the number of device cores if not specified
	DefaultCores uint `yaml:"defaultCores"`
	// DefaultGPUNum is the number of device number if not specified
	DefaultHCUNum uint `yaml:"defaultGPUNum"`
	// DeviceSplitCount is the number of fake-devices reported by VHCU-device-plugin per GPU
	DeviceSplitCount uint `yaml:"deviceSplitCount"`
	// DeviceMemoryScaling is the device memory oversubscription factor
	DeviceMemoryScaling float64 `yaml:"deviceMemoryScaling"`
	// DeviceCoreScaling is the device core oversubscription factor
	DeviceCoreScaling float64 `yaml:"deviceCoreScaling"`
}
