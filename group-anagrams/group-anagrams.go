func groupAnagrams(strs []string) [][]string {
    // create a list with length 26 (a-z) char
    // iterate ever char in it, store result in the list
    // create hash map and use that list as key
    // result = iterate every values in map

    result := make(map[[26]int][]string)
    for _, v := range strs {
        count := [26]int{} // use array instead of slice

        for _, s := range v {
            count[s - 97] += 1 // 97 = ascii a
        }
        result[count] = append(result[count], v)    
    }
    var res [][]string
    for _, v := range result {
        res = append(res, v)
    }
    
    return res
}