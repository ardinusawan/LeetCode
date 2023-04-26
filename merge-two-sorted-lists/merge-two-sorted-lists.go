/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    result := &ListNode{Val: 0}
    head := result
    
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            head.Next = list1
            list1 = list1.Next
        } else {
            head.Next = list2
            list2 = list2.Next
        }
        head = head.Next
    }
    
    if list1 != nil {
        head.Next = list1
    } else {
        head.Next = list2
    }
    
    return result.Next
}