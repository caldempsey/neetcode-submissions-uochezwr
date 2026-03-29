func hasDuplicate(nums []int) bool {
    set := make(map[int]struct{}, len(nums))
    
    for _, n := range nums {
        if _, ok := set[n] ; ok {
            return true 
        }

        set[n] = struct{}{}
    }

    return false
}
