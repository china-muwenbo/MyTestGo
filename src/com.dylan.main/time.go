package main

import (
	"time"
	"fmt"
)

func main() {
	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	duration := time.Now().Sub(start)
	fmt.Println(int64(duration / time.Millisecond))
}
