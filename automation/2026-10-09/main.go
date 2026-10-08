/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    ptr1 := head
    ptr2 := head

    for (ptr1 != nil && ptr2 != nil) {
        ptr1 = ptr1.Next
        ptr2 = ptr2.Next

        if (ptr2 == nil) {
            break
        }

        ptr2 = ptr2.Next

        if (ptr1 == ptr2) {
            return true
        }
    }
    return false;
}