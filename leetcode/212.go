package main

// trie 树减支
// 应该多结合一下
// 写的比较丑

var dis = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var res []string
var now []byte
var visted [][]bool
var bo [][]byte
var lenc int
var lenr int

type Trie struct {
	child [26]*Trie // child
	flag bool // last
}

func Constructor() Trie {
	return Trie{
		child: [26]*Trie{},
		flag:  false,
	}
}

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

func (this *Trie) StartsWith(prefix string) (bool, bool) {
	lenw := len(prefix)
	p := this
	for i := 0; i < lenw; i++ {
		b := prefix[i] - 'a'
		if p.child[b] == nil {
			return false, false
		}
		p = p.child[b]
	}

	flag := p.flag
	if p.flag {
		p.flag = false
	}

	return true, flag
}

var trie Trie

func findWords(board [][]byte, words []string) []string {
	res = []string{}
	bo = board
	lenc = len(board)
	lenr = len(board[0])

	visted = make([][]bool, lenc)
	for i := 0; i < lenc; i ++ {
		visted[i] = make([]bool, lenr)
	}

	trie = Constructor()
	for _, v := range words {
		trie.Insert(v)
	}

	for i := 0; i < lenc; i ++ {
		for j := 0; j < lenr; j ++ {
			now = append(now, bo[i][j])
			dfs(i, j)
			visted[i][j] = false
			now = []byte{}
		}
	}

	return res
}

func dfs(sc, sr int) {
	if visted[sc][sr] {
		return
	}

	startWith, exist := trie.StartsWith(string(now))
	if !startWith {
		return
	}
	if exist {
		res = append(res, string(now))
	}

	visted[sc][sr] = true

	for i := 0; i < 4; i++ {
		nsr := sr + dis[i][0]
		nsc := sc + dis[i][1]
		if nsr > -1 && nsc > -1 && nsr < lenr && nsc < lenc && !visted[nsc][nsr] {
			now = append(now, bo[nsc][nsr])
			dfs(nsc, nsr)
			// rowback
			visted[nsc][nsr] = false
			now = now[:len(now)-1]
		}
	}
}

