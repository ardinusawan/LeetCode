/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "fmt"

func reverseList(head *ListNode) *ListNode {
    // iterative
    // T O(n), M O(1)
    var prev *ListNode
    curr := head
    
    for curr != nil {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }
    
    return prev
    
    // recursive
    // T O(n), M O(n)
//     if head == nil {
//         return nil
//     }
    
//     newHead := head
//     if head.Next != nil {
//         newHead = reverseList(head.Next)
//         head.Next.Next = head
//     }
//     head.Next = nil
    
//     return newHead
}