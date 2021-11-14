package main

import "fmt"

func main() {
	slice := []int{1, 5, 2, 6, 54, 32, 121, 43, 32}
	bubbleSort(slice)
	fmt.Println(slice)
}

func bubbleSort(slice []int) {
	for partIndex := len(slice) - 1; partIndex > 0; partIndex-- {
		for i := 0; i < partIndex; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
}
