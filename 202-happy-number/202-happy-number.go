import (
    "fmt"
)

func isHappy(n int) bool {
    visited := make(map[int]bool)
    
    return checkVisit(visited, n)
}

func checkVisit(visited map[int]bool, n int) bool {
    if _, visit := visited[n]; !visit {
        visited[n] = true
        n := sumOfSquare(n)
        fmt.Println(n)
        if n == 1 {
            return true
        }
        
        result := checkVisit(visited, n)
        if result {
            return true
        }
    }
    return false
}

func sumOfSquare(n int) int {
    // modulus to get latest element
    // divider to remove last element
    
    var result int
    for n > 0 {
        lastElement := n % 10
        lastElement = lastElement * lastElement
        
        result += lastElement
        n = n / 10
        fmt.Println(lastElement, result, n)
    }
    return result
}