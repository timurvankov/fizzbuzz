package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"testing"
)

const COUNTER = 100000000

var wg sync.WaitGroup

var FIZZ = []byte("Fizz\n")
var BUZZ = []byte("Buzz\n")
var FIZZBUZZ = []byte("FizzBuzz\n")

func countOffset(start int) int {
	end := start
	start = 1

	countFIZZBUZZ := end / 15
	countFIZZ := end/3 - countFIZZBUZZ
	countBUZZ := end/5 - countFIZZBUZZ
	countNUMBERS := end - countFIZZ - countBUZZ - countFIZZBUZZ - 1

	return countFIZZBUZZ*len(FIZZBUZZ) + countFIZZ*len(FIZZ) + countBUZZ*len(BUZZ) + countNUMBERS*len(FIZZBUZZ)
}

func fizzBuzz(start int, end int) {
	file, err := os.Create("fizzbuzz_gouroutine")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	file.Seek(int64(countOffset(start)), 0)
	w := bufio.NewWriter(file)
	w.WriteString("\n")

	for i := start; i <= end; i++ {
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
	wg.Done()
}

// RealizationWithBufio using bufio
func BenchmarkFizzbuzzBufio(b *testing.B) {
	routines := 2
	runtime.GOMAXPROCS(routines)
	remove := []int{}
	for i := 0; i != routines; i++ {
		wg.Add(1)
		remove = append(remove, i*(COUNTER/routines)+1)
		go fizzBuzz(1+i*(COUNTER/routines), i*(COUNTER/routines)+COUNTER/routines)
	}
	wg.Wait()
	/*for _, rem := range remove {
		println(rem)
		removeLines("fizzbuzz_gouroutine", rem, 1)
	}*/
}

func removeLines(fn string, start, n int) (err error) {
	if start < 1 {
		return errors.New("invalid request.  line numbers start at 1.")
	}
	if n < 0 {
		return errors.New("invalid request.  negative number to remove.")
	}
	var f *os.File
	if f, err = os.OpenFile(fn, os.O_RDWR, 0); err != nil {
		return
	}
	defer func() {
		if cErr := f.Close(); err == nil {
			err = cErr
		}
	}()
	var b []byte
	if b, err = ioutil.ReadAll(f); err != nil {
		return
	}
	cut, ok := skip(b, start-1)
	if !ok {
		return fmt.Errorf("less than %d lines", start)
	}
	if n == 0 {
		return nil
	}
	tail, ok := skip(cut, n)
	if !ok {
		return fmt.Errorf("less than %d lines after line %d", n, start)
	}
	t := int64(len(b) - len(cut))
	if err = f.Truncate(t); err != nil {
		return
	}
	if len(tail) > 0 {
		_, err = f.WriteAt(tail, t)
	}
	return
}

func skip(b []byte, n int) ([]byte, bool) {
	for ; n > 0; n-- {
		if len(b) == 0 {
			return nil, false
		}
		x := bytes.IndexByte(b, '\n')
		if x < 0 {
			x = len(b)
		} else {
			x++
		}
		b = b[x:]
	}
	return b, true
}
