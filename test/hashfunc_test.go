// Author: Shubhankar Ranade
// contains benchmarks for hash functions

package test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/neelkshah/pandora/pkg/hashfunc"
)

var numWords = 12823

func getData() ([]string, error) {
	var lines []string = make([]string, 0, numWords)
	file, err := os.Open("../assets/arabian_nights.txt")

	if err != nil {
		return lines, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		if len(line) > 2 {
			lines = append(lines, line)
		}

		if err != nil {
			break
		}
	}
	if err != nil && err != io.EOF {
		return nil, err
	}

	return lines, nil
}

var testStrings, err = getData()

func BenchmarkHashing(b *testing.B) {
	hashFuncs := []struct {
		name     string
		hashFunc func([]byte, uint64) uint64
	}{
		{"FNV-1a", hashfunc.FNV64a},
		{"FNV-1", hashfunc.FNV64},
		{"DJB2a", hashfunc.DJB2a},
		{"DJB2", hashfunc.DJB2},
		{"SDBM", hashfunc.SDBM},
		{"PJW", hashfunc.PJW},
		{"Murmur", hashfunc.Murmur},
	}

	hashtableSize := uint64(len(testStrings))
	fmt.Println("Hashtable size: ", hashtableSize)

	for _, hasher := range hashFuncs {
		b.Run(fmt.Sprintf("%s", hasher.name), func(b *testing.B) {
			var collisions int64 = 0
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				hashArr := make([]int64, hashtableSize) // for counting collisions
				for i := range hashArr {
					hashArr[i] = -1
				}
				var num uint64 = 0
				b.StartTimer()

				for _, str := range testStrings {
					hashArr[hasher.hashFunc([]byte(str), hashtableSize)]++
					// hashArr[hasher.hashFunc(helper.IntToByte(num), hashtableSize)]++
					num++
				}

				b.StopTimer()
				for i := range hashArr {
					if hashArr[i] > 0 {
						collisions += hashArr[i]
					}
				}
				b.StartTimer()
			}
			b.ReportMetric(float64(collisions)/float64(b.N), "collisions/op")
			b.ReportMetric(float64(len(testStrings)), "keys")
		})
	}
}
