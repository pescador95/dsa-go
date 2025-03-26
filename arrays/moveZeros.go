package main

func moveZeros(arr []int) []int {
	j := 0
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		if arr[i] != 0 && arr[j] == 0 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		if arr[j] != 0 {
			j++
		}
	}
	return arr
}
