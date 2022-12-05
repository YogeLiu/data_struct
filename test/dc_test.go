package test

import (
	"errors"
	"math/rand"
	"sort"
	"testing"

	dc "github.com/YogeLiu/data_struct/divide_conquer"
)

func TestMax(t *testing.T) {
	// arr := []int{1, 3, 7, 8, 10, 9, 8, 45, 56}
	// fmt.Printf("dc.Max(arr): %v\n", dc.Max(arr))
	// return
	for i := 0; i < 100; i++ {
		length := rand.Intn(20)
		if length < 1 {
			continue
		}
		arr := []int{}
		for cnt := 0; cnt < length; cnt++ {
			arr = append(arr, rand.Int())
		}
		max := dc.Max(arr)
		sort.Ints(arr)
		if arr[len(arr)-1] != max {
			panic(errors.New("element not equal"))
		}

	}
}
