/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type QueueNode struct {
    Node *TreeNode
    Depth int
 }

func averageOfLevels(root *TreeNode) []float64 {
    n := depthOfTreeNode(root)
    avgArr := make([]float64, n)
    cntArr := make([]int, n)

    q := make([]QueueNode, 0)

    q = append(q, QueueNode{root, 1})
    avgArr[0] = float64(root.Val)
    cntArr[0] = 1

    for (len(q) != 0) {
        cur := q[0]
        q = q[1:]

        sum := 0
        cnt := 0

        if (cur.Node.Left != nil) {
            sum += cur.Node.Left.Val
            cnt += 1
            q = append(q, QueueNode{cur.Node.Left, cur.Depth + 1})
        }

        if (cur.Node.Right != nil) {
            sum += cur.Node.Right.Val
            cnt += 1
            q = append(q, QueueNode{cur.Node.Right, cur.Depth + 1})
        }

        if (cnt >= 1){
            avgArr[cur.Depth] += float64(sum)
            cntArr[cur.Depth] += cnt
        }
    }

    for i := 0; i < n; i++ {
        avgArr[i] = avgArr[i] / float64(cntArr[i])
    }


    return avgArr
}

func depthOfTreeNode(root *TreeNode) int {
    if (root == nil) {
        return 0
    }
    return max(depthOfTreeNode(root.Left), depthOfTreeNode(root.Right)) + 1
}