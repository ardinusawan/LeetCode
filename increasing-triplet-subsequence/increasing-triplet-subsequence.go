func increasingTriplet(nums []int) bool {
    // init two variable, first and second with max value
    // iterate nums. for each element do check
    // if current elem < first? -> first = elem
    // else if current elem > first BUT current < second? -> second = current elem
    // check if first > second > current elem
    
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
