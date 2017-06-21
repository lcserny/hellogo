package main

import "fmt"

func main() {
	r := returnLocalVariableAddress()
	fmt.Printf("%d", r.i)
}

type Record struct {
	i int
}

func returnLocalVariableAddress() *Record {
	return &Record{3}
}