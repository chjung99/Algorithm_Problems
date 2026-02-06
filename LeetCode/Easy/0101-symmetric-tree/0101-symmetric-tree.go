/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    var leftSide []int
    var rightSide []int

    inOrderLeftSide(&leftSide, root, root, 0) 
    inOrderRightSide(&rightSide, root, root, 0)
    fmt.Println(leftSide)
    fmt.Println(rightSide)
    return slices.Equal(leftSide, rightSide)
}

func inOrderLeftSide(arr *[]int, node *TreeNode, root *TreeNode, cnt int) {
    if (node == nil) {
        return
    }

    inOrderLeftSide(arr, node.Left, root, cnt+1)
    // visit
    *arr = append(*arr, cnt)
    *arr = append(*arr, node.Val)
    if (node == root) {
        return
    }
    
    inOrderLeftSide(arr, node.Right, root, cnt+1)
}

func inOrderRightSide(arr *[]int, node *TreeNode, root *TreeNode, cnt int) {
    if (node == nil) {
        return
    }
    inOrderRightSide(arr, node.Right, root, cnt+1)
    // visit
    *arr = append(*arr, cnt)
    *arr = append(*arr, node.Val)
    if (node == root) {
        return
    }
    inOrderRightSide(arr, node.Left, root, cnt+1)

}