// buckets sort algorithm
// 1. create map[int]int, iterate nums and set map with key = num of nums, value = frequent of num in nums
// [1,1,1,2,2,3] -> map[1:3, 2:2, 3: 1]
// 2. create buckets type is list of list, where index = frequent and value = list of num
// [1,1,1,2,2,3] -> [[], [3], [2], [1]]
// 3. result = []int
// for loop descending of buckets, dont stop while len(result) < k

func topKFrequent(nums []int, k int) []int {
    if (len(nums)<=1) {
        return nums
    }

    m := make(map[int]int)
    for _, v := range nums {
        m[v]++
    }

    buckets := make([][]int, len(nums)+1)
    for num, freq := range m {
        buckets[freq] = append(buckets[freq], num)
    }

    var result []int
    for i:=len(buckets)-1; i>0 && len(result)<k; i-- {
        if len(buckets[i]) > 0 {
            result = append(result, buckets[i]...)
        }
    }
    return result
}
