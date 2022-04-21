func plusOne(digits []int) []int {
    result := []int{}
    incr := 1
    
    for i:=len(digits)-1; i>=0; i-- {
        if digits[i] + incr == 10 {
            result = append([]int{0}, result...) // prepend
        } else {
            result = append([]int{(digits[i] + incr)}, result...) // prepend
            incr = 0
        }
    }
    
    // edge case
    // input = 9, output = [1, 0]
    if incr == 1 {
        result = append([]int{incr}, result...)
    }
    return result
}