package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

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

	file, openErr := os.OpenFile(filePath, os.O_RDONLY, 0)
	if openErr != nil {
		panic(openErr)
	}
	defer file.Close()

	// 128 * 4096 - 954ms
	const readSize int64 = 128 * 4096
	buf := make([]byte, readSize)
	for {
		_, readErr := file.Read(buf)
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			panic(readErr)
		}
	}
}
