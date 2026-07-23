type Node struct {
    key int
    val int
    prev *Node
    next *Node
}

func initNode(key int, val int) *Node {
    return &Node{key, val, nil, nil}
}

func pushNode(tail *Node, node *Node){
    tmp := tail.prev

    tmp.next = node
    node.prev = tmp

    node.next = tail
    tail.prev = node
}

func popNode(node *Node){
    prevNode := node.prev
    nextNode := node.next

    prevNode.next = nextNode
    nextNode.prev = prevNode

    node.prev = nil
    node.next = nil
}

type LRUCache struct {
    head *Node
    tail *Node
    capacity int
    m map[int]*Node
    size int
}


func Constructor(capacity int) LRUCache {
    head := initNode(-1, -1)
    tail := initNode(-1, -1)

    head.next = tail
    tail.prev = head
    return LRUCache{head, tail, capacity, make(map[int]*Node), 0}
}


func (this *LRUCache) Get(key int) int {
    ptr, exists := (*this).m[key]
    if (!exists) {
        return -1
    }
    popNode(ptr)
    pushNode((*this).tail, ptr)
    return (*ptr).val
}


func (this *LRUCache) Put(key int, value int)  {
    ptr, exists := (*this).m[key]
    if (!exists) {
        node := initNode(key, value)
        (*this).m[key] = node
        if ((*this).size == (*this).capacity) {
            delete((*this).m, (*this).head.next.key)
            popNode((*this).head.next)
            (*this).size -= 1

        }
        pushNode((*this).tail, node)
        (*this).size += 1
    } else {
        popNode(ptr)
        pushNode((*this).tail, ptr)
        ptr.val = value
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */