/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    cur := head
    prev := head
    for cur != nil {
        if (cur.Val == val) {
            next := findNext(cur.Next, val)
            prev.Next = next
            cur = next

        } else {
            prev = cur
            cur = cur.Next
        }

    }

    if (head != nil && head.Val == val) {
        head = head.Next
    }

    return head
}

func findNext(cur *ListNode, val int) *ListNode {
    for cur != nil {
        if (cur.Val != val) {
            break
        }
        cur = cur.Next
    }
    return cur
}