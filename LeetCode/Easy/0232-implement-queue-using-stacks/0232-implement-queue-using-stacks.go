type MyQueue struct {
    s []int 
}


func Constructor() MyQueue {
    return MyQueue{
        s : make([]int, 0),
    }
}


func (this *MyQueue) Push(x int)  {
    this.s = append(this.s, x)
}


func (this *MyQueue) Pop() int {
    front := this.s[0]
    this.s = this.s[1:]
    return front
}


func (this *MyQueue) Peek() int {
    return this.s[0]
}


func (this *MyQueue) Empty() bool {
    return len(this.s) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */