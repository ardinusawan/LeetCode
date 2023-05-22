type Stack []byte

func (s *Stack) Push(value byte) {
    *s = append(*s, value)
}

func (s *Stack) Pop() (byte, bool) {
    stack := *s
    if len(stack) == 0 {
        return 0, false
    }
    index := len(stack)-1
    popped := stack[index]
    *s = stack[:index]
    return popped, true
}

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}

func isValid(s string) bool {
    stack := Stack{}
    for _, val := range s {
        v := byte(val)
        if v == '(' || v == '{' || v == '[' {
            stack.Push(v)
            continue
        }
        if stack.IsEmpty() {
            return false
        }
        
        if v == ')' && stack[len(stack)-1] == '(' {
            stack.Pop()
        } else if v == '}' && stack[len(stack)-1] == '{' {
            stack.Pop()
        } else if v == ']' && stack[len(stack)-1] == '[' {
            stack.Pop()
        } else {
            return false
        }
    }
    return stack.IsEmpty()
}