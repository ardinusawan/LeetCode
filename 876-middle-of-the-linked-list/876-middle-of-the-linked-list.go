/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
    // using fast and slow pointer
    // fast pointer move twice than slow pointer.
    // when fast pointer reach end, slow pointer is in the middle
    
    s, f := &ListNode{}, &ListNode{}
    s.Next = head
    f.Next = head
    
    for f != nil {
        s = s.Next
        if f.Next == nil {
            f = f.Next
        } else {
            f = f.Next.Next    
        }
    }
    
    return s
}