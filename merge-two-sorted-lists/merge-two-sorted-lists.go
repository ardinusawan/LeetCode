// init result with empty node
// iterate list1 and list2 using Next pointer
// when one of it nil already, next is rest of the other one
// use pointer to move between node
// result.Next is the result

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    result := &ListNode{Val: 0}
    pointer := result
    
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            pointer.Next = list1
            list1 = list1.Next
        } else {
            pointer.Next = list2
            list2 = list2.Next
        }
        pointer = pointer.Next
    }
    
    if list1 != nil {
        pointer.Next = list1
    } else {
        pointer.Next = list2
    }
    
    return result.Next
}