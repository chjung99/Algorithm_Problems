/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    cnt := 0
    visitNode(&cnt, root)
    return cnt
}

func visitNode(cnt *int, cur *TreeNode) {
    if (cur != nil) {
        *cnt += 1
        visitNode(cnt, cur.Left)
        visitNode(cnt, cur.Right)
    }
    return
}