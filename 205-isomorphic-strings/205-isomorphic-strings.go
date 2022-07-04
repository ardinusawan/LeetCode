func isIsomorphic(s string, t string) bool {
    m := make(map[string]string, len(s))
    mReverse := make(map[string]string, len(s))
    
    for i, v := range s {
        if _, ok := m[string(v)]; ok {
            if m[string(v)] != string(t[i]) {
                return false
            }
        } else {
            if _, ok := mReverse[string(t[i])]; ok {
                if mReverse[string(t[i])] != string(v) {
                    return false
                }
            }
            
            m[string(v)] = string(t[i])
            mReverse[string(t[i])] = string(string(v))
        }
    }
    return true
}