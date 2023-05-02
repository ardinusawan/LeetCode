// make two pointer, pointerA = headA, pointerB = headB
// those two pointer walk one node everytime
// when node reach eol (end of linked list), it continue walk starting from their pair node head
// when it meet each other (in same node, or in end of linked list) return on of the pointer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    pointerA, pointerB := headA, headB
    for pointerA != pointerB {
        if pointerA != nil {
            pointerA = pointerA.Next    
        } else {
            pointerA = headB
        }
        
        if pointerB != nil {
            pointerB = pointerB.Next    
        } else {
            pointerB = headA
        }
    }
    return pointerA
}