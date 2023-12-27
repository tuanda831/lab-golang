package main

import (
	"fmt"
)

func Solution(s string) string {
	srune := []rune(s)

	if len(srune) > 0 {
		last := srune[0]
		index := 0
		for i := 1; i < len(srune); i++ {
			if last > srune[i] {
				return string(append(srune[0:index], srune[index+1:]...))
			}
			last = srune[i]
			index = i
		}
	}

	return string(srune[0 : len(s)-1])
}

func main() {
	fmt.Println("Hello Tuan")
}
