package main

import (
    "fmt"
    "pub/arrlist/list"
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
                    list.Add([]int{nums[i], nums[j], nums[k]})
                }
            }
        }
    }

    result := make([][]int, list.Size())
    for i := 0; i < list.Size(); i++ {
        result[i] = list.Get(i).([]int)
    }
    return result;
}