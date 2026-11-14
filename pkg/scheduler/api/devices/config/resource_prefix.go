// Copyright (c) 2026 Hygon Information Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"
	"strings"
)

const (
	// ResourceNamespace is the vendor prefix for all K8s resource names
	// (hygon.com/<family>…), matching k8s-hcu-device-plugin.
	ResourceNamespace = "hygon.com"

	// DefaultResourcePrefix is the product-family prefix used in K8s resource
	// names and HAMi annotations (hcu → hygon.com/hcunum, …).
	DefaultResourcePrefix = "hcu"
)

// resourcePrefix is the lowercase product-family prefix (hcu|*cu|dcu|…).
// Mutated only by InitResourcePrefix at scheduler startup.
var resourcePrefix = DefaultResourcePrefix

// HygonVHCURegister is the node annotation that carries the virtual GPU
// information registered by the device plugin (hami.io/node-<prefix>-register).
// Rebuilt by InitResourcePrefix for *cu compatibility.
var HygonVHCURegister = "hami.io/node-hcu-register"

// HygonVHCUHandshake is the node annotation used for the handshake between the
// scheduler and the device plugin (hami.io/node-handshake-<prefix>).
// Rebuilt by InitResourcePrefix for *cu compatibility.
var HygonVHCUHandshake = "hami.io/node-handshake-hcu"

// InitResourcePrefix validates and applies deviceshare.ResourcePrefix.
// Same rule as k8s-hcu-device-plugin: the product-family token must contain
// "cu" (hcu, *cu, dcu, acu, bcu, ccu, …).
// Must be called once at startup, before InitDevicesConfig, so that the vHCU
// resource names and HAMi annotation keys are derived from the active prefix.
// Keep this value identical to k8s-hcu-device-plugin's RESOURCE_PREFIX.
func InitResourcePrefix(prefix string) error {
	prefix = strings.ToLower(strings.TrimSpace(prefix))
	if prefix == "" {
		prefix = DefaultResourcePrefix
	}
	if !strings.Contains(prefix, "cu") {
		return fmt.Errorf("invalid resource-prefix %q, want a product family containing \"cu\" (e.g. hcu, *cu, dcu)", prefix)
	}
	resourcePrefix = prefix
	HygonVHCURegister = fmt.Sprintf("hami.io/node-%s-register", prefix)
	HygonVHCUHandshake = fmt.Sprintf("hami.io/node-handshake-%s", prefix)
	return nil
}

// ResourcePrefix returns the lowercase product-family prefix (hcu|*cu|dcu|…).
func ResourcePrefix() string {
	return resourcePrefix
}

// ResourceNameNum returns hcunum / *cunum.
func ResourceNameNum() string {
	return resourcePrefix + "num"
}

// ResourceNameCores returns hcucores / *cucores.
func ResourceNameCores() string {
	return resourcePrefix + "cores"
}

// ResourceNameMem returns hcumem / *cumem.
func ResourceNameMem() string {
	return resourcePrefix + "mem"
}

// ResourceCountName returns hygon.com/hcunum (or hygon.com/*cunum).
func ResourceCountName() string {
	return ResourceNamespace + "/" + ResourceNameNum()
}

// ResourceMemoryName returns hygon.com/hcumem (or hygon.com/*cumem).
func ResourceMemoryName() string {
	return ResourceNamespace + "/" + ResourceNameMem()
}

// ResourceCoreName returns hygon.com/hcucores (or hygon.com/*cucores).
func ResourceCoreName() string {
	return ResourceNamespace + "/" + ResourceNameCores()
}

// applyResourcePrefix overrides the Hygon vHCU resource names with the names
// derived from the active resource prefix, keeping the scheduler consistent
// with the HCU device plugin.
func applyResourcePrefix(c *Config) {
	if c == nil {
		return
	}
	c.HygonConfig.ResourceCountName = ResourceCountName()
	c.HygonConfig.ResourceMemoryName = ResourceMemoryName()
	c.HygonConfig.ResourceCoreName = ResourceCoreName()
}
