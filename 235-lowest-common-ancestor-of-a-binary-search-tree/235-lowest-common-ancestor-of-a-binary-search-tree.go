/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // if p and q is smaller than current node: go to left
    // if p and q is bigger than current node: go to right
    // else, current node is LCA
    
    curr := root
    for curr != nil {
        if p.Val < curr.Val && q.Val < curr.Val {
            curr = curr.Left
        } else if p.Val > curr.Val && q.Val > curr.Val {
            curr = curr.Right
        } else {
            break
        }
    }
    
    return curr
}
