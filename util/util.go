package util

func DeleteSlice(arr []int, key int, tag string) []int {
	arrLen := len(arr)
	for i, val := range arr {
		if val == key {
			if i == arrLen { //最后一个元素在加则越界
				i--
			}

			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}