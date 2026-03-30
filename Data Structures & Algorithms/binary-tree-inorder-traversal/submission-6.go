/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
        return []int{root.Val}
    }

    left := inorderTraversal(root.Left)
    right := inorderTraversal(root.Right)
	// The first value will always be the first value from the left
	curr := root.Val 
	// We always scan the left hand side first so initialise an array as our answer
    result := []int{}
	// Combine the downstream left value to the result first
    result = append(result, left...)
	// Then add the current node
    result = append(result, curr)
	// Then add the right values 
    result = append(result, right...)
	// Now we have 4, 2, 5
	// Continue back up the callstack 
    return result
}
