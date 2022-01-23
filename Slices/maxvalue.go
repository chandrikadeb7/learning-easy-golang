package main

import "fmt"

func main() {
	nums := []int{15, 10, 38, 89, 4, 78, 90, 89}
	fmt.Println(nums)
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if result < nums[i] {
			result = nums[i]
		}

	}
	fmt.Printf("Max value: %v\n", result)

}
