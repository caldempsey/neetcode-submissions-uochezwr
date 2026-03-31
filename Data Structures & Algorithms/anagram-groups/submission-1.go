func groupAnagrams(strs []string) [][]string {
	// Create a map of sorted strings 

	groupBySortedString := make(map[string][]string, len(strs))
	
	// Dump all the strings into the map of sorted strings if they are matched when they are sorted 
	for _, str := range strs {
		sorted := sortString(str)
		groupBySortedString[sorted] = append(groupBySortedString[sorted], str)
	}


	// Output the answer as each sorted string and return the result
	ans := make([][]string, 0, len(strs))

	for _, v := range groupBySortedString{
		ans = append(ans, v)
	}
	
	return ans
}

func sortString(s string) string {
	// convert the string to a rune slice
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		// Sort characters ascending
		return r[i] < r[j]
	})
	return string(r)
}