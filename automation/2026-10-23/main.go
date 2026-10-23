/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3 *ListNode
    var head *ListNode

    carry := 0
    x := 0
    for (true) {
        if (l1 == nil && l2 == nil) {
            break
        }
        x = 0
        if (l1 != nil) {
            x += l1.Val
            l1 = l1.Next
        }

        if (l2 != nil) {
            x += l2.Val
            l2 = l2.Next
        }

        x += carry
        carry = 0

        if (x / 10 >= 1) {
            carry = 1
            x = x % 10
        }

        if (l3 == nil) {
            l3 = &ListNode{x, nil}
            head = l3
        } else {
            l3.Next = &ListNode{x, nil}
            l3 = l3.Next
        }
    }
    if (carry != 0) {
        l3.Next = &ListNode{carry, nil}
        l3 = l3.Next
    }

    return head
}