package y2020

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type d8Op int

const (
	OpAcc d8Op = iota + 1
	OpJmp
	OpNop
)

type d8VM struct {
	Acc   int64
	Ip    int64
	Insts []d8Inst
}

func (vm *d8VM) Step() error {
	if vm.Ip >= int64(len(vm.Insts)) || vm.Ip < 0 {
		return errors.New("Instruction pointer out of reach")
	}

	inst := vm.Insts[vm.Ip]

	switch inst.Op {
	case OpAcc:
		vm.Acc += inst.Arg
		vm.Ip++
	case OpJmp:
		vm.Ip += inst.Arg
	case OpNop:
		vm.Ip++
	default:
		return errors.New("unknown instruction!")
	}

	return nil
}

type d8Inst struct {
	Op  d8Op
	Arg int64
}

func d8ParseProgram(input []byte) ([]d8Inst, error) {
	insts := make([]d8Inst, 0, 100)
	for _, line := range bytes.Split(bytes.TrimSpace(input), []byte{'\n'}) {
		s := string(bytes.TrimSpace(line))
		parts := strings.SplitN(s, " ", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid input!")
		}

		x, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return nil, err
		}

		op := d8Inst{Arg: x}

		switch parts[0] {
		case "jmp":
			op.Op = OpJmp
		case "acc":
			op.Op = OpAcc
		case "nop":
			op.Op = OpNop
		default:
			return nil, errors.New("Invalid input!")
		}

		insts = append(insts, op)
	}

	return insts, nil
}

func Day08Part01(input []byte) (string, error) {
	insts, err := d8ParseProgram(input)
	if err != nil {
		return "", err
	}
	vm := &d8VM{
		Insts: insts,
		Acc:   0,
		Ip:    0,
	}
	vm = vm
	// seen := make(map[int64]

	return "", errors.New("Not implemented")
}

func Day08Part02(input []byte) (string, error) {
	return "", errors.New("Not implemented")
}
