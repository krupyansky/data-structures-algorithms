package main

import "fmt"

func main() {
	slice := []int{1, 5, 2, 6, 54, 32, 121, 43, 32}
	shellSort(slice)
	fmt.Println(slice)
}

func shellSort(slice []int) {
	gap := 1
	for gap < len(slice)/3 {
		gap = 3*gap + 1
	}

	for gap >= 1 {
		for i := gap; i < len(slice); i++ {
			for j := i; j >= gap && slice[j] < slice[j-gap]; j -= gap {
				slice[j], slice[j-gap] = slice[j-gap], slice[j]
			}
		}

		gap /= 3
	}
}
