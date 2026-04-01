func longestConsecutive(nums []int) int {
    set := make(map[int]bool)
    for _, num := range nums {
        set[num] = true
    }

    // build explicit "next" edges so that we can walk the sequence
    next := make(map[int]int)
    for num := range set {
        if set[num+1] {
            next[num] = num + 1
        }
    }

    // find chain heads (no predecessor)
    var heads []int
    for num := range set {
        if !set[num-1] {
            heads = append(heads, num)
        }
    }

    // walk each chain
    best := 0
    for _, head := range heads {
        length := 1
        curr := head
		// the 'trick' to this problem is it looks multiplicative
		// but its not, because walking each head is only as complex as however many heads there are
		// so its O(c₁) + O(c₂) + ... + O(cₕ) = O(c₁ + c₂ + ... + cₕ) = O(n) since its purely additive
        for dest, exists := next[curr]; exists; dest, exists = next[curr] {
            curr = dest
            length++
        }
        if length > best {
            best = length
        }
    }
    return best
}