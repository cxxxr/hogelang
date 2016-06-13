package main

import (
	"bufio"
	"fmt"
	"github.com/cxxxr/hogelang/compiler"
	"github.com/cxxxr/hogelang/parser"
	"github.com/cxxxr/hogelang/vm"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	exe := vm.NewExecutable(nil)
	m := vm.NewMachine(exe)
	vm.Init(m)

	for {
		fmt.Print("> ")
		line, _, err := reader.ReadLine()

		if err != nil {
			break
		}

		if len(line) == 0 {
			continue
		}

		p := parser.NewParser(strings.NewReader(string(line)))
		x, err := p.Parse()
		if err != nil {
			fmt.Println(err)
		} else {
			exe = compiler.Compile(exe, x)
			m.SetExecutable(exe)
			fmt.Println(vm.Run(m))
		}
	}
}
