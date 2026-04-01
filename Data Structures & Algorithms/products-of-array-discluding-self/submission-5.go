func productExceptSelf(nums []int) []int {
	product := 1
	zeroCount := 0
	zeroIdx := -1

	for i, num := range nums {
		if num == 0 {
			zeroCount++
			zeroIdx = i
			continue 
		} 
		product *= num
	}

	out := make([]int, len(nums))

	if zeroCount > 1 {
		return out // all zeros already
	}
	if zeroCount == 1 {
		out[zeroIdx] = product
		return out
	}

	// no zeros
	for i := range out {
		out[i] = product / nums[i]
	}
	return out
}