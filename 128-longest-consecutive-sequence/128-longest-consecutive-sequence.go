// use map[int]bool, fill map from nums
// nums: [100,4,200,1,3,2] -> map[100: true, 4: true, 200: true, 1: true, 3: true, 2: true]
// iterate map, find key which dont have previous value in map. i.e: 100, 200, 1
// for every value, iterate to find longest seq exist in map

func longestConsecutive(nums []int) int {
    m := make(map[int]bool)
    longestSeq := 0
    for _, v := range nums {
        m[v] = true
    }
    
    for k, _ := range m {
        if ok := m[k-1]; !ok {
            start := 0
            for(m[k+start]) {
                start += 1
                if longestSeq < start {
                    longestSeq = start
                }
            }
        }
    }

    return longestSeq
}
