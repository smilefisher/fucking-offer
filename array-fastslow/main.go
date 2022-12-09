package main

import "fmt"

func main() {

	arr := []int{1, 1}

	f2 := func(arr []int) []int {
		var fast, slow = 1, 0
		length := len(arr)
		if length == 1 {
			return arr
		}
		for fast < length {
			if arr[fast] != arr[slow] {
				slow++
				arr[slow] = arr[fast]
			}
			fast++
		}
		return arr[:slow+1]
	}
	fmt.Println(f2(arr))

}
