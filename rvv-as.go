package main

import (
	"fmt"
	"strings"
)

// Function to assemble the instruction
func assembleInstruction(instruction string) (uint32, error) {
	parts := strings.Fields(instruction)
	switch parts[0] {
	case "vand.vv":
		return assembleVand(parts)
	case "vsrl.vi":
		return assembleVsrl(parts)
	case "vrgather.vv":
		return assembleVrgather(parts)
	case "vxor.vv":
		return assembleVxor(parts)
	default:
		panic(fmt.Sprintf("Unknown instruction: %v", parts))
	}
}

func main() {
	// Example usage
	instruction := "vand.vv v31, v0, v0"
	opcode, err := assembleInstruction(instruction)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print opcode as hexadecimal
	fmt.Printf("Opcode: 0x%08x\n", opcode)
}
