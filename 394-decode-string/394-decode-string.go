func decodeString(s string) string {
    var stack []string
	for _, v := range s {
		v := string(v)
		if v != string("]") { // push element to stack except close bracket
			stack = append(stack, v)
		} else {
            var substr string
            for string(stack[len(stack)-1]) != "[" { // substr = element inside []
                substr = fmt.Sprintf("%s%s", string(stack[len(stack)-1]), substr)
                stack = stack[:len(stack)-1]
            }
            stack = stack[:len(stack)-1] // pop "["

            var multiplier string
            for len(stack) > 0 { // get all number to multiply with substr
                if _, err := strconv.Atoi(string(stack[len(stack)-1])); err != nil {
                    break
                }
                multiplier = fmt.Sprintf("%s%s", string(stack[len(stack)-1]), multiplier)
                stack = stack[:len(stack)-1]
            }
            
            n, _ := strconv.Atoi(multiplier)
            for i:=0;i<n;i++ { // multiply number with substr
                for _, vs := range substr {
                    stack = append(stack, string(vs))
                }
            }
		}
	}
	return strings.Join(stack, "")

}