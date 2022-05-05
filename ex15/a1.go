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
    var resList = list.New();
    for i := 0; i < len(nums) - 2; i++ {
        for j := i + 1; j < len(nums) - 1; j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[i] + nums[j] + nums[k] == 0 {    
                    var element = []int{nums[i], nums[j], nums[k]}   
                    if !isResListContained(resList, element) {          
                        resList.Add(element)
                    }
                }
            }
        }
    }

    result := make([][]int, resList.Size())
    for i := 0; i < resList.Size(); i++ {
        result[i] = resList.Get(i).([]int)
    }
    return result;
}

func isResListContained(resList *list, elements []int) bool {
    flagRes := make([]bool, resList.Size())
    flagElems := make([]bool, len(elements))

    for i < resList.Size() {
        resElements = resList.Get(i).([]int)
        
        for j < len(resElements) {
            for k < len(elements) {
                if resElements[j] == elements[k] && !flagElems[k] {
                    flagRes[j] = true
                    flagElems[k] = true
                }                    
            }
        }
    }    

    var result = true

    for _, v := range flagRes {
        result = result && v
    }

    for _, v := range flagElems {
        result = result && v
    }
        
    return result
}