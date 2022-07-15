
func getHint(secret string, guess string) string {
    bull, cow := 0, 0
    cache := make([]int, 10)
    for i := range secret {
        a, b := secret[i] - '0', guess[i] - '0'
        if a == b { bull++; continue }
        if cache[a] < 0 { cow++ }
        if cache[b] > 0 { cow++ }
        cache[a]++
        cache[b]--
    }
    return fmt.Sprintf("%dA%dB", bull, cow)
}