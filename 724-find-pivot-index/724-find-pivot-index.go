import "fmt"

func pivotIndex(nums []int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    
    left := 0
    right := sum
    for i, n := range nums {
        right = right - n
        
        if left == right {
            return i
        }
        
        left = left + n
    }
    return -1
}