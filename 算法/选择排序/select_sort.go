package main

import "fmt"

// 时间复杂度 n*n
func main() {
	arr := []int{1, 7, 3, 88, 4, 0, 2}
	selectSort(arr, 0)
	fmt.Println(arr)
}

func selectSort(arr []int, start int) {
	if start == len(arr) {
		return
	}
	minIndex := start
	minNum := arr[start]
	for i := start + 1; i < len(arr); i++ {
		guess := arr[i]
		if guess < minNum {
			minIndex = i
			minNum = guess
		}
	}

	// 交换位置
	arr[start], arr[minIndex] = arr[minIndex], arr[start]
	selectSort(arr, start+1)
}
