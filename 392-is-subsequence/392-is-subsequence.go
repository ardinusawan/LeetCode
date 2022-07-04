import "fmt"

func isSubsequence(s string, t string) bool {
    lp, rp := 0, 0
    
    for i:=0;i<len(t);i++ {
        if lp == len(s) {
            break
        }
        
        if s[lp] == t[rp] {
            lp += 1
        }
        rp += 1
    }
    
    return lp == len(s)
}