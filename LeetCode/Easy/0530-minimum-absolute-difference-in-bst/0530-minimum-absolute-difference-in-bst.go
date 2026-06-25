/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    ans := 100000
    nums := make([]int, 0)

    q := make([]*TreeNode, 0)
    q = append(q, root)
    nums = append(nums, root.Val)

    for (len(q) != 0) {
        cur := q[0]
        q = q[1:]

        if (cur.Left != nil) {
            q = append(q, cur.Left)
            nums = append(nums, cur.Left.Val)
        }

        if (cur.Right != nil) {
            q = append(q, cur.Right)
            nums = append(nums, cur.Right.Val)
        }
    }
    sort.Ints(nums)
    for i := 0; i < len(nums) - 1; i++ {
        ans = min(ans, nums[i+1] - nums[i])
    }
    return ans
}