/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    mapper := make(map[int][]int)
    dfs(root, 0, mapper)
    result := make([][]int, len(mapper))
    for i, v := range(mapper) {
        result[i] = v
    }
    
    return result
}

func dfs(root *TreeNode, level int, mapper map[int][]int) {
    if root == nil {
        return
    }
    mapper[level] = append(mapper[level], root.Val)
    
    if root.Left != nil {
        dfs(root.Left, level+1, mapper)
    }
    if root.Right != nil {
        dfs(root.Right, level+1, mapper)
    }
}