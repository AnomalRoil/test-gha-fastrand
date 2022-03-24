package main

import (
	"fmt"
	_ "unsafe"
)

//go:linkname FastRand runtime.fastrand
func FastRand() uint32

//go:linkname FastRandN runtime.fastrandn
func FastRandN(n uint32) uint32

func main() {
	fmt.Println("vim-go")
	fmt.Println("Testing 1,2,3,4 cases:")
	for i := uint32(1); i < 5; i++ {
		for j := 0; j < 20; j++ {
			fmt.Println("\t", i, "\t", FastRandN(i))
		}
	}
}
