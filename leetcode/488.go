package main

import "bytes"

// WRRBBW
// RB
// -1

// 不得不说，模拟谁都会，但是 ... 细节是魔鬼呀
// 题解是真的牛皮
// 分情况剪枝
// 栈维护新的情况

// TODO

type status struct {
	now    string
	remain string
}

func findMinStep(board string, hand string) int {
	// BFS
	queue := []*status{&status{
		now:    board,
		remain: hand,
	}}
	m := make(map[status]struct{})
	appendNew := func(now string, b byte) []string {
		var res []string

		// 是可以在最后一个地方插的
		for i := 0; i <= len(now); i++ {
			if i > 0 && now[i-1] == b { // 剪枝，前一个不能和当前的一样，即只在头部插入（因为是一样的）
				continue
			}
			if now[i] == b || (i > 0 && i < len(now) && now[i-1] == now[i] && now[i-1] != b) {
				var bt bytes.Buffer
				bt.WriteByte()
				if now[i] == b {

					for now[i] ==
				}
			}
		}

		return res
	}

	step := 1
	for len(queue) > 0 {
		lenq := len(queue)

		for i := 0; i < lenq; i++ {
			// generate next
			now, remain := queue[i].now, queue[i].remain
			for i := 0; i < len(remain); i++ {
				if i > 0 && remain[i] == remain[i-1] {
					continue
				}
				next := appendNew(now, remain[i])
				nextremain := remain[:i] + remain[i+1:]
				for k := range next {
					if next[k] == "" {
						return step
					}
					add := status{now: next[k], remain: nextremain}
					if _, ok := m[add]; !ok {
						queue = append(queue, &add)
					}
				}
			}
		}

		step++
		queue = queue[lenq:]
	}

	return -1
}
