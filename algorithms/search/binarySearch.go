package main

import "fmt"

func main() {
	slice := []int{1, 5, 10, 15, 54, 234, 300}
	indexFirst := binarySearch(slice, 5643)
	indexSecond := recursiveBinarySearch(0, len(slice), slice, 5643)
	fmt.Println(indexFirst, indexSecond)
}

func binarySearch(slice []int, value int) int {
	low := 0
	high := len(slice)

	for low < high-1 {
		mid := (low + high) / 2

		if slice[mid] == value {
			return mid
		}
		if slice[mid] < value {
			low = mid
		}
		if slice[mid] > value {
			high = mid
		}
	}

	return -1
}

func recursiveBinarySearch(low, high int, slice []int, value int) int {
	if low >= high {
		return -1
	}

	mid := (low + high) / 2

	if slice[mid] == value {
		return mid
	}

	if slice[mid] < value {
		return recursiveBinarySearch(mid+1, high, slice, value)
	} else {
		return recursiveBinarySearch(low, mid, slice, value)
	}
}
