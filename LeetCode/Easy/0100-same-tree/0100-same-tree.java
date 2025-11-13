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
    public boolean isSameTree(TreeNode p, TreeNode q) {
        Queue<Integer> OrderP = new ArrayDeque<>();
        Queue<Integer> OrderQ = new ArrayDeque<>();
        trevasalTree(p, OrderP);
        trevasalTree(q, OrderQ);

        if (OrderP.size() != OrderQ.size()) return false;

        while (!OrderP.isEmpty() && !OrderQ.isEmpty()){ 
            if (!OrderP.poll().equals(OrderQ.poll())){
                return false;
            }
        }

        return true;
    }
    public void trevasalTree(TreeNode t, Queue<Integer> queue){
        // System.out.print(t.val+ " ");
        if (t != null){
            queue.add(t.val);
        } else {
            return;
        }

        if (t.left != null){
            trevasalTree(t.left, queue);
        } else {
            queue.add(100000);
        }
        
        if (t.right != null){
            trevasalTree(t.right, queue);
        } else {
            queue.add(100000);
        }
    }
}