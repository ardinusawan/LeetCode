// calculation start from last 3 element
func minCostClimbingStairs(cost []int) int {
    for i:=len(cost)-3;i>=0;i--{
        if cost[i + 1] < cost[i + 2] {
            cost[i] += cost [i + 1]    
        } else {
            cost[i] += cost [i + 2]
        }
    }
    if cost[0] < cost[1] {
        return cost[0]
    }
    return cost[1]
}