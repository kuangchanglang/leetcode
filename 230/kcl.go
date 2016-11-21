/**
 * Detion for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    if k <= 0 {
        return 1
    }

    cnt := count(root.Left)
	if cnt > k-1{
        return kthSmallest(root.Left, k)
    }
    
    if cnt == k-1{
        return root.Val
    }
    
    return kthSmallest(root.Right, k-1-cnt)
}

func count(root *TreeNode) int{
	if root == nil{
		return 0
	}
	return count(root.Left) + count(root.Right) + 1
}
