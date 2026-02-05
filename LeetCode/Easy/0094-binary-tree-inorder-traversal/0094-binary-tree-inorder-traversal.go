/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var ret []int
    traversalNode(&ret, root)
    return ret
}

func traversalNode(ret *[]int, node *TreeNode) {
    if (node == nil) {
        return
    }
    if (node.Left != nil) {
        traversalNode(ret, node.Left)
    }
    //visit node
    *ret = append(*ret, node.Val)
    if (node.Right != nil) {
        traversalNode(ret, node.Right)
    }
}