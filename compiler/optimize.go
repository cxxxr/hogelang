package compiler

import (
	"github.com/cxxxr/hogelang/vm"
	//"fmt"
)

type optimizer func(*vm.Executable, []vm.Instr, int) (int, bool)

var optTable = []optimizer{
	func(exe *vm.Executable, code []vm.Instr, i int) (int, bool) {
		if !(i < len(code)-2) {
			return 0, false
		}
		if !(code[i].Opcode == vm.CONST && code[i+1].Opcode == vm.CONST) {
			return 0, false
		}
		a, ok1 := exe.GetConst(code[i].Arg1).(vm.Fixnum)
		b, ok2 := exe.GetConst(code[i+1].Arg1).(vm.Fixnum)

		if !(ok1 && ok2) {
			return 0, false
		}

		switch code[i+2].Opcode {
		case vm.ADD:
			a += b
		case vm.SUB:
			a -= b
		case vm.MUL:
			a *= b
		case vm.DIV:
			a /= b
		case vm.MOD:
			a %= b
		default:
			return 0, false
		}
		exe.Gen(code[i].Pos, vm.CONST, exe.SetConst(a), 0)
		return i + 2, true
	},
}

func findOptimizer(exe *vm.Executable, code []vm.Instr, i int) (int, bool) {
	for _, f := range optTable {
		if j, ok := f(exe, code, i); ok {
			return j, ok
		}
	}
	return 0, false
}

func optimize(exe *vm.Executable) *vm.Executable {
	var flag bool = false
	exe2 := vm.NewExecutable(exe)
	code := exe.GetCode()
	for i := 0; i < len(code); i++ {
		if j, ok := findOptimizer(exe2, code, i); ok {
			i = j
			flag = true
		} else {
			exe2.GenCopy(code[i])
		}
	}

	if flag {
		return optimize(exe2)
	} else {
		return exe2
	}
}
