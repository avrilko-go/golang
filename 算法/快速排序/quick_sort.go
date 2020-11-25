package main

import "fmt"

// 时间复杂度 nlogn

func main() {
	list3 := []int{5, 9, 1, 8, 14, 6, 49, 25, 4, 3}
	quickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)
}

func quickSort(arr []int, begin, end int) {
	if begin < end {
		mid := partition(arr, begin, end)
		quickSort(arr, begin, mid-1)
		quickSort(arr, mid+1, end)
	}
}

func partition(arr []int, begin, end int) int {
	i := begin + 1
	j := end

	for i < j {
		if arr[i] > arr[begin] { // 从左边开始有大于基值的直接丢到最右边
			arr[i], arr[j] = arr[j], arr[i]
			j-- // 最右边的位置被占用了需要往前面移动一个位置
		} else { // 没有的话不动，最左边占用了一个位置需要往右边一定一个
			i++
		}
	}

	if arr[i] >= arr[begin] { // 这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
		i--
	}
	// 将基础位置和中间位置的值进行交换
	arr[begin], arr[i] = arr[i], arr[begin]

	return i
}
