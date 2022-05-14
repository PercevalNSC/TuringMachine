package main

import (
	"fmt"
	"turing/dtm"
)

func main() {
	w := "ab#"
	tm := dtm.TuringMachineSet('$')
	fmt.Println("turing machine output:", tm.Run(w))
}
