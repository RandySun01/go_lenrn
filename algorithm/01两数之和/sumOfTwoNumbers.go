package main

import "fmt"

/*
@author RandySun
@create 2022-04-18-21:26
*/
func twoSum(nums []int, target int) []int {
	// key:num, values: index
	dictMap := make(map[int]int)
	for index, num := range nums{
		res := target-num
		v, ok := dictMap[res]
		if ok{
			return []int{v, index}
		}
		dictMap[num] = index
	}
	return nil
}

func main() {
	nums := []int{2,7,11,15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}