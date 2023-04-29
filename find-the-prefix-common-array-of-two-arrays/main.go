// iterate by A
// create empty map
// for every iteration, mark A and B enum with increase value by one
// for every iteration, for loop map to check how many value > 1, result[i] += 1 when found

// Link: https://leetcode.com/contest/biweekly-contest-103/problems/find-the-prefix-common-array-of-two-arrays/

func findThePrefixCommonArray(A []int, B []int) []int {
    result := make([]int, len(A))
    m := make(map[int]int)
    
    for i, n := range A {
        m[n] += 1
        m[B[i]] += 1
        for _, value := range m {
           if(value) > 1 {
                result[i] += 1
            }
        }
    }
    fmt.Println(m)
    return result
}
