package main

import "fmt"

type loserTree struct {
	size  int     // lists 的大小
	tree  []int   // 每个父节点记录的是败者，0 是胜者
	lists [][]int // 多路归并
}

func newLoserTree(lists [][]int) *loserTree {
	return &loserTree{
		size:  len(lists),
		tree:  make([]int, len(lists)*2-1),
		lists: lists,
	}
}

func (l *loserTree) initLoserTree() {
	l.initWinner(0)
}

func (l *loserTree) initWinner(index int) int {
	if index == 0 {
		l.tree[0] = l.initWinner(index*2 + 1)
		return l.tree[0]
	}
	if index >= l.size {
		return index - l.size
	}

	// 递归计算左右子树的胜者
	winnerLeft, winnerRight := l.initWinner(index*2), l.initWinner(index*2+1)
	// 处理一下使 left 表示败者
	if l.lists[winnerLeft] == nil {
		// 左边没东西啦
		// 胜利的肯定是右边
		// 并且每次的败者肯定都是这边，右边一定会浮上去
	} else if (len(l.lists[winnerRight]) == 0) || (l.lists[winnerLeft][0] < l.lists[winnerRight][0]) {
		winnerLeft, winnerRight = winnerRight, winnerLeft
	}
	l.tree[index] = winnerLeft

	// 返回胜者
	return winnerRight
}

// and refresh
func (l *loserTree) winnerGet() *int {
	// 取出胜者，如果胜者都是空的了，那肯定是空的
	winnerList := l.lists[l.tree[0]]
	if len(winnerList) == 0 {
		return nil
	}

	// 取出胜者
	winner := winnerList[0]
	// 删除胜者
	l.lists[l.tree[0]] = l.lists[l.tree[0]][1:]

	// 刷新胜者
	//l.initLoserTree()
	// treeIndex 表示胜者的父节点
	treeIndex := (l.tree[0] + l.size) / 2
	treeWinner := l.tree[0]
	for treeIndex > 0 {
		// 调整 使 treeWinner 表示胜者
		// 这个节点的上一个败者，因为要更新，所以可以只从有更新的那个地方的败者树上去就可以
		treeLoser := l.tree[treeIndex]
		if len(l.lists[treeLoser]) == 0 {
			// 这个败者没东西了，说明胜者还是出自上一个的那个 list
		} else if (len(l.lists[treeWinner]) == 0) || (l.lists[treeLoser][0] < l.lists[treeWinner][0]) {
			// 找到了新的胜者，更新 treeWinner 使其一直代表胜者
			treeLoser, treeWinner = treeWinner, treeLoser
		}

		// treeLoser 在这一轮败了
		l.tree[treeIndex] = treeLoser
		// 继续向上比赛
		treeIndex /= 2
	}
	// 更新新的 winner
	l.tree[0] = treeWinner

	return &winner
}

func main() {
	lists := [][]int{
		[]int{5, 6, 7},
		[]int{1, 2, 3},
		[]int{1, 4, 8},
	}
	l := newLoserTree(lists)
	l.initLoserTree()
	for winner := l.winnerGet(); winner != nil; winner = l.winnerGet() {
		fmt.Println(*winner)
	}
}
