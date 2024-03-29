// If the length of s1 is greater than s2, return false.
// Initialize count maps for s1 and s2 with a size of 26.
// Count the occurrences of each character in the initial window of s1 and s2.
// Check if the counts of all characters in s1 and s2 are the same. Increment matches for each match.
// Slide the window over s2, updating the counts accordingly.
// If matches reaches 26, return true.
// If a character count in s2 is one more or one less than its count in s1, decrement matches.
// Return false if no match is found.
func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    s1Count, s2Count := make(map[int]int, 26), make(map[int]int, 26)
    for i := byte('a')-97; i <= byte('z'); i++ {
		s1Count[int(i)-97] = 0
        s2Count[int(i)-97] = 0
	}
    
    for i := 0; i < len(s1); i++ {
        s1Index := int(s1[i] - 'a')
        s2Index := int(s2[i] - 'a')
        s1Count[s1Index]++
        s2Count[s2Index]++
    }

    matches := 0
    
    for i:=0; i<26; i++ {
        if s1Count[i] == s2Count[i] {
            matches++
        }
    }

    left := 0
    for right:=len(s1); right<len(s2); right++ {
        if matches == 26 {
            return true
        }

        index := int(s2[right] - 'a')
        s2Count[index]++
        if s1Count[index] == s2Count[index] { 
            matches++
        } else if s1Count[index] + 1 == s2Count[index] {
            matches--
        }

        index2 := int(s2[left] - 'a')
        s2Count[index2]--
        if s1Count[index2] == s2Count[index2] { 
            matches++
        } else if s1Count[index2] - 1 == s2Count[index2] {
            matches--
        }

        left++
    }
    return matches == 26
}