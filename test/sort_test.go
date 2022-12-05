package test

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"testing"

	my_sort "github.com/YogeLiu/data_struct/sort"
)

func TestSelectSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		length := rand.Intn(100)
		arr := []int64{}
		for cnt := 0; cnt < length; cnt++ {
			arr = append(arr, rand.Int63())
		}
		stdArr := make([]int64, length)
		copy(stdArr, arr)
		my_sort.SelectSort(arr)
		sort.Slice(stdArr, func(i, j int) bool { return stdArr[i] < stdArr[j] })
		for i := 0; i < len(stdArr); i++ {
			if stdArr[i] != arr[i] {
				panic(errors.New("element not equal"))
			}
		}
	}
}

func TestQucikSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		length := rand.Intn(100)
		arr := []int{}
		for cnt := 0; cnt < length; cnt++ {
			arr = append(arr, rand.Int())
		}
		testify(arr, my_sort.QuickSort)
	}
}

func TestMergeSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		length := rand.Intn(100)
		arr := []int{}
		for cnt := 0; cnt < length; cnt++ {
			arr = append(arr, rand.Int())
		}
		stdArr := make([]int, len(arr))
		copy(stdArr, arr)
		sort.Ints(stdArr)
		arr = my_sort.MergeSort(arr)
		for i := 0; i < len(stdArr); i++ {
			if stdArr[i] != arr[i] {
				panic(errors.New("element not equal"))
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		length := rand.Intn(100)
		arr := []int{}
		for cnt := 0; cnt < length; cnt++ {
			arr = append(arr, rand.Int())
		}
		testify(arr, my_sort.BubbleSort)
	}
}

func TestInsertSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		length := rand.Intn(100)
		arr := []int{}
		for cnt := 0; cnt < length; cnt++ {
			arr = append(arr, rand.Int())
		}
		testify(arr, my_sort.InsertSort)
	}
}

func TestHeapSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		length := rand.Intn(100)
		arr := []int{}
		for cnt := 0; cnt < length; cnt++ {
			arr = append(arr, rand.Int())
		}
		testify(arr, my_sort.HeapSort)
	}
}

func testify(arr []int, algorithm func(data []int)) {
	stdArr := make([]int, len(arr))
	copy(stdArr, arr)
	algorithm(arr)
	sort.Ints(stdArr)
	for i := 0; i < len(stdArr); i++ {
		if stdArr[i] != arr[i] {
			panic(errors.New("element not equal"))
		}
	}
}

func TestCountingSort(t *testing.T) {
	arr := []int{5, 10, 12, 15, 7, 9, 11, 5, 6, 7, 9, 12, 13, -1, -5, -5, -3, -2, -3, -1}
	arr = my_sort.CountingSort(arr)
	fmt.Printf("arr: %+v\n", arr)
}
