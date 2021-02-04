package main

type Trie struct {
	child [26]*Trie // child
	flag bool // last
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		child: [26]*Trie{},
		flag:  false,
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	lenw := len(word)
	p := this
	for i := 0; i < lenw; i++ {
		b := word[i] - 'a'
		if p.child[b] == nil {
			p.child[b] = &Trie{
				child: [26]*Trie{},
				flag:  false,
			}
		}
		p = p.child[b]
	}
	p.flag = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	lenw := len(word)
	p := this
	for i := 0; i < lenw; i++ {
		b := word[i] - 'a'
		if p.child[b] == nil {
			return false
		}
		p = p.child[b]
	}

	return p.flag
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	lenw := len(prefix)
	p := this
	for i := 0; i < lenw; i++ {
		b := prefix[i] - 'a'
		if p.child[b] == nil {
			return false
		}
		p = p.child[b]
	}

	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
