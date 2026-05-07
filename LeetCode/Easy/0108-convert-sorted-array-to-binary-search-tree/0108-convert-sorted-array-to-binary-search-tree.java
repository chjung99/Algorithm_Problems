/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public TreeNode sortedArrayToBST(int[] nums) {
        int n = nums.length;
        return divdeAndConquer(0, n-1, nums);
    }
    public TreeNode divdeAndConquer(int left, int right, int[] nums) {
        if (left > right) {
            return null;
        }
        int mid = (left + right) / 2; // idx: 2
        TreeNode node = new TreeNode(nums[mid]); // nums[2] = 0
        node.left = divdeAndConquer(left, mid - 1, nums); // idx: 0
        node.right = divdeAndConquer(mid + 1, right, nums); // idx: 3
        
        return node;
    }
}