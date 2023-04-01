import (
    "fmt"
    "math"
)

func lengthOfLongestSubstring(s string) int {
    // sliding windows
    // input: "pwwkew"
    
    // use the byte type as the key type for the charSet map, and store the index of each character in the map instead of the character itself. When a repeated character is encountered, we check if its index is greater than or equal to the left pointer. If it is, we update the left pointer to the next position after the index of the repeated character. Finally, we update the charSet map with the index of the current character, and update the result variable if necessary.
    
    var result, left int
    charSet := make(map[byte]int) 
    // key: char, value: index
    
    for right:=0;right<len(s);right++ {

        if index, found := charSet[s[right]]; found && index >= left {
            left = index + 1 // left = 2
        }
        
        charSet[s[right]] = right
 
        result = int(math.Max(float64(result), float64(right - left + 1)))
    }
    
    return result
}