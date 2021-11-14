package main

import "fmt"

func main() {
	slice := []int{1, 5, 2, 6, 54, 32, 121, 43, 32}
	insertionSort(slice)
	fmt.Println(slice)
}

func insertionSort(slice []int) {
	for partIndex := 1; partIndex < len(slice); partIndex++ {
		curUnsorted := slice[partIndex]
		i := 0
		for i = partIndex; i > 0 && slice[i-1] > slice[i]; i-- {
			slice[i-1], slice[i] = slice[i], slice[i-1]
		}
		slice[i] = curUnsorted
	}
}
