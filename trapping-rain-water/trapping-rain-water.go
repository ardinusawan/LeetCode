// pointer: left = 0, right = len(height)-1
// maxLeft = height[0], maxRight = height[-1]
// main idea: for left < right, find distance between maxValue - currentVal. Do this thing from left and right pointer
// when height[left] < height[right], left++ and vice versa
// then update maxValue with max(maxValue, height[i])

func trap(height []int) int {
    left, right := 0, len(height)-1
    maxLeft, maxRight := height[left], height[right]
    result := 0
    for left < right {
        if height[left] < height[right] {
            result += max(maxLeft - height[left], 0)
            left++
            maxLeft = max(maxLeft, height[left])
        } else {
            result += max(maxRight - height[right], 0)
            right--
            maxRight = max(maxRight, height[right])
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}