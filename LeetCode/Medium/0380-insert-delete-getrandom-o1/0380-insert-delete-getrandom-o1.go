type RandomizedSet struct {
    set map[int]int
    keys []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        set: make(map[int]int),
        keys: make([]int, 0),
        }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.set[val]; exists {
        return false
    }
    this.set[val] = len(this.keys)
    this.keys = append(this.keys, val)
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, exists := this.set[val]; !exists {
        return false
    }

    removeIdx := this.set[val]
    LastIdx := len(this.keys) - 1
    LastElem := this.keys[LastIdx]

    this.set[LastElem] = removeIdx
    this.keys[LastIdx], this.keys[removeIdx] = this.keys[removeIdx], this.keys[LastIdx]
    this.keys = this.keys[:LastIdx]

    delete(this.set, val)
    
    return true
}


func (this *RandomizedSet) GetRandom() int {
    randomIdx := rand.Intn(len(this.keys))
    return this.keys[randomIdx]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */