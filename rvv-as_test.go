package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAssembleInstruction(t *testing.T) {
	// Define test cases with expected opcodes
	tests := []struct {
		expected    uint32
		instruction string
	}{
		{0x26000057, "vand.vv v0, v0, v0"},
		{0x260000d7, "vand.vv v1, v0, v0"},
		{0x260001d7, "vand.vv v3, v0, v0"},
		{0x260003d7, "vand.vv v7, v0, v0"},
		{0x260007d7, "vand.vv v15, v0, v0"},
		{0x26000fd7, "vand.vv v31, v0, v0"},
		{0x26100057, "vand.vv v0, v1, v0"},
		{0x26300057, "vand.vv v0, v3, v0"},
		{0x26700057, "vand.vv v0, v7, v0"},
		{0x26f00057, "vand.vv v0, v15, v0"},
		{0x27f00057, "vand.vv v0, v31, v0"},
		{0x26008057, "vand.vv v0, v0, v1"},
		{0x26018057, "vand.vv v0, v0, v3"},
		{0x26038057, "vand.vv v0, v0, v7"},
		{0x26078057, "vand.vv v0, v0, v15"},
		{0x260f8057, "vand.vv v0, v0, v31"},
		//
		{0xa2003057, "vsrl.vi v0, v0, 0"},
		{0xa200b057, "vsrl.vi v0, v0, 1"},
		{0xa200b0d7, "vsrl.vi v1, v0, 1"},
		{0xa200b1d7, "vsrl.vi v3, v0, 1"},
		{0xa200b3d7, "vsrl.vi v7, v0, 1"},
		{0xa200b7d7, "vsrl.vi v15, v0, 1"},
		{0xa200bfd7, "vsrl.vi v31, v0, 1"},
		{0xa210b057, "vsrl.vi v0, v1, 1"},
		{0xa230b057, "vsrl.vi v0, v3, 1"},
		{0xa270b057, "vsrl.vi v0, v7, 1"},
		{0xa2f0b057, "vsrl.vi v0, v15, 1"},
		{0xa3f0b057, "vsrl.vi v0, v31, 1"},
		{0xa201b057, "vsrl.vi v0, v0, 3"},
		{0xa203b057, "vsrl.vi v0, v0, 7"},
		{0xa207b057, "vsrl.vi v0, v0, 15"},
		{0xa20fb057, "vsrl.vi v0, v0, 31"},
		//
		{0x32000057, "vrgather.vv v0, v0, v0"},
		{0x320000d7, "vrgather.vv v1, v0, v0"},
		{0x320001d7, "vrgather.vv v3, v0, v0"},
		{0x320003d7, "vrgather.vv v7, v0, v0"},
		{0x320007d7, "vrgather.vv v15, v0, v0"},
		{0x32000fd7, "vrgather.vv v31, v0, v0"},
		{0x32100057, "vrgather.vv v0, v1, v0"},
		{0x32300057, "vrgather.vv v0, v3, v0"},
		{0x32700057, "vrgather.vv v0, v7, v0"},
		{0x32f00057, "vrgather.vv v0, v15, v0"},
		{0x33f00057, "vrgather.vv v0, v31, v0"},
		{0x32008057, "vrgather.vv v0, v0, v1"},
		{0x32018057, "vrgather.vv v0, v0, v3"},
		{0x32038057, "vrgather.vv v0, v0, v7"},
		{0x32078057, "vrgather.vv v0, v0, v15"},
		{0x320f8057, "vrgather.vv v0, v0, v31"},
		//
		{0x2e000057, "vxor.vv v0, v0, v0"},
		{0x2e0000d7, "vxor.vv v1, v0, v0"},
		{0x2e0001d7, "vxor.vv v3, v0, v0"},
		{0x2e0003d7, "vxor.vv v7, v0, v0"},
		{0x2e0007d7, "vxor.vv v15, v0, v0"},
		{0x2e000fd7, "vxor.vv v31, v0, v0"},
		{0x2e100057, "vxor.vv v0, v1, v0"},
		{0x2e300057, "vxor.vv v0, v3, v0"},
		{0x2e700057, "vxor.vv v0, v7, v0"},
		{0x2ef00057, "vxor.vv v0, v15, v0"},
		{0x2ff00057, "vxor.vv v0, v31, v0"},
		{0x2e008057, "vxor.vv v0, v0, v1"},
		{0x2e018057, "vxor.vv v0, v0, v3"},
		{0x2e038057, "vxor.vv v0, v0, v7"},
		{0x2e078057, "vxor.vv v0, v0, v15"},
		{0x2e0f8057, "vxor.vv v0, v0, v31"},
	}

	// Iterate over test cases
	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing %s", test.instruction), func(t *testing.T) {
			// Assemble the instruction
			got, err := assembleInstruction(test.instruction)
			if err != nil {
				t.Errorf("Error assembling instruction %s: %v", test.instruction, err)
				return
			}

			// Check if the result matches the expected opcode
			if got != test.expected {
				fmt.Printf(" got: %32s\n", strconv.FormatUint(uint64(got), 2))
				fmt.Printf("want: %32s\n", strconv.FormatUint(uint64(test.expected), 2))
				t.Errorf("assembleInstruction(%s) = 0x%08x; want 0x%08x", test.instruction, got, test.expected)
			}
		})
	}
}
