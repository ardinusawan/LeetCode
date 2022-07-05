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
    var prev *ListNode
    curr := head
    
    for curr != nil {
        temp := curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
    }
    
    return prev
}