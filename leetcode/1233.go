package main

import (
	"fmt"
	"sort"
)

// trie 树
// 这道题的出现，让我们思考 trie 树的应用场景就很有意义
// 字典树 和 暴力 的区别？
// 字典树 快速检索（最长前缀匹配），统计，排序和保存大量的字符串，最坏是 O(n)
// 前缀？

// 本来写的是先插入再算，现在就是先排序，然后在插入时算

// 这道题更简单的做法是 排序 + HasPrefix ...

func main() {
	fmt.Println(removeSubfolders([]string{"/a/b/c", "/a/b/d", "/a/b/ca"}))
}

type trieNode struct {
	child [27]*trieNode
	root  bool
}

func (t *trieNode) Add(str string) bool {
	root := t
	str = str + "/"
	last := ' '
	for _, v := range str {
		sym := t.getSym(v)
		if root.root && last == '/' {
			return true
		}
		if root.child[sym] == nil {
			root.child[sym] = &trieNode{}
		}
		root = root.child[sym]
		last = v
	}
	root.root = true

	return false
}

func (t *trieNode) getSym(sym int32) int {
	if sym <= 'z' && sym >= 'a' {
		sym -= 'a'
	} else {
		sym = 26
	}
	return int(sym)
}

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	tire := &trieNode{}

	var res []string
	for _, v := range folder {
		if !tire.Add(v) {
			res = append(res, v)
		}
	}
	return res
}
