// https://leetcode.com/problems/product-of-array-except-self/
// init: create var result with length == nums
// step 1: calculate prefix inside result
// goal: make [1,2,3,4] to [1,1,2,6]
// iterate from begining of nums. firstPrefix = 1, result[i] = currentPrefix. new currentPrefix = currentPrefix * nums[i]

// step 2: calculate postfix inside result as well
// goal: make [1,1,2,6] to [24,12,8,6]
// iterate from last to first nums. currentPostfix = 1, result = currentPostfix * nums[i]. new currentPostfix = currentPostfix * nums[i]

func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums))
    currentPrefix := 1
    for i:=0; i<len(nums); i++ {
        result[i] = currentPrefix
        currentPrefix *= nums[i]
    }

    currentPostfix := 1
    for i:=len(nums)-1; i>=0; i-- {
        result[i] *= currentPostfix
        currentPostfix *= nums[i]
    }
    
    return result
}
