/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    var ans []string
    findPath(root, &ans, strconv.Itoa(root.Val))
    
    return ans
}

func findPath(cur *TreeNode, ans *[]string, path string) {
    if (cur.Left == nil && cur.Right == nil) {
        *ans = append(*ans, path)
        return
    }
    if (cur.Left != nil) {
        findPath(cur.Left, ans, path + "->" + strconv.Itoa(cur.Left.Val))
    }
    if (cur.Right != nil) {
        findPath(cur.Right, ans, path + "->" + strconv.Itoa(cur.Right.Val))
    }
    
}