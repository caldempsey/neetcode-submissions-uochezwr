// sortKeysByFrequency returns keys as a tuple of sortOrder, key, sorted by frequency
func sortKeysByFrequency(freq map[int]int) [][]int {
	tuples := make([][]int, 0, len(freq))

	for k, v := range freq { 
		tuple := make([]int, 2)
		tuple[0], tuple[1] = k, v
		
		tuples = append(tuples, tuple)
	}

	sort.Slice(tuples, func(i, j int) bool {
		// ascending sort O(n log n) on the frequency
		return tuples[i][1] > tuples[j][1] 
	})
	
	return tuples
}



func topKFrequent(nums []int, k int) []int {
	var freq = make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
	}
	
	keysByFrequency := sortKeysByFrequency(freq)
	kFrequent := keysByFrequency[:k]

	ans := make([]int, 0, len(kFrequent)) 
	
	for _, t := range kFrequent {
		ans = append(ans, t[0])
	}
	
	return ans
}
