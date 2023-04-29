// create additional 3 pointer
// odd -> head, even -> head.Next, eventHead = even
// while even != nil && even.Next != nil (in case total of linked list is even)
// walk the node: odd.Next = even.Next, odd = odd.Next. do same with event
// in the end, last node of odd.Next = evenHead
// return head

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if (head == nil) { return head }
    odd := head

    even := head.Next
    evenHead := even

    for even != nil && even.Next != nil {
        odd.Next = even.Next;
        odd = odd.Next;
        even.Next = odd.Next;
        even = even.Next;
    }
    odd.Next = evenHead;
    return head
}