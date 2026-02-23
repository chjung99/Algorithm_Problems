/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var order []int
    traversal(&order, root)
    return order
}

func traversal(order *[]int, cur *TreeNode) {
    if (cur == nil) {
        return
    }

    *order = append(*order, cur.Val)

    traversal(order, cur.Left)
    traversal(order, cur.Right)
}