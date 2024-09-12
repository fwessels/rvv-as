package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Function to convert register name (like "v0") to its numeric value
func registerToNumber(register string) (int, error) {
	// Remove any trailing commas
	register = strings.TrimSuffix(register, ",")

	if !strings.HasPrefix(register, "v") {
		return 0, fmt.Errorf("invalid register: %s", register)
	}

	// Drop the 'v' prefix and convert the remaining string to an integer
	numberStr := register[1:] // Remove the first character 'v'
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, fmt.Errorf("invalid register number: %s", register)
	}

	// Check if the register number is within the valid range (0 to 31)
	if number < 0 || number > 31 {
		return 0, fmt.Errorf("register out of range: %d", number)
	}

	return number, nil
}

// Function to convert immediate value to 5-bit zimm5 field
func immediateToZimm5(immediate string) (int, error) {
	value, err := strconv.Atoi(immediate)
	if err != nil {
		return 0, fmt.Errorf("invalid immediate value: %s", immediate)
	}

	// Check if the immediate value is within the valid range (0 to 31)
	if value < 0 || value > 31 {
		return 0, fmt.Errorf("immediate value out of range: %d", value)
	}

	return value, nil
}

// Function to assemble the instruction
func assembleVand(parts []string) (uint32, error) {
	if len(parts) != 4 || parts[0] != "vand.vv" {
		return 0, fmt.Errorf("invalid instruction format")
	}

	// Convert register names to numeric values
	rd, err := registerToNumber(parts[1])
	if err != nil {
		return 0, err
	}

	rs1, err := registerToNumber(parts[2])
	if err != nil {
		return 0, err
	}

	rs2, err := registerToNumber(parts[3])
	if err != nil {
		return 0, err
	}

	// Construct the 32-bit instruction
	funct6 := uint32(0x26000000)
	vm := uint32(0) << 25
	rs1Val := uint32(rs1) << 20
	rs2Val := uint32(rs2) << 15
	rdVal := uint32(rd) << 7
	opcode := uint32(0x57)

	// Combine the instruction parts
	instructionWord := funct6 | vm | rs2Val | rs1Val | rdVal | opcode

	return instructionWord, nil
}

// Function to assemble the vsrl.vi instruction
func assembleVsrl(parts []string) (uint32, error) {
	if len(parts) != 4 || parts[0] != "vsrl.vi" {
		return 0, fmt.Errorf("invalid instruction format")
	}

	// Convert register names to numeric values
	rd, err := registerToNumber(parts[1])
	if err != nil {
		return 0, err
	}

	rs1, err := registerToNumber(parts[2])
	if err != nil {
		return 0, err
	}

	// Convert immediate value to zimm5
	zimm5, err := immediateToZimm5(parts[3])
	if err != nil {
		return 0, err
	}

	// Construct the 32-bit instruction
	funct6 := uint32(0xa2003000)
	rs1Val := uint32(rs1) << 20
	zimm5Val := uint32(zimm5) << 15
	rdVal := uint32(rd) << 7
	opcode := uint32(0x57)

	// Combine the instruction parts
	instructionWord := funct6 | zimm5Val | rs1Val | rdVal | opcode

	return instructionWord, nil
}

// Function to assemble the vxor.vv instruction
func assembleVxor(parts []string) (uint32, error) {
	if len(parts) != 4 || parts[0] != "vxor.vv" {
		return 0, fmt.Errorf("invalid instruction format")
	}

	// Convert register names to numeric values
	rd, err := registerToNumber(parts[1])
	if err != nil {
		return 0, err
	}

	rs1, err := registerToNumber(parts[2])
	if err != nil {
		return 0, err
	}

	rs2, err := registerToNumber(parts[3])
	if err != nil {
		return 0, err
	}

	// Construct the 32-bit instruction
	funct6 := uint32(0x2e000000)
	vm := uint32(0) << 25
	rs1Val := uint32(rs1) << 20
	rs2Val := uint32(rs2) << 15
	rdVal := uint32(rd) << 7
	opcode := uint32(0x57)

	// Combine the instruction parts
	instructionWord := funct6 | vm | rs2Val | rs1Val | rdVal | opcode

	return instructionWord, nil
}

// Function to assemble the vrgather.vv instruction
func assembleVrgather(parts []string) (uint32, error) {
	if len(parts) != 4 || parts[0] != "vrgather.vv" {
		return 0, fmt.Errorf("invalid instruction format")
	}

	// Convert register names to numeric values
	rd, err := registerToNumber(parts[1])
	if err != nil {
		return 0, err
	}

	rs1, err := registerToNumber(parts[2])
	if err != nil {
		return 0, err
	}

	rs2, err := registerToNumber(parts[3])
	if err != nil {
		return 0, err
	}

	// Construct the 32-bit instruction
	funct6 := uint32(0x32000000)
	vm := uint32(0) << 25
	rs1Val := uint32(rs1) << 20
	rs2Val := uint32(rs2) << 15
	rdVal := uint32(rd) << 7
	opcode := uint32(0x57)

	// Combine the instruction parts
	instructionWord := funct6 | vm | rs2Val | rs1Val | rdVal | opcode

	return instructionWord, nil
}
