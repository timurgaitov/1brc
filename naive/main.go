package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type result struct {
	Min   float64
	Max   float64
	Sum   float64
	Count float64
}

func (r *result) Mean() float64 {
	return r.Sum / r.Count
}

// 2min
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

	file, openErr := os.Open(filePath)
	if openErr != nil {
		panic(openErr)
	}
	defer file.Close()

	res := make(map[string]*result)
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		split := strings.Split(line, ";")
		key := split[0]
		val, readErr := strconv.ParseFloat(split[1], 64)
		if readErr != nil {
			break
		}
		cur := res[key]
		if cur == nil {
			cur = &result{}
			res[key] = cur
		}
		cur.Min = math.Min(cur.Min, val)
		cur.Max = math.Max(cur.Max, val)
		cur.Sum += val
		cur.Count++
	}

	for k, v := range res {
		fmt.Println(k, v, v.Mean())
	}
}
