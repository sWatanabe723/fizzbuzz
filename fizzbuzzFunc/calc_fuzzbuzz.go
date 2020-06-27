package fizzbuzzFunc

import (
	"strconv"
	"strings"
)

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
