func twoSum(nums []int, target int) []int {
    numsH := make(map[int]int)
    var result []int
    
    for i, v := range nums {
        numsH[v] = i
    }
    
    for i, v := range nums {
        if v1, ok := numsH[target -v]; ok == true && i != v1 {
            result = append(result, []int{i, v1}...)
            break
        }
    }
    
    return result
}