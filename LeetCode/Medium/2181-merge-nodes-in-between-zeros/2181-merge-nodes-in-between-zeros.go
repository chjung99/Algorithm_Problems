/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    merged := &ListNode{-1, nil}
    ptr := merged
    var tmp *ListNode
    for (head != nil) {
        if (head.Val == 0) {
            if (tmp != nil) {
                ptr.Next = tmp
                ptr = ptr.Next
            }
            tmp = &ListNode{0, nil}
        } else {
            tmp.Val += head.Val
        }
        head = head.Next
    }
    return merged.Next
}