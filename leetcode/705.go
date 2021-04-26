package main

func main() {

}

//type HashNode struct {
//	val string
//	next *HashNode
//}

type HashNode int

type MyHashSet struct {
	val [10000][]HashNode
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{val: [10000][]HashNode{}}
}

func (this *MyHashSet) Add(key int)  {
	if this.Contains(key) {
		return
	}
	k := hash(key)
	this.val[k] = append(this.val[k], HashNode(k))
}


func (this *MyHashSet) Remove(key int)  {
	k := hash(key)
	t := this.val[k]
	l := len(t)
	for i := 0; i < l; i++ {
		if t[i] == HashNode(key) {
			this.val[k] = append(this.val[k][:i], this.val[k][i+1:] ...)
			return
		}
	}
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	k := hash(key)
	t := this.val[k]
	l := len(t)
	for i := 0; i < l; i++ {
		if t[i] == HashNode(key) {
			// this.val[k] = append(this.val[k][:i], this.val[k][i+1:] ...)
			return true
		}
	}

	return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func hash(key int) int {
	return key/4057
}
