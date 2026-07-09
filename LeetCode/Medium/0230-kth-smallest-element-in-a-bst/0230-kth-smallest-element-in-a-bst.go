/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    ret := -1
    idx := 0
    searchInorder(root, &idx, k, &ret)
    return ret
}

func searchInorder(node *TreeNode, idx *int, k int, ret *int) {

    if (node == nil) {
        return
    }

    searchInorder(node.Left, idx, k, ret)

    //visit
    (*idx) += 1
    // fmt.Println(node.Val, *idx)
    if (*idx == k) {
        *ret = node.Val
    }

    searchInorder(node.Right, idx, k, ret)
}