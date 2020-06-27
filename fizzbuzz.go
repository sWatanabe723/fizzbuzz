package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for i := 0; i < 101; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if strings.Contains(strconv.Itoa(i), "7") {
			fmt.Println("GitHub-fixed")
		} else {
			fmt.Println(i)
		}
	}

}
