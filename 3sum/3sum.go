// first, sort nums
// 3 pointer. i(current value), l(left), r(right)
// init: i=0, i = nums[i], l = nums[i+1], r=nums[length-1]
// if i > 0 and prev value == current value, continue
// while l < r
    // if totalSum > 0, decrease r
    // if totalSum < 0, increase l
    // if totalSum = 0, increase l
        // increase l
        // keep increase l when nums[l] == nums[l-1] and l < r

import (
    "sort"
)

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{}
    for i, _ := range nums {
        // check if duplicate for i with left
        if (i > 0 && nums[i] == nums[i-1]) {
            continue
        }
        
        l , r := i + 1, len(nums)-1
        for (l < r ) {
            totalSum := nums[i] + nums[l] + nums[r]
            if totalSum > 0 {
                r -= 1
            } else if totalSum < 0 {
                l += 1
            } else {
                result = append(result, []int{nums[i], nums[l], nums[r]})
                l += 1

                // check if duplicate for new left pointer
                for (nums[l] == nums[l-1] && l < r) {
                    l+=1
                }
            }
        }
    }
    return result
}
