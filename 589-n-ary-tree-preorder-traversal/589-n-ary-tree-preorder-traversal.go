/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    // dfs
    result := dfs(root, []int{})
    return result
}

func dfs(root *Node, result []int) []int {
    if root == nil {
        return result
    }
    
    result = append(result, root.Val)
    for _, r := range root.Children {    
        result = dfs(r, result)
    }
    return result    
}