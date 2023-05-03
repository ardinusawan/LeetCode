// dfs left and right node same time
// when value of two node is same, keep walking the dfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return dfs(root.Left, root.Right)
}

func dfs(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    } else if left == nil || right == nil {
        return false
    }
 
    
    return left.Val == right.Val && dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}