func maxProfit(prices []int) int {
    biggest, lowest := prices[0], prices[0]
    maxDiff := 0
    
    for _, v := range prices {
        if v > biggest {
            biggest = v
        }
               
        if maxDiff < (biggest - lowest) {
            maxDiff = biggest - lowest
        }
        
        if v < lowest {
            biggest = v
            lowest = v
        }
    }
    
    return maxDiff
}