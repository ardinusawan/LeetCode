// two pointer, fast and slow
// fast walk two node, slow one node
// when fast is nil, walk must be located in the middle
// from middle until end, reverse linked list
// when start - middle == middle - end then it polindrom
// note: is it ok Next link is disconnected from start - middle with middle - end (we goona have 2 linked list)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    // result := false
    // locate middle pointer
    fast, slow := head, head
    for fast != nil {
        fast = fast.Next
        if fast != nil {
            fast = fast.Next
        }
        
        slow = slow.Next
    }
    
    middle := reverseList(slow)
    start := head    
    for middle != nil {
        if start.Val != middle.Val {
            return false
        }
        start = start.Next
        middle = middle.Next
    }
    
    return true
}

func reverseList(head *ListNode) *ListNode {
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