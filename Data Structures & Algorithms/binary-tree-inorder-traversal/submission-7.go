
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	_, result := dfs(root, []int{})

	return result
}

func dfs(root *TreeNode, result []int) (*TreeNode, []int) {
	if root == nil {
		return root, result 
	}

	// In order traversal of a binary search is taking the result of the entire left subtree first, and descending all the way down to the bottom
	// Then for each node on the way back up that callstack we check out the right subtree in the binary tree
	// The relative depth for the first call will be at most 1.
	_, leftResult := dfs(root.Left, result)
	_, rightResult := dfs(root.Right, result)

	// On our way back up the callstack:

	// The first value will always be the first value from the left
	curr := root.Val
	// Combine the downstream left values to the result first, this is because the callstack first visits the entire left subtree
	// So we need to grep all of the values from that root node.
	// In the first call of the stack, there will be no nodes
	result = append(result, leftResult...)
	// Then add the current node
	result = append(result, curr)
	// Then add all of the right values after the current node.
	// Because we only visit the right hand side, after, visiting the parent, and the entire left hand side when performing the DFS.
	result = append(result, rightResult...)
	// Continue back up the callstack
	// If this is the first traversal, we had no left values, and no right values, and just the node, so we
	return root, result
}
