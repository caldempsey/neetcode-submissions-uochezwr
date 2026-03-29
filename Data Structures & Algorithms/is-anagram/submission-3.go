func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false 
    }

    var freq = make(map[rune]int, len(s))
    for _, c := range s {
        freq[c]++
    }

    for _, c := range t {
        _, ok := freq[c] 
        if !ok {
            return false 
        }

        freq[c]-- 

        if freq[c] < 0 {
           return false  
        }
    }

    return true 
 }
