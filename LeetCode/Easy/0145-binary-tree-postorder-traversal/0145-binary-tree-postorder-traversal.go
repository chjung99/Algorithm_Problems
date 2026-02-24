/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var order []int
    traversal(&order, root)
    return order
}

func traversal(order *[]int, cur *TreeNode) {
    if (cur == nil) {
        return
    }
    
    traversal(order, cur.Left)
    traversal(order, cur.Right)

    *order = append(*order, cur.Val)
}