package main

type HashNode struct {
	key int
	value int
}

type MyHashMap struct {
	val [4057][]HashNode
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{val: [4057][]HashNode{}}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	hkey := hash(key)
	lenk := len(this.val[hkey])
	for i:=0; i < lenk; i ++ {
		if this.val[hkey][i].key == key {
			this.val[hkey][i].value = value
			return
		}
	}

	this.val[hkey] = append(this.val[hkey], HashNode{key: key, value: value})
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hkey := hash(key)
	lenk := len(this.val[hkey])
	for i:=0; i < lenk; i ++ {
		if this.val[hkey][i].key == key {
			return this.val[hkey][i].value
		}
	}

	return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	hkey := hash(key)
	lenk := len(this.val[hkey])
	for i:=0; i < lenk; i ++ {
		if this.val[hkey][i].key == key {
			this.val[hkey] = append(this.val[hkey][:i], this.val[hkey][i+1:]...)
			return
		}
	}
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func hash(key int) int {
	return key/4057
}
