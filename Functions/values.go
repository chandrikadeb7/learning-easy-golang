package main

import (
	"fmt"
)

//triples index i at slice nums
func double(values []int, i int) {
	values[i] *= 3
}

func doublePtr(a *int) {
	*a *= 2
}

func main() {
	//declare a slice
	nums := []int{1, 2, 3, 4}
	//triples index 3
	double(nums, 3)
	fmt.Println(nums)

	val := 5
	doublePtr(&val)
	fmt.Printf("Double int value 5 is: %v\n", val)

}
