package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type result struct {
	Min   int16
	Max   int16
	Sum   int64
	Count int64
}

func (r *result) Mean() int64 {
	return r.Sum / r.Count
}

func main() {
	whenStarted := time.Now()
	defer func() {
		fmt.Printf("%v", time.Since(whenStarted))
		fmt.Println()
	}()

	if len(os.Args) < 1 {
		panic("pass file path")
	}
	filePath := os.Args[1]

	impl(filePath)
}

func impl(filePath string) {
	file, openErr := os.OpenFile(filePath, os.O_RDONLY, 0)
	if openErr != nil {
		panic(openErr)
	}
	defer file.Close()

	res := make(map[[100]byte]*result)

	r := bufio.NewReaderSize(file, 4096*4096)
	buf := [100]byte{}
out:
	for {
		bufI := 0

		var b byte
		var rec *result
		var ok bool
		for {
			b = rd(r)
			if b == 0 {
				break out
			}
			if b == ';' {
				for bufI < len(buf) {
					buf[bufI] = 0
					bufI++
				}
				rec, ok = res[buf]
				if !ok {
					rec = &result{}
					rec.Max = -10000
					res[buf] = rec
				}
				break
			}
			buf[bufI] = b
			bufI++
		}
		sign := int16(1)
		b = rd(r)
		f := int16(0)
		if b == '-' {
			sign = -1
		} else {
			f = int16(b - '0')
		}
		for {
			b = rd(r)
			if b == '.' {
				continue
			}
			if b == '\n' {
				break
			}
			f = f*10 + int16(b-'0')
		}
		f = f * sign
		if f > rec.Max {
			rec.Max = f
		}
		if f < rec.Min {
			rec.Min = f
		}
		rec.Sum += int64(f)
		rec.Count++
	}
}

func rd(r *bufio.Reader) byte {
	b, readErr := r.ReadByte()
	if readErr == io.EOF {
		return 0
	}
	if readErr != nil {
		panic(readErr)
	}
	return b
}
