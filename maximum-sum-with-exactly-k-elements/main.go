func maximizeSum(nums []int, k int) int {
    sum := 0
    if len(nums) == 0 {
        return sum
    }

    biggest := nums[0]
    for _, n := range nums {
        if n > biggest {
            biggest = n
        }
    }
    
    for i:=0;i<k;i++ {
        sum += biggest
        biggest += 1
    }
  
    return sum
}
