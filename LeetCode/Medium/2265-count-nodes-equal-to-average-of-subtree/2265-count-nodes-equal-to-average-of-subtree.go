/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
    var sumOfSubtree func(node *TreeNode) (int, int)
    cnt := 0
    
    sumOfSubtree = func(node *TreeNode) (int, int) {
        if (node == nil) {
            return 0, 0
        }

        leftSum, leftCnt := sumOfSubtree(node.Left)
        rightSum, rightCnt := sumOfSubtree(node.Right)

        totalSum := leftSum + rightSum + node.Val
        totalCnt := leftCnt + rightCnt + 1

        if (node.Val == int(totalSum / totalCnt) ) {
            cnt += 1
        }

        return totalSum, totalCnt
    }

    sumOfSubtree(root)

    return cnt
}