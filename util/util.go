package util

import "log"

func DeleteSlice(arr []int, key int) []int {
	for i, val := range arr {
		if val == key {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}
