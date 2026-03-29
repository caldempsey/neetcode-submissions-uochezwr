func hasDuplicate(nums []int) bool {
    const size = 10007 // prime number for better distribution
    table := make([][]int, size)

    for _, num := range nums {
        idx := abs(num) % size
        // linear scan in bucket for collision resolution
        for _, v := range table[idx] {
            if v == num {
                return true
            }
        }
        table[idx] = append(table[idx], num)
    }
    return false
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}