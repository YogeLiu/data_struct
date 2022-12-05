package sort

import (
	"math/rand"
)

/*
* Select Sort
* Time Complexity: O(n^2)
 */

func SelectSort(data []int64) {
	for i := 0; i < len(data); i++ {
		minIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[minIndex] > data[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			data[minIndex], data[i] = data[i], data[minIndex]
		}
	}
}

/*
* Bubble Sort
* Time Complexity: O(n^2)
 */

func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}

/*
* Insert Sort
* Time Complexity: O(n^2)
 */

func InsertSort(data []int) {
	for i := 1; i < len(data); i++ {
		j := i - 1
		cur := data[i]
		for j >= 0 && cur < data[j] {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = cur
	}
}

/*
* Qucik Sort
* Time Complexity: O(nlogn)
 */

func QuickSort(data []int) {
	if len(data) < 2 {
		return
	}
	left, right := 0, len(data)-1
	std := rand.Int() % len(data)
	data[std], data[right] = data[right], data[std]
	for i := 0; i < len(data); i++ {
		if data[i] < data[right] {
			data[left], data[i] = data[i], data[left]
			left++
		}
	}
	data[left], data[right] = data[right], data[left]
	QuickSort(data[0:left])
	QuickSort(data[left+1:])
}

/*
* Merge Sort
* Time Complexity: O(nlogn)
 */
func MergeSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	return merge(MergeSort(data[0:len(data)/2]), MergeSort(data[len(data)/2:]))
}

func merge(a, b []int) []int {
	ret := make([]int, 0, len(a)+len(b))
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			ret = append(ret, a[i])
			i++
		} else {
			ret = append(ret, b[j])
			j++
		}
	}
	if i < len(a) {
		ret = append(ret, a[i:]...)
	}
	if j < len(b) {
		ret = append(ret, b[j:]...)
	}
	return ret
}

/*
* Heap Sort
* Time Complexity: O(nlogn)
 */

func HeapSort(data []int) {
	length := len(data)
	buildMaxHeap(data, length)
	right := len(data) - 1
	for right >= 0 {
		data[right], data[0] = data[0], data[right]
		length--
		right--
		heapify(data, 0, length)
	}

}

func buildMaxHeap(data []int, length int) {
	for i := length / 2; i >= 0; i-- {
		heapify(data, i, length)
	}
}

func heapify(data []int, i, length int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < length && data[largest] < data[left] {
		largest = left
	}
	if right < length && data[right] > data[largest] {
		largest = right
	}
	if largest != i {
		data[largest], data[i] = data[i], data[largest]
		heapify(data, largest, length)
	}
}

/*
* Counting Sort
* Description: 一定范围的数组数组映射到[0,Max]数组中，统计数组每个元素出现的次数
* Time Complexity: O(n)
 */

func CountingSort(data []int) []int {
	min, max := data[0], data[0]
	for _, val := range data {
		if val < min {
			min = val
		}
		if max < val {
			max = val
		}
	}
	count := make([]int, max-min+1)
	for _, val := range data {
		count[val-min]++
	}
	ret := make([]int, 0, len(data))
	for i, val := range count {
		if val == 0 {
			continue
		}
		for cnt := 0; cnt < val; cnt++ {
			ret = append(ret, min+i)
		}
	}
	return ret
}
