package main

import "fmt"

func main() {
	slice := []int{1, 5, 2, 6, 54, 32, 121, 43, 32}
	mergeSort(slice)
	fmt.Println(slice)
}

func mergeSort(slice []int) {
	length := len(slice)
	if length <= 1 {
		return
	}

	sort(slice, 0, length-1)
}

func sort(slice []int, low, high int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2

	sort(slice, low, mid)
	sort(slice, mid+1, high)

	merge(slice, low, mid, high)
}

func merge(slice []int, low, mid, high int) {
	aux := make([]int, high-low+1)

	i := low
	j := mid + 1
	k := 0

	for ; i <= mid && j <= high; k++ {
		if slice[i] < slice[j] {
			aux[k] = slice[i]
			i++
		} else {
			aux[k] = slice[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		aux[k] = slice[i]
		k++
	}

	for ; j <= high; j++ {
		aux[k] = slice[j]
		k++
	}

	copy(slice[low:high+1], aux)
}
