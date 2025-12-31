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
    boolean answer = true;
    public boolean isValidBST(TreeNode root) {

        inorderTraversal(root, "root", root);
        return answer;
    }

    public void inorderTraversal(TreeNode node, String direction, TreeNode parent) {
        if (node.left != null && node.right != null){
            if (node.left.val < node.val && node.val < node.right.val){
                if (direction.equals("left")){
                    if (node.right.val < parent.val){
                        ;
                    }else {
                        answer = false;
                        return;
                    }
                }
                if (direction.equals("right")){
                    if (node.left.val > parent.val){
                        ;
                    }else {
                        answer = false;
                        return;
                    }
                }
                inorderTraversal(node.left, "left", node);
                inorderTraversal(node.right, "right", node);

            } else {
                answer = false;
            }
            return;
        }

        if (node.left == null && node.right == null) {
            return;
        }

        if (node.left == null) {
            if (node.val < node.right.val){
                if (direction.equals("left")){
                    if (node.right.val < parent.val){
                        ;
                    }else {
                        answer = false;
                        return;
                    }
                }
                
                inorderTraversal(node.right, "right", node);
            }else {
                answer = false;
            }
            return;
        }

        if (node.right == null) {
            if (node.left.val < node.val){
                
                if (direction.equals("right")){
                    if (node.left.val > parent.val){
                        ;
                    }else {
                        answer = false;
                        return;
                    }
                }
                inorderTraversal(node.left, "left", node);
            } else {
                answer = false;
            }
            return;
        }
    }
}