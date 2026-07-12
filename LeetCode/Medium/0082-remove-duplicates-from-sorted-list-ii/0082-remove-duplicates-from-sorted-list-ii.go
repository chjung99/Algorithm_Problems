/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    m := make(map[int]int)

    ptr := head

    for (ptr != nil) {
        _, exists := m[ptr.Val]
        if (!exists) {
            m[ptr.Val] = 0
        }
        m[ptr.Val] += 1
        ptr = ptr.Next
    }

    ptr = head

    if (ptr != nil && m[ptr.Val] >= 2) {
        head = findNext(ptr, m)
        ptr = head
    }

    for (ptr != nil && ptr.Next != nil) {
        val, _ := m[ptr.Next.Val]
        if (val >= 2) {
            nxt := findNext(ptr.Next, m)
            ptr.Next = nxt
        }
        ptr = ptr.Next
    }

    return head
}

func findNext(ptr *ListNode, m map[int]int) *ListNode {
    for (ptr != nil) {
        if (m[ptr.Val] == 1) {
            break
        }
        ptr = ptr.Next
    }
    return ptr
}