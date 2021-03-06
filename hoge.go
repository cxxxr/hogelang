package main

import (
	"bufio"
	"fmt"
	"github.com/cxxxr/hogelang/compiler"
	"github.com/cxxxr/hogelang/parser"
	"github.com/cxxxr/hogelang/vm"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func eval(m *vm.Machine, exe *vm.Executable, str string) (vm.Object, error) {
	p := parser.NewParser(strings.NewReader(str))
	x, err := p.Parse()
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		exe, err = compiler.Compile(exe, x)
		if err == nil {
			m.SetExecutable(exe)
			return vm.Run(m)
		} else {
			return nil, err
		}
	}
}

func main() {
	exe := vm.NewExecutable(nil)
	m := vm.NewMachine(exe)
	vm.Init(m)

	if len(os.Args) > 1 {
		for _, file := range os.Args[1:] {
			dat, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			_, err = eval(m, exe, string(dat))
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			line, _, err := reader.ReadLine()

			if err != nil {
				break
			}

			if len(line) == 0 {
				continue
			}

			v, err := eval(m, exe, string(line))
			if err != nil {
				fmt.Println(err)
			} else if v != nil {
				fmt.Println(v)
			}
		}
	}
}
