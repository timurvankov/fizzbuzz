package main

import (
	"crypto/md5"
	"io"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCheckCheckSum test
func TestCheckCheckSum(t *testing.T) {
	checkSum1 := getCheckSum("default/fizzbuzz_default")
	checkSum2 := getCheckSum("pre/fizzbuzz_pre")
	checkSum3 := getCheckSum("bufio/fizzbuzz_bufio")
	checkSum4 := getCheckSum("goroutine/fizzbuzz_gouroutine")

	for i := range checkSum1 {
		//assert.Equal(t, checkSum1[i], checkSum2[i])
		//assert.Equal(t, checkSum1[i], checkSum3[i])
		assert.Equal(t, checkSum2[i], checkSum4[i])
	}
}

func getCheckSum(name string) []byte {
	file, err := os.Open(name)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	const filechunk = 8192

	// calculate the file size
	info, _ := file.Stat()

	filesize := info.Size()

	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))

	hash := md5.New()

	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)

		file.Read(buf)
		io.WriteString(hash, string(buf)) // append into the hash
	}
	return hash.Sum(nil)
}
