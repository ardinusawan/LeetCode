var stack []string
var result []string

func generateParenthesis(n int) []string {
    // only add open paranthesis if open < n
    // only add close paranthesis if close < open
    // valid if open == close == n

    stack = make([]string, 0) // reset
    result = make([]string, 0) // reset
    backtracking(n, 0, 0)
    return result
}

func backtracking(n, open, close int) {
    // to visualize
    // fmt.Println("------")
    // fmt.Println(n, open, close)
    // fmt.Println(stack)
    // fmt.Println("------")
    // fmt.Println("")

    if open == n && close == n {
        result = append(result, strings.Join(stack, ""))
        return
    }

    if open < n {
        stack = append(stack, "(")
        backtracking(n, open + 1, close)
        stack = stack[:len(stack)-1] // after stack is joined and to result, it need to find other combination if any
    }

    if close < open {
        stack = append(stack, ")")
        backtracking(n, open, close + 1)
        stack = stack[:len(stack)-1] // after stack is joined and to result, it need to find other combination if any
    }
}