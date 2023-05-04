// recursive dfs
// go to farther left first, when reach nil append root.Val to result
// then go to right node

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var result []int

func inorderTraversal(root *TreeNode) []int {
    result = []int{}
    dfs(root)

    return result
}

func dfs(root *TreeNode) {
    if root == nil {
        return
    }
    dfs(root.Left)
    result = append(result, root.Val)
    dfs(root.Right)
}