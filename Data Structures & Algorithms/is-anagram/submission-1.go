func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false 
    }

    var m = make(map[rune]int, len(s))
    for _, c := range s {
        m[c]++
    }

    for _, t := range t {
        if _, ok := m[t]; !ok {
            return false 
        }
        if _, ok := m[t]; ok {
            m[t]-- 
        }
        if m[t] < 0 {
           return false  
        }
    }

    for _, v := range m {
        if v != 0 {
            return false
        }
    }
     
    return true 
}
