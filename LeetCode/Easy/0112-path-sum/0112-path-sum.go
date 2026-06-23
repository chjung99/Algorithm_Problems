/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    return hasPathSumWithLen(root, targetSum, 0)
}

func hasPathSumWithLen(root *TreeNode, targetSum int, size int) bool {
    if (root == nil) {
        return targetSum == 0 && size >= 1
    }
    if (root.Left != nil && root.Right == nil) {
        return hasPathSumWithLen(root.Left, targetSum - root.Val, size + 1)
    } else if (root.Left == nil && root.Right != nil) {
        return hasPathSumWithLen(root.Right, targetSum - root.Val, size + 1)
    } else {
        return hasPathSumWithLen(root.Left, targetSum - root.Val, size + 1) || hasPathSumWithLen(root.Right, targetSum - root.Val, size + 1)
    }
    
}