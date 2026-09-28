/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    sum := 0
    var traverse func(cur *TreeNode)

    traverse = func(cur *TreeNode) {
        if (cur == nil) {
            return
        }

        traverse(cur.Left)
        if (cur.Val >= low && cur.Val <= high) {
            sum += cur.Val
        }
        // fmt.Println(cur.Val)
        traverse(cur.Right)
    }

    traverse(root)
    return sum
}