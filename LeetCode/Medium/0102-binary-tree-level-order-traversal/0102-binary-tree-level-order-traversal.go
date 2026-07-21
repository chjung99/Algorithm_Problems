/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    order := make([][]int, 0)

    if (root != nil) {
        traverse(root, &order, 0)
    }
    
    return order
}

func traverse(node *TreeNode, order *[][]int, depth int) {
    if (len(*order) == depth) {
        (*order) = append(*order, []int{})
    }

    (*order)[depth] = append((*order)[depth], node.Val)

    if (node.Left != nil) {
        traverse(node.Left, order, depth+1)
    }

    if (node.Right != nil) {
        traverse(node.Right, order, depth+1)
    }

}