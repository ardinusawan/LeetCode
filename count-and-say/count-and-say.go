// iterate for range n
// init base result to "1"
// start iteration from i=1 until i==n
// for every walk, store uniq value found until now.
// when previous value is different with current, then append result to uniq value found + previous uniq value.
// when hit end of result, do same thing like before

import (
    "strconv"
)

func countAndSay(n int) string {
    result := "1"
    for i:=1;i<n;i++ {
        var updatedResult string
        uniqVal := result[0]
        uniqValSum := 0
        for _, v := range result {
            if uniqVal != byte(v) { // when char different with previous char
                updatedResult += strconv.Itoa(uniqValSum) + string(uniqVal)
                uniqVal = byte(v)
                uniqValSum = 0
            }
            uniqValSum += 1
        }
        
        // when reach end of characters
        updatedResult += strconv.Itoa(uniqValSum) + string(uniqVal)
        result = updatedResult
    }
    return result
}

