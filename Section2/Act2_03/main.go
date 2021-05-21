package main

import "fmt"

func main() {
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println("Before:", nums)
	for j := 0; j < len(nums); j++ {
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}

	fmt.Println("After: ", nums)

}
