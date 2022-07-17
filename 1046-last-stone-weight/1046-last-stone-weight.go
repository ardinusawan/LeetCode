import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's lenght, not just it contents
    *h = append(*h, x.(int))
}

func(h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

type MaxHeap struct {
    IntHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }

func lastStoneWeight(stones []int) int {
    h := &MaxHeap{stones}
    fmt.Println(h)
    heap.Init(h)
    for h.Len() > 1 {
        first := heap.Pop(h).(int)
        second := heap.Pop(h).(int)
        if first > second { 
            heap.Push(h, first - second)
        }
    }
    
    fmt.Println(h, stones)
    heap.Push(h, 0) // edge case: if heap is empty
    return h.IntHeap[0]
}