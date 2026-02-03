/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    visit := make(map[int]bool)
    cur := head
    var prev *ListNode
    for cur != nil {
        _, exists := visit[cur.Val]
        
        if !exists {
            visit[cur.Val] = true
            prev = cur

        } else {
            prev.Next = cur.Next
        }

        cur = cur.Next
    }
    return head
}