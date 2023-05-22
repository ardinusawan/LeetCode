// create stack data structure with byte as value
// implement Push, Pop, IsEmpty
// when open bracket, push. 
// when pair is correct, pop.
// else, return false
// in the end, if stack not empty return false

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

        // below code is try to pop(). when it empty, nothing can be popped hence it must be fail
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