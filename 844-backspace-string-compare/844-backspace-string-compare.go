import "fmt"

func backspaceCompare(s string, t string) bool {
    // stack
    var ss, st []string
    for _, v := range s {
        if string(v) != "#" {
            ss = append(ss, string(v)) // push
        } else if len(ss) > 0 {
            ss = ss[:len(ss)-1] // pop
        }
    }
    
    for _, v := range t {
        if string(v) != "#" {
            st = append(st, string(v)) // push
        } else if len(st) > 0 {
            st = st[:len(st)-1] // pop
        }
    }
    
    fmt.Println(ss, st)
    if len(ss) != len(st) { return false }
    for i, _ := range ss {
        if ss[i] != st[i] { return false }
    }
    
    return true
}