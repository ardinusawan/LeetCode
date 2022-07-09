import "fmt"

var beenThere map[int]bool

func search(nums []int, target int) int {
    // iterative binary search
    // cause we need the index
    // recursive had to maintain the index since will modify arrays
    
    left, right := 0, len(nums) -1
    
    for left <= right {
        pivot := left + (right - left)/2
        
        if nums[pivot] == target {
            return pivot
        }
        
        if nums[pivot] < target {
            left = pivot + 1
        } else {
            right = pivot - 1
        }
    }
    
    return -1
}