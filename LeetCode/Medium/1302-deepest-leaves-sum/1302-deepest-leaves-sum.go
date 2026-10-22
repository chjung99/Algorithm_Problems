/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    depthSum := make([]int, 0)

    var dfs func(node *TreeNode, cur int)

    dfs = func(node *TreeNode, cur int) {
        if node == nil {
            return
        }

        if len(depthSum) < cur {
            depthSum = append(depthSum, 0)
        }

        depthSum[cur-1] += node.Val

        dfs(node.Left, cur + 1)
        dfs(node.Right, cur + 1)
    }

    dfs(root, 1)

    return depthSum[len(depthSum)-1]
}