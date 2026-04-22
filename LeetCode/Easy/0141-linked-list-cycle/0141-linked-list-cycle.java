/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public boolean hasCycle(ListNode head) {
        int MARK = 100_000;
        while (head != null) {
            if (head.val == MARK) {
                return true;
            }
            head.val = MARK;
            head = head.next;
        }
        return false;
    }
}