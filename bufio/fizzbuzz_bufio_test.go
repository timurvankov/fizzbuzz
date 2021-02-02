package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

const COUNTER = 100000000

// RealizationWithBufio using bufio
func BenchmarkFizzbuzzBufio(b *testing.B) {
	FIZZ := []byte("Fizz\n")
	BUZZ := []byte("Buzz\n")
	FIZZBUZZ := []byte("FizzBuzz\n")

	// Открываем файл
	file, err := os.Create("fizzbuzz_bufio")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	for i := 1; i <= COUNTER; i++ {
		if i%15 == 0 {
			w.Write(FIZZBUZZ)
		} else if i%3 == 0 {
			w.Write(FIZZ)
		} else if i%5 == 0 {
			w.Write(BUZZ)
		} else {
			w.WriteString(fmt.Sprintf("%d\n", i))
		}
	}

	w.Flush()
}
