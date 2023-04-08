func myAtoi(s string) int {
    var result int
    sign, found := 1, false
    
    // Remove leading whitespaces
    s = strings.TrimSpace(s)
    
    for _, c := range s {
        if c == '-' && !found {
            sign, found = -1, true
        } else if c == '+' && !found {
            found = true
        } else if '0' <= c && c <= '9' {
            found = true
            result = result*10 + int(c-'0')
            // Check overflow
            if sign*result > math.MaxInt32 {
                return math.MaxInt32
            } else if sign*result < math.MinInt32 {
                return math.MinInt32
            }
        } else {
            // Break the loop on the first non-digit character
            break
        }
    }
    
    return sign*result
}
