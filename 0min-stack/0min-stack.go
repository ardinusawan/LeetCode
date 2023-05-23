import (
    "math"
)

type MinStack struct {
    Values []int
    MinValue int
}


func Constructor() MinStack {
    return MinStack {
        Values: []int{},
        MinValue: math.MaxInt,
    }
}


func (this *MinStack) Push(val int)  {
    if this.MinValue >= val {
        this.Values = append(this.Values, this.MinValue)
        this.MinValue = val
    }
    this.Values = append(this.Values, val)
}


func (this *MinStack) Pop()  {
    length := len(this.Values)
    if length == 0 {
        return
    }
    top := this.Top()
    this.Values = this.Values[:length-1]

    if top == this.MinValue {
        if len(this.Values) > 0 {
            this.MinValue = this.Values[len(this.Values)-1]
            this.Values = this.Values[:len(this.Values)-1]
        } else {
            this.MinValue = math.MaxInt
        }
    }
}


func (this *MinStack) Top() int {
    length := len(this.Values)
    if length == 0 {
        return 0
    }
    return this.Values[length-1]
}


func (this *MinStack) GetMin() int {
    return this.MinValue
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */