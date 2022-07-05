/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    mergedList := ListNode{}
    tail := &mergedList
    
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else if list1.Val > list2.Val {
            tail.Next = list2
            list2 = list2.Next
        }
        tail = tail.Next
    }
    
    for list1 != nil && list2 == nil {
            tail.Next = list1
            list1 = list1.Next
            tail = tail.Next
    }
    
    for list1 == nil && list2 != nil {        
        tail.Next = list2
        list2 = list2.Next
        tail = tail.Next
    }
    
    
    return mergedList.Next
}