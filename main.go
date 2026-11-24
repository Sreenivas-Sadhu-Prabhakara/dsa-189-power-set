package main

import "fmt"

// PowerSet returns all 2^n subsets of nums. Each number from 0 to 2^n-1 is
// read as a membership bitmask: bit i set means nums[i] is in the subset.
func PowerSet(nums []int) [][]int {
	n := len(nums)
	out := make([][]int, 0, 1<<n)
	for mask := 0; mask < 1<<n; mask++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				subset = append(subset, nums[i])
			}
		}
		out = append(out, subset)
	}
	return out
}

func main() {
	fmt.Println(PowerSet([]int{1, 2, 3}))
}
