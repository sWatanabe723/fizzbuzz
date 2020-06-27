package fizzbuzzFunc

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{17, "GitHub-fixed"},
		{27, "GitHub-fixed"},
		{75, "GitHub-fixed"},
		{77, "GitHub-fixed"},
		{1, "1"},
		{13, "13"},
		{12, "Fizz"},
		{99, "Fizz"},
		{5, "Buzz"},
		{80, "Buzz"},
		{15, "FizzBuzz"},
		{45, "FizzBuzz"},
	}

	for _, test := range tests {
		if got := FizzBuzz(test.input); got != test.want {
			t.Errorf("FizzBuzz(%v) = %s", test.input, got)
		}

	}
}
