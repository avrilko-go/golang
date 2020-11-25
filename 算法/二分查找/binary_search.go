package main

import "fmt"

// 时间复杂度 log(n)
func main() {
	arr := []int{1, 3, 5, 9, 11, 100}
	num := 100
	index := binarySearch(num, arr)
	fmt.Println(index)
}

func binarySearch(num int, arr []int) int {
	low := 0
	high := len(arr) - 1

	for {
		if high < low {
			return -1
		}
		mid := (high + low) / 2
		guess := arr[mid]
		if num == guess { // 如果相同则找到了
			return mid
		}
		if num > guess { // 查找的元素大于猜测的元素，则讲low设置为mid + 1
			low = mid + 1
		}
		if num < guess { // 查找的元素小于猜测的元素，则将high设置为mid - 1
			high = mid - 1
		}
	}
}
