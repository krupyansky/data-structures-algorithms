package main

import "fmt"

func main() {
	array := []int{1, 5, 2, 6, 54, 32, 121, 43, 32}
	insertionSort(array)
	fmt.Println(array)
}

func insertionSort(array []int) {
	for partIndex := 1; partIndex < len(array); partIndex++ {
		curUnsorted := array[partIndex]
		i := 0
		for i = partIndex; i > 0 && array[i-1] > array[i]; i-- {
			array[i-1], array[i] = array[i], array[i-1]
		}
		array[i] = curUnsorted
	}
}
