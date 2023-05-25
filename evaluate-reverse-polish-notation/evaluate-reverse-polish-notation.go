// ["2","1","+","3","*"]
// you can see the pattern, when it hit '+', '-', '*', '/' calculate last two data
// else push the value to the stack

import (
	"strconv"
)

func evalRPN(tokens []string) int {
    var stack []int
    for i, _ := range tokens {
        token := string(tokens[i])
        if token == "+" || token == "-" || token == "*" || token == "/" {
            secondVal := int(stack[len(stack)-1])
            firstVal := int(stack[len(stack)-2])
            var result int
            switch (token) {
                case "+": result = firstVal + secondVal;
                case "-": result = firstVal - secondVal;
                case "*": result = firstVal * secondVal;
                case "/": result = firstVal / secondVal;
            }            
            stack = stack[:len(stack)-2]
            stack = append(stack, result)
        } else {
            num, _ := strconv.Atoi(tokens[i])
            stack = append(stack, num)
        }
    }
    return stack[0]
}