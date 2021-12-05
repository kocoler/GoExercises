package main

// 倒序 Tire 树

type Node struct {
	word   [26]*Node
	isWord bool
}

type StreamChecker struct {
	root    *Node
	history []byte // max history, if overflow 3000, pure back to 2000
}

func Constructor(words []string) StreamChecker {
	root := &Node{} // alloc
	for _, v := range words {
		now := root
		for i := len(v) - 1; i >= 0; i-- {
			if next := now.word[v[i]-'a']; next == nil {
				now.word[v[i]-'a'] = &Node{}
			}
			now = now.word[v[i]-'a']
		}
		now.isWord = true
	}

	return StreamChecker{root: root}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.history = append(this.history, letter)
	if len(this.history) > 3000 {
		this.history = this.history[:2000]
	}
	now := this.root
	index := len(this.history) - 1
	for now != nil && index >= 0 && !now.isWord {
		now = now.word[this.history[index]-'a']
		index--
	}

	return now != nil && now.isWord
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
