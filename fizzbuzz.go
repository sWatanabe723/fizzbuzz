package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for i := 0; i < 101; i++ {
		fmt.Println(FizzBuzz(i))
	}

}

func FizzBuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else if strings.Contains(strconv.Itoa(i), "7") {
		return "GitHub-fixed"
	} else {
		return strconv.Itoa(i)
	}
}
