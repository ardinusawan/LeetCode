import "fmt"

func longestPalindrome(s string) int {
    dict := make(map[string]int)
    
    for _, v := range s {
        dict[string(v)] += 1
    }
    
    sum := 0
    center := 0
    for i, v := range dict {
        if v % 2 == 0 {
            sum += v
            fmt.Println("1", sum, v)
            dict[i] -= v
        } else if v > 1 {
            sum += v - 1
            fmt.Println("2", sum, v - 1)

            dict[i] -= v - 1
        }
        
        if v % 2 == 1 {
            center = 1
        }
    }
    
    return sum + center
}