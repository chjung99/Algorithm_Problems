/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    ptr := head

    for (ptr.Next != nil) {
        nxt := ptr.Next
        gcd := getGCD(ptr.Val, nxt.Val)

        ptr.Next = &ListNode{gcd, nxt}
        ptr = nxt
    }

    return head
}

func getGCD(x int, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}