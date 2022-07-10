// first round
// left = -inf < root.Left.Val < root.Val
// right = root.Val > root.Val < root.Val

// when it zig-zag value of left and right will be prev 2 parent node
// this usefull to make sure current value not invalid BST

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
    "math"
)
func isValidBST(root *TreeNode) bool {
    // -inf < root < inf
    return validator(root, math.Inf(-1), math.Inf(1))
}

func validator(root *TreeNode, left, right float64) bool {
    if root == nil {
        return true
    }
    
    if (left >= float64(root.Val) || right <= float64(root.Val)) {
        return false
    }
    
    return validator(root.Left, float64(left), float64(root.Val)) && 
    validator(root.Right, float64(root.Val), float64(right))
}