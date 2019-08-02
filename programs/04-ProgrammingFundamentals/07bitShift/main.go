package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("kb in decimal: %d\t\tkb in binary: %b\n", kb, kb)
	fmt.Printf("mb in decimal: %d\t\tmb in binary: %b\n", mb, mb)
	fmt.Printf("gb in decimal: %d\tgb in binary: %b\n", gb, gb)
}
