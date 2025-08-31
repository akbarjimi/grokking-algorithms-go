package main

import (
	"GrokkingAlgorithmsGo/Chapter01"
	"fmt"
)

func main() {
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39}
	var index = Chapter01.BinarySearch(20, arr)
	if index == -1 {
		fmt.Println("We don't have it")
	} else {
		fmt.Println(arr[index])
	}
	// switch on chapter and run its function via CLI
}
