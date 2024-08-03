func winningPlayerCount(n int, pick [][]int) int {
    winner := 0
    for i:=0;i<n;i++{
        picked := make(map[int]int, len(pick))
        for j:=0;j<len(pick);j++ {
            if pick[j][0] == i {
                picked[pick[j][1]] += 1
            }
        }
        for _, v := range picked {
            if v >= i + 1 {
                winner += 1
                break
            }
        }
    }   
    return winner
}