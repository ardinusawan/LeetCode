// using slow and fast pointer
// slow = run 1 node, fast run 2 node
// when end of the node is not nil, then it must be in cycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    var slow, fast *ListNode
    slow = head
    fast = head
    for fast != nil && fast.Next != nil && fast.Next.Next != nil && slow != nil && slow.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        
        if fast == slow {
            return true
        }
    }
    
    return false
}