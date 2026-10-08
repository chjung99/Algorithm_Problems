/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    prev := head
    sz := 0

    for (prev != nil) {
        prev = prev.Next
        sz += 1
    }

    prev = head

    for i := 0; i < sz - n - 1; i++ {
        prev = prev.Next
    }

    // prev -> ptr -> next
    if (sz - n - 1 < 0) {
        if (sz == 1) {
            head = nil
        } else {
            head = head.Next
        }
    } else {
        prev.Next = prev.Next.Next
    }
    // prev -> prev.next -> nil
    // prev -> prev.next -> prev.next.next -> nil
    return head
}