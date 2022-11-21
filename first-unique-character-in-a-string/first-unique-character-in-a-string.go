import (
    "fmt"
)

func firstUniqChar(s string) int {
    // use array to mapping only 26 char
    // content will be count of char occur in string    
    
    var charCount [26]int
    for i := 0; i < len(s); i++ {
        charCount[int(s[i])-97]+=1
    }
    
    for i := 0; i < len(s); i++ {    
        if charCount[int(s[i])-97] == 1 {
            return i
        }
    }
    
    return -1
}