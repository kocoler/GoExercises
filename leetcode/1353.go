package main

import "sort"

type item struct {
	start int
	end   int
}

type items []item

func (i items) Len() int {
	return len(i)
}

func (i items) Less(a, b int) bool {
	return i[a].end < i[b].end
}

func (i items) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func maxEvents(events [][]int) int {
	var eventsItems items
	lene := len(events)
	eventsItems = make(items, lene)
	m := 0
	for i := 0; i < lene; i++ {
		eventsItems[i] = item{
			start: events[i][0],
			end:   events[i][1],
		}
		if m < events[i][1] {
			m = events[i][1]
		}
	}
	sort.Sort(eventsItems)
	num := 0
	l := 0
	for i := 0; i <= m; i ++ {
		for j := l; j < lene; j ++ {
			if eventsItems[j].start <= i && eventsItems[j].end >= i {
				num ++
				l ++
				break
			}
		}
	}

	return num
}
