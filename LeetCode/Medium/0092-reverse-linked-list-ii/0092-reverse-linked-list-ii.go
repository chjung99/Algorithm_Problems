/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    arr := make([]int, 0)
    ptr := head
    
    for ptr != nil {
        arr = append(arr, ptr.Val)
        ptr = ptr.Next
    }

    ptr = head
    idx := 0
    for ptr != nil {
        if (idx >= left - 1 && idx <= right - 1 ) {
            ptr.Val = arr[right - 1 - (idx - left + 1)]
        }
        idx += 1
        ptr = ptr.Next
    }
    
    return head
}