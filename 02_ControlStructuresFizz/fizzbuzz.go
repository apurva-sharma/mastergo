package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzbuzz(n int) {
	for cnt := 1; cnt <= n; cnt++ {
		switch {
		case cnt%15 == 0:
			fmt.Println("FizzBuzz")
		case cnt%3 == 0:
			fmt.Println("Fizz")
		case cnt%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(cnt)
		}
	}
}

func main() {
	n := 50
	if len(os.Args) > 1 {
		max, err := strconv.Atoi(os.Args[1])
		if err == nil {
			n = max
		}
	}
	fizzbuzz(n)
}
