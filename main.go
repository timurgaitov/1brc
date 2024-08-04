package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

const fileSize int64 = 15874209967
const readBufSize int64 = 128 * 4096

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

	wg := sync.WaitGroup{}

	// concurrency 4 - 320ms
	// concurrency 8 - 257ms
	const concurrency = 8
	const n int64 = fileSize / concurrency
	var off int64 = 0

	for range concurrency {
		wg.Add(1)
		go readSection(filePath, off, n, &wg)
		off += n + 1
	}

	wg.Wait()
}

func readSection(filePath string, off int64, n int64, wg *sync.WaitGroup) {
	file, openErr := os.OpenFile(filePath, os.O_RDONLY, 0)
	if openErr != nil {
		panic(openErr)
	}
	defer file.Close()

	sect := io.NewSectionReader(file, off, n)
	buf := make([]byte, readBufSize)
	for {
		_, readErr := sect.Read(buf)
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			panic(readErr)
		}
	}

	wg.Done()
}
