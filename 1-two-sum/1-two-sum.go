func twoSum(nums []int, target int) []int {
    result := []int{}
    hash := make(map[int]int)
    for _, v := range nums {
        hash[v] += 1
    }
    
    for i, v := range nums {
        if _, ok := hash[target - v]; ok {
            if target - v == v && hash[v] == 1 {
                continue
            }
            result = append(result, i)
        }
    }
    return result 
}