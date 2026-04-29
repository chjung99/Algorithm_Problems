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
    public int countNodes(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int leftDepth = getDepth(root.left);
        int rightDepth = getDepth(root.right);

        if (leftDepth == rightDepth) {
            return (1 << leftDepth) - 1 + countNodes(root.right) + 1;
        }

        return (1 << rightDepth) - 1 + countNodes(root.left) + 1;
    }

    public int getDepth(TreeNode node) {
        if (node == null) {
            return 0;
        }
        return 1 + getDepth(node.left);
    }
}