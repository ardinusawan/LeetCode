import (
    "fmt"
)

func longestPalindrome(s string) string {
    // start from the middle, then expand to left and right
    // when left and right is same, continue
    // when left in boundary, move middle value. so same with right
    // stop: when middle == right
    
    maxLeft := 0
    maxRight := 0
    resultLen :=0
    
    for i:=0; i<len(s); i++ {
        
        // odd cases
        left , right := i, i
        for (left >=0 && right < len(s) && s[left] == s[right]) {
            if ((right - left + 1) > resultLen) {
                maxLeft = left
                maxRight = right  
                resultLen = right - left + 1
            }
            left -= 1
            right += 1
        }
        
        // even cases
        left , right = i, i+1
        for (left >=0 && right < len(s) && s[left] == s[right]) {
            if ((right - left + 1) > resultLen) {
                maxLeft = left
                maxRight = right
                resultLen = right - left + 1

            }
            left -= 1
            right += 1
        }
        
    }
    
    result := ""
    for i:=maxLeft; i<=maxRight; i++ {
        result = result + string(s[i])
    }
    
    return result
}