/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var adder int
    result := &ListNode{Val: 0}
    head := result
    for l1 != nil || l2 != nil {
        var l1Val, l2Val int
        if l1 != nil {
            l1Val = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            l2Val = l2.Val
            l2 = l2.Next
        }
        sum := l1Val + l2Val + adder
        lastDigit := sum % 10
        adder = sum / 10
        
        head.Next = &ListNode{Val: lastDigit}
        head = head.Next
    }
    if adder > 0 {
        head.Next = &ListNode{Val: adder}
    }
    return result.Next
}