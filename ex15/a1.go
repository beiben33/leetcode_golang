package main

import (
    "fmt"
    "arrlist/list"
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}

func threeSum(nums []int) [][]int {
    var list = list.New();
    for i := 0; i < len(nums) - 2; i++ {
        for j := i + 1; j < len(nums) - 1; j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[i] + nums[j] + nums[k] == 0 {
                    list.Add([3]int{i, j, k})
                }
            }
        }
    }

    result := make([][]int, len(list))
    for idx, v := range list.elements {
        result[idx] = v; 
	}
    return result;
}