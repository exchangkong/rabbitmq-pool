package util

func DeleteSlice(arr []int, key int) []int {
	for i, val := range arr {
		if val == key {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}