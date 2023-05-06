// hashmap

func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _, v := range nums {
        if m[v] == true {
            return true    
        }
        m[v] = true
    }
    return false
}
