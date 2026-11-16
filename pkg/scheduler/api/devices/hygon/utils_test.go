// Copyright (c) 2026 Hygon Information Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package vhcu

import (
	"reflect"
	"testing"
)

// TestDecodeContainerDevices covers the AI-review finding: annotation
// segments carrying fewer than 4 comma-separated fields must be skipped
// with a log instead of causing an index-out-of-range panic.
func TestDecodeContainerDevices(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    ContainerDevices
		wantLen int
	}{
		{
			name:    "empty string",
			input:   "",
			want:    ContainerDevices{},
			wantLen: 0,
		},
		{
			name:    "segment without comma is skipped",
			input:   "no-comma-segment",
			want:    ContainerDevices{},
			wantLen: 0,
		},
		{
			name:  "valid segment with 4 fields",
			input: "HCU-TPX1,HCU,2048,30",
			want: ContainerDevices{
				{UUID: "HCU-TPX1", Type: "HCU", Usedmem: 2048, Usedcores: 30},
			},
			wantLen: 1,
		},
		{
			name:    "segment with 3 fields is skipped without panic",
			input:   "HCU-TPX1,HCU,2048",
			want:    ContainerDevices{},
			wantLen: 0,
		},
		{
			name:    "segment with 2 fields is skipped without panic",
			input:   "HCU-TPX1,HCU",
			want:    ContainerDevices{},
			wantLen: 0,
		},
		{
			name:    "segment with 1 field is skipped without panic",
			input:   "HCU-TPX1",
			want:    ContainerDevices{},
			wantLen: 0,
		},
		{
			name:  "valid and short segments mixed: short one is skipped",
			input: "HCU-TPX1,HCU,2048,30:HCU-TPX2,HCU",
			want: ContainerDevices{
				{UUID: "HCU-TPX1", Type: "HCU", Usedmem: 2048, Usedcores: 30},
			},
			wantLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decodeContainerDevices(tt.input)
			if len(got) != tt.wantLen {
				t.Fatalf("decodeContainerDevices(%q) got %d devices, want %d: %+v", tt.input, len(got), tt.wantLen, got)
			}
			if tt.wantLen > 0 && !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("decodeContainerDevices(%q) got %+v, want %+v", tt.input, got, tt.want)
			}
		})
	}
}

// TestDecodePodDevices covers the container-level split: a malformed
// container segment must not abort parsing of the remaining containers.
func TestDecodePodDevices(t *testing.T) {
	t.Run("two valid containers", func(t *testing.T) {
		got := decodePodDevices("HCU-TPX1,HCU,2048,30;HCU-TPX2,HCU,1024,10")
		if len(got) != 2 || len(got[0]) != 1 || len(got[1]) != 1 {
			t.Fatalf("decodePodDevices got %+v", got)
		}
	})

	t.Run("second container malformed is skipped without panic", func(t *testing.T) {
		got := decodePodDevices("HCU-TPX1,HCU,2048,30;HCU-TPX2")
		if len(got) != 2 || len(got[1]) != 0 {
			t.Fatalf("decodePodDevices malformed got %+v", got)
		}
		if len(got[0]) != 1 || got[0][0].UUID != "HCU-TPX1" {
			t.Fatalf("decodePodDevices first container got %+v", got[0])
		}
	})

	t.Run("empty input", func(t *testing.T) {
		got := decodePodDevices("")
		if len(got) != 0 {
			t.Fatalf("decodePodDevices empty got %+v", got)
		}
	})
}
