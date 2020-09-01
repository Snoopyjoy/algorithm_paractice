package main

import "fmt"

func main() {
	arr := []int{0, 2, 3, 1, 1}
	QuickSort(arr, 0, 4)
	// arr = []int{1, 2, 3, 4, 5}
	// patition(arr, 0, 3)
	// fmt.Println(arr)
}

func QuickSort(arr []int, i, j int) {
	if i >= j {
		return
	}
	pos := patition(arr, i, j)
	QuickSort(arr, i, pos-1)
	QuickSort(arr, pos+1, j)
	fmt.Println(arr)
}

func patition(arr []int, i, j int) int {
	start := i
	pivot := arr[i]
	for i < j {
		for i < j && arr[j] > pivot {
			j--
		}
		for i < j && arr[i] <= pivot {
			i++
		}
		if i < j {
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[start], arr[i] = arr[i], arr[start]
	return i
}
