func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return dfs(root, []int{})
}

func dfs(root *TreeNode, acc []int) []int {
	if root == nil {
		return acc
	}

	leftResult := dfs(root.Left, acc)
	rightResult := dfs(root.Right, acc)

	// Grep the values from the left, empty list for first traversal
	acc = append(acc, leftResult...)
	// Add the current root node 
	acc = append(acc, root.Val)
	// Grep all the values from the right, empty list for first traversal
	acc = append(acc, rightResult...)

	return acc
}