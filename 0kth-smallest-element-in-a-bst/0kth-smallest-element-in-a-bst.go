// dfs and store value in stack (array)
// when root.Left != nil, go to the left again
// when root.Left = nil, append current node
// when root.Right != nil, go to the right again
// if len(array) == k, stop. Result it top of array

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var arr []int

func kthSmallest(root *TreeNode, k int) int {
    arr = []int{}
   
    dfs(root, k)
    return arr[k-1]
}

func dfs(root *TreeNode, k int) {
    if root == nil {
        return
    }
    
    if root.Left != nil {
        dfs(root.Left, k)
    }
    
    arr = append(arr, root.Val)
    if len(arr) == k {
        return
    }
    
    if root.Right != nil {
        dfs(root.Right, k)
    }
}
