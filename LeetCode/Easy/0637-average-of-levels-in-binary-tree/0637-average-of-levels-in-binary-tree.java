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
    static class Node {
        TreeNode treenode;
        int depth;
        Node (TreeNode treenode, int depth) {
            this.treenode = treenode;
            this.depth = depth;
        }
    }
    public List<Double> averageOfLevels(TreeNode root) {
        ArrayDeque<Node> q = new ArrayDeque<>();
        List<Double> ans = new ArrayList<>();
        double[] sumOfLevels = new double[10000];
        int[] countOfLevels = new int[10000];
        int maxDepth = 0;

        q.add(new Node(root, 0));
        sumOfLevels[0] = root.val;
        countOfLevels[0] = 1;

        
        while (!q.isEmpty()) {
            Node cur = q.removeFirst();
            if (cur.treenode.left != null) {
                q.add(new Node(cur.treenode.left, cur.depth + 1));
                sumOfLevels[cur.depth + 1] += cur.treenode.left.val;
                countOfLevels[cur.depth + 1] += 1;
                maxDepth = Math.max(maxDepth, cur.depth + 1);
            }
            if (cur.treenode.right != null) {
                q.add(new Node(cur.treenode.right, cur.depth + 1));
                sumOfLevels[cur.depth + 1] += cur.treenode.right.val;
                countOfLevels[cur.depth + 1] += 1;
                maxDepth = Math.max(maxDepth, cur.depth + 1);
            }
        }

        for (int i = 0; i <= maxDepth; i++) {
            ans.add(Double.valueOf(sumOfLevels[i]/countOfLevels[i]));
        }

        return ans;
    }
}