package main

import (
	"fmt"
	"time"
)

func main() {
	n := 1000000
	testSlice1 := make([]int, 0)    // Without preallocation
	testSlice2 := make([]int, 0, n) // Preallocate slice with capacity n

	fmt.Printf("Total time without preallocation: %v \n", timeLoop(testSlice1, n))
	fmt.Printf("Total time with preallocation: %v \n", timeLoop(testSlice2, n))

}

// function timeLoop returns data type time.Duration specifically for time measurement.
func timeLoop(slice []int, n int) time.Duration {
	start := time.Now()
	for i := 0; i < n; i++ {
		slice = append(slice, i) //append a value to the slice at position i.
	}
	elapsed := time.Since(start)
	return elapsed
}
