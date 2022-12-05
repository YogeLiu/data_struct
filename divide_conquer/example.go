package divideconquer

/*
* found the max number in the input list
 */

func Max(arr []int) int {
	if len(arr) <= 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		return max(arr[0], arr[1])
	}
	return max(Max(arr[0:len(arr)/2]), Max(arr[len(arr)/2:]))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
