type MinStack struct {
    st []int
    minSt []int
    minValue int
}


func Constructor() MinStack {
    return MinStack{
        make([]int, 0),
        make([]int, 0),
        math.MaxInt,
    }
}


func (this *MinStack) Push(value int)  {
    if (value <= this.minValue) {
        this.minValue = value
        this.minSt = append(this.minSt, value)
    }
    this.st = append(this.st, value)
}


func (this *MinStack) Pop()  {
    if (this.Top() == this.minValue) {
        this.minSt = this.minSt[:len(this.minSt)-1]
    }
    this.st = this.st[:len(this.st)-1]

    if (len(this.minSt) != 0) {
        this.minValue = this.minSt[len(this.minSt)-1]
    } else {
        this.minValue = math.MaxInt
    }

}


func (this *MinStack) Top() int {
    return this.st[len(this.st)-1]
}


func (this *MinStack) GetMin() int {
    return this.minValue
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */