// 2 pointer, left and right
// init max amount to 0
// left: index 0, right: last index
// while left < right, calculate currentAmount and compare with maxAmount. store the biggest to maxAmount
//   if height[l] < height[r], left move to the right and vice versa

import (
    "math"
)

func maxArea(height []int) int {
    l , r := 0, len(height)-1
    maxAmount := 0
    for l < r {
        currentAmount := int(math.Min(float64(height[l]), float64(height[r]))) * (r - l)
        if currentAmount > maxAmount {
            maxAmount = currentAmount
        }

        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }
    return maxAmount
}
