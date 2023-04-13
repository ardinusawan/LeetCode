func increasingTriplet(nums []int) bool {    
    if len(nums) < 3 {
        return false
    }

    first := math.MaxInt32
    second := math.MaxInt32

    for _, current := range nums {
        if (current < first) {
            first = current
        } else if (current > first && current < second ) {
            second = current
        }
        
        if (current > second && second > first ) {
            return true
        }
    }
    
    return false
}
