package main

import "fmt"

func main() {
	slice := []int{1, 5, 2, 6, 54, 32, 121, 43, 32}
	selectionSort(slice)
	fmt.Println(slice)
}

func selectionSort(slice []int) {
	for partIndex := len(slice) - 1; partIndex > 0; partIndex-- {
		largestAt := 0
		for i := 1; i <= partIndex; i++ {
			if slice[i] > slice[largestAt] {
				largestAt = i
			}
		}
		slice[largestAt], slice[partIndex+1] = slice[partIndex+1], slice[largestAt]
	}
}
