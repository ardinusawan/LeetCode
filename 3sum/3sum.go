import (
    "fmt"
    "sort"
)

func threeSum(nums []int) [][]int {
    // first, sort nums
    // 3 pointer. v(current value), l(left), r(right)
    // init: i=0, v = nums[i], l = nums[i+1], r=nums[length-1]
    // if a same with prev value, skip
    // while l < r
        // if totalSum > 0, decrease r
        // if totalSum < 0, increase l
        // if totalSum = 0, increase l
            // increase l
            // keep increase l when nums[l] == nums[l-1] and l < r
    
    var result [][]int
    sort.Ints(nums)
    for i, v := range(nums) {
        if (i > 0 && v == nums[i-1]) {
            continue
        }
        
        l, r := i+1, len(nums)-1
        for(l < r) {
            totalSum := v + nums[l] + nums[r]
            if (totalSum > 0) {
                r -= 1
            } else if (totalSum < 0) {
                l += 1
            } else {
                result = append(result, []int{v, nums[l], nums[r]})
                l += 1
                // edge case
                // Input:
                // [0,0,0,0]
                // Output:
                // [[0,0,0],[0,0,0]]
                // Expected:
                // [[0,0,0]]
                for (nums[l] == nums[l-1] && l < r) {
                    l+=1
                }
            }
        }
    }
    
    return result
}