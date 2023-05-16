// two pointer, start from start and end
// when start + end > target, decrease end.
// when start + end < target, increase start.

func twoSum(numbers []int, target int) []int {
    start, end := 0, len(numbers)-1
    for (numbers[start] + numbers[end] != target) {
        if numbers[start] + numbers[end] > target {
            end -= 1
        } else {
            start += 1
        }
    }
    return []int{start+1, end+1}
}
