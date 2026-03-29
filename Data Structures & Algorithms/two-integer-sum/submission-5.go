func twoSum(nums []int, target int) []int {
    needed := make(map[int]int, len(nums))  // what I need -> my index
    
    for i, num := range nums {
        if partnerIdx, ok := needed[num]; ok {
            return []int{partnerIdx, i}
        }
        needed[target - num] = i  // Store what I'm looking for
    }
    return nil
}
