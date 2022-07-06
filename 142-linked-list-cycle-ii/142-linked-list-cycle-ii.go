/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "fmt"

func detectCycle(head *ListNode) *ListNode {
    // slow and fast pointer
    h := make(map[*ListNode]bool)
    
	slow := &ListNode{}
	fast := &ListNode{}

	slow.Next = head
	fast.Next = head

	for fast != nil && fast.Next != nil {
        if ok, _ := h[slow]; ok {
			return slow
		}
        
        h[slow] = true

		slow = slow.Next
		fast = fast.Next.Next
	}
    
	return nil
}