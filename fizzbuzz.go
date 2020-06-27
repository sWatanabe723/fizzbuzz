package main

import (
	"fmt"

	"./fizzbuzzFunc"
)

func main() {
	for i := 0; i < 101; i++ {
		fmt.Println(fizzbuzzFunc.FizzBuzz(i))
	}

}
