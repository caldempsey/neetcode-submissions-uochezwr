func twoSum(nums []int, target int) []int {
	var indexByComputedTarget = make(map[int]int, len(nums))

	for i, v := range nums {
		var computedTarget = target - v
		targetIdx, ok := indexByComputedTarget[v]
		if ok {
			return []int{targetIdx, i}
		}

		indexByComputedTarget[computedTarget] = i
	}

	return []int{-1, -1}
}
