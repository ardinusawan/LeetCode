import (
    "fmt"
    "math"
)

func lengthOfLongestSubstring(s string) int {
    // sliding windows
    
    var result, left int
    charSet := make(map[byte]int)
    for right:=0;right<len(s);right++ {
        // for (charSet[s[right]]) {
        //     delete(charSet, s[right])
        //     left+=1
        // }
        if index, found := charSet[s[right]]; found && index >= left {
            // delete(charSet, s[right])
            left = index + 1
        }
        // fmt.Println(index, found)
        charSet[s[right]] = right
        result = int(math.Max(float64(result), float64(right - left + 1)))
    }
    // fmt.Println(charSet)
    return result
}