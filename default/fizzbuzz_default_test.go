package main

import (
	"fmt"
	"os"
	"testing"
)

const COUNTER = 10000000

// DefaultRealization fizzbuzz function
func BenchmarkFizzbuzzDefault(b *testing.B) {
	// Открываем файл
	file, err := os.Create("fizzbuzz_default")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Наивная реализация
	for i := 1; i <= COUNTER; i++ {
		if i%15 == 0 {
			file.WriteString("FizzBuzz\n")
		} else if i%3 == 0 {
			file.WriteString("Fizz\n")
		} else if i%5 == 0 {
			file.WriteString("Buzz\n")
		} else {
			file.WriteString(fmt.Sprintf("%d\n", i))
		}
	}
}
