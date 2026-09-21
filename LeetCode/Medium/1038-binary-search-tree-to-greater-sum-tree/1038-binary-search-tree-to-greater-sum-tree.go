/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
    var greaterSum func(cur *TreeNode)
    sum := 0

    greaterSum = func(cur *TreeNode) {
        if (cur == nil) {
            return
        }

        greaterSum(cur.Right)
        sum += cur.Val
        cur.Val = sum
        greaterSum(cur.Left)

    }
    greaterSum(root)

    return root
}
