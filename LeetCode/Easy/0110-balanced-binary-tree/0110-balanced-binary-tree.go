/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if (root == nil) {
        return true
    }
    
    return abs(getHeight(root.Left) - getHeight(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(cur *TreeNode) int {
    if (cur == nil) {
        return 1
    }
    return max(getHeight(cur.Left) + 1, getHeight(cur.Right) + 1)
}

func abs(x int) int {
    if (x < 0) {
        return -x
    }
    return x
}