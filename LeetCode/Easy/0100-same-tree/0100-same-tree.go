/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    pVals := make([]int, 0)
    qVals := make([]int, 0)

    bfs(p, &pVals)
    bfs(q, &qVals)

    pLen := len(pVals)
    qLen := len(qVals)

    fmt.Println(pVals)
    fmt.Println(qVals)
    if (pLen != qLen) {
        return false
    }

    for i := 0; i < pLen; i++ {
        if (pVals[i] != qVals[i]) {
            return false
        }
    }

    return true
}

func bfs(p *TreeNode, pVals *[]int) {
    st := make([]*TreeNode, 0)
    st = append(st, p)
    if (p!= nil){
        *pVals = append(*pVals, p.Val)
    } else {
        *pVals = append(*pVals, 100000)
    }
    

    for (len(st) != 0) {
        cur := st[len(st)-1]
        st = st[:len(st)-1]

        if (cur == nil) {
            continue
        }

        if (cur.Left != nil) {
            st = append(st, cur.Left)
            *pVals = append(*pVals, cur.Left.Val)
        } else {
            st = append(st, cur.Left)
            *pVals = append(*pVals, 100000)
        }

        if (cur.Right != nil) {
            st = append(st, cur.Right)
            *pVals = append(*pVals, cur.Right.Val)
        } else {
            st = append(st, cur.Right)
            *pVals = append(*pVals, 100000)
        }
    }
}