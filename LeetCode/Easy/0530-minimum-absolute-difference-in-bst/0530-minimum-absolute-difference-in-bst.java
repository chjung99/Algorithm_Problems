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
    public int getMinimumDifference(TreeNode root) {
        List<Integer> nums = new ArrayList<>();

        BFS(root, nums);
        nums.sort(Comparator.naturalOrder());
        System.out.print(nums);
        int ans = nums.get(1) - nums.get(0);
        
        for (int i = 0; i < nums.size() - 1; i++) {
            ans = Math.min(ans, nums.get(i+1)-nums.get(i));
        }
        return ans;
    }
    public void BFS(TreeNode node, List<Integer> nums) {
        Deque<TreeNode> q = new ArrayDeque<>();
        q.add(node);
        nums.add(node.val);

        while (!q.isEmpty()) {
            TreeNode cur = q.removeFirst();
            if (cur.left != null) {
                nums.add(cur.left.val);
                q.add(cur.left);
            }
            if (cur.right != null) {
                nums.add(cur.right.val);
                q.add(cur.right);
            }
        }
    }
}