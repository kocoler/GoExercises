package main

import (
	"fmt"
	"sort"
)

func main() {
	members := []int{1, 2, 3, 4, 10, 44}
	preMembers := []int{2, 4, 6, 8, 10}

	sort.Ints(members)
	sort.Ints(preMembers)
	lenNew := len(members)
	lenPre := len(preMembers)

	var outdated []int
	var added []int

	// two-pointer -> O(m+n)
	n := 0
	p := 0
	for n < lenNew && p < lenPre {
		if members[n] < preMembers[p] {
			added = append(added, members[n])
			n ++
		} else if members[n] > preMembers[p] {
			outdated = append(outdated, preMembers[p])
			p ++
		} else if members[n] == preMembers[p] {
			n++
			p++
		}
	}

	for p < lenPre {
		outdated = append(outdated, preMembers[p])
		p ++
	}

	for n < lenNew {
		added = append(added, members[n])
		n ++
	}

	fmt.Println(outdated, added)
}
