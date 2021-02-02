package main

import (
	"fmt"
	"os"
	"testing"
)

const COUNTER = 10000000

// RealizationWithPreBytes реализация, где я не создаю Fizz и Buzz каждый раз, а только в начале
func BenchmarkFizzbuzzPre(b *testing.B) {
	FIZZ := []byte("Fizz\n")
	BUZZ := []byte("Buzz\n")
	FIZZBUZZ := []byte("FizzBuzz\n")

	// Открываем файл
	file, err := os.Create("fizzbuzz_pre")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Наивная реализация
	for i := 1; i <= COUNTER; i++ {
		if i%15 == 0 {
			file.Write(FIZZBUZZ)
		} else if i%3 == 0 {
			file.Write(FIZZ)
		} else if i%5 == 0 {
			file.Write(BUZZ)
		} else {
			file.WriteString(fmt.Sprintf("%d\n", i))
		}
	}
}
