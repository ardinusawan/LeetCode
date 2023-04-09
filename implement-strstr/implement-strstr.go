func strStr(haystack string, needle string) int {
    if (len(haystack) < len(needle)) { return -1 }
    for i := 0; i<len(haystack); i++ {
        left := i
        right := 0
        for left < len(haystack) && right < len(needle) && haystack[left] == needle[right] {
            left += 1
            right += 1
        }
        if right == len(needle) {
            return i
        }
    }
    return -1
}