import (
    "unicode"
)

func isPalindrome(s string) bool {
    cleanString := ""
    for i:=0;i<len(s);i++ {
        v := s[i]
        if v >= 65 && v <=90 || v >=97 && v<=122 {
            cleanString += string(unicode.ToLower(rune(v)))
        } else if v >= 48 && v <= 57 {
            cleanString += string(v)
        }
    }
    
    // if len(cleanString) == 1 && len(s) != 1 { return false }
    
    left, right := 0, len(cleanString)-1
    for left <= right {
        if (cleanString[left] != cleanString[right]) {
            return false
        }
        left += 1
        right -= 1
    }
    
    return true
}