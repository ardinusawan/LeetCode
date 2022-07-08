/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "fmt"

func levelOrder(root *TreeNode) [][]int {
    return bfs(root, 0, [][]int{})
}

func bfs(root *TreeNode, level int, result [][]int) [][]int {
    if root == nil {
        return result
    }
    
    if len(result) == level {
        result = append(result, []int{root.Val})    
    } else {
        result[level] = append(result[level], root.Val)
    }
    
    level += 1
    if root.Left != nil {
        left := bfs(root.Left, level, result)    
        result = left
    }
    
    if root.Right != nil {
        right := bfs(root.Right, level, result)
        result = right
    }
    
    return result
}