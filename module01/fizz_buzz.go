package module01

import (
	"fmt"
	"strings"
	"strconv"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	var sequence []string
	var next string

	if n > 1 {
		for i:=1; i <= n; i++ {
			switch {
			case i%3 == 0 && i % 5 == 0:
				next = "Fizz Buzz"
			case i%3 == 0:
				next = "Fizz"
			case i%5 == 0:
				next = "Buzz"
			default:
				next = strconv.Itoa(i)
			}

			sequence = append(sequence, next)
		}
		fmt.Println(strings.Join(sequence, ", "))
		return
	}

	fmt.Println("1")
}
