// use first array as comparison, NO NEED FIND SHORTEST WORD!
// init result with empty string
// for every element in first array (index: i)
//   for every string in list
//     if length of first array is out of bound || char is different then return result
//     else continue
//   append result with strs[0][i]

func longestCommonPrefix(strs []string) string {
    result := ""
    if (len(strs) == 0) {
        return result
    }
    
    for i := range strs[0] {
        for _, s := range strs {
               if (i == len(s) || s[i] != strs[0][i]) {
                return result
            }
        }
        result += string(strs[0][i])
    }
    return result
}