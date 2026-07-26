/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if (head == nil || head.Next == nil) {
        return head
    }

    var prev *ListNode
    slow := head
    fast := head

    for (fast != nil && fast.Next != nil) {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    prev.Next = nil

    left := sortList(head)
    right := sortList(slow)

    return merge(left, right)
}

func merge(left *ListNode, right *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy

    for (left != nil && right != nil) {
        if (left.Val <= right.Val) {
            curr.Next = left
            left = left.Next
        } else {
            curr.Next = right
            right = right.Next
        }
        curr = curr.Next
    }
    if (left != nil) {
        curr.Next = left
    }

    if (right != nil) {
        curr.Next = right
    }

    return dummy.Next
}