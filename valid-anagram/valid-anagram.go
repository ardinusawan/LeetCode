import (
    "reflect"
)

// using map
// {char: countOfChar}
// compare
func isAnagram(s string, t string) bool {
    sMap, tMap := make(map[string]int), make(map[string]int)
    
    for _, v := range s {
        char := string(v)
        sMap[char] += 1
    }
    
    for _, v := range t {
        char := string(v)
        tMap[char] += 1
    }
    
    return reflect.DeepEqual(sMap, tMap)
}