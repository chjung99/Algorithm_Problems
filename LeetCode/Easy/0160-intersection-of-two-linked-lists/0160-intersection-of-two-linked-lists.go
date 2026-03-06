/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var ans *ListNode

    var stackA []*ListNode
    var stackB []*ListNode

    for headA != nil {
        stackA = append(stackA, headA)
        headA = headA.Next
    }

    for headB != nil {
        stackB = append(stackB, headB)
        headB = headB.Next
    }

    for len(stackA) != 0 && len(stackB) != 0 {
        a := stackA[len(stackA) - 1]
        stackA = stackA[:len(stackA) - 1]

        b := stackB[len(stackB) - 1]
        stackB = stackB[:len(stackB) - 1]

        if (a == b) {
            ans = a
        }
    }

    return ans
}