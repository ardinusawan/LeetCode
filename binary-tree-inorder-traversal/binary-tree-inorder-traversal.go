// recursive

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
    inorder(root)

    return result
}

func inorder(root *TreeNode) {
    if root == nil {
        return
    }
    inorder(root.Left)
    result = append(result, root.Val)
    inorder(root.Right)
}