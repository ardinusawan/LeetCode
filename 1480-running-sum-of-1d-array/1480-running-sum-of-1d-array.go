import (
"fmt"
)

func runningSum(nums []int) []int {
    for i, _ := range nums[1:] {
        nums[i+1] = nums[i] + nums[i+1]
    }
    return nums
}