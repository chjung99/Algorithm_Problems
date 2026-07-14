/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if (len(preorder) == 0 || len(inorder) == 0) {
        return nil
    }

    inMap := make(map[int]int)
    for i, val := range inorder {
        inMap[val] = i
    }

    var helper func(preStart, inStart, inEnd int) *TreeNode
    helper = func(preStart, inStart, inEnd int) *TreeNode {
        if inStart > inEnd {
            return nil
        }

        rootVal := preorder[preStart]
        root := &TreeNode{Val: rootVal}

        inRootIdx := inMap[rootVal]

        leftTreeSize := inRootIdx - inStart

        root.Left = helper(preStart + 1, inStart, inRootIdx - 1)

        root.Right = helper(preStart + leftTreeSize + 1, inRootIdx + 1, inEnd)

        return root
    }

    return helper(0, 0, len(inorder)-1)
}

