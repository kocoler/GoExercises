package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func buildList(nums []int) *ListNode {
	var res *ListNode

	for i := len(nums) - 1; i >= 0; i-- {
		node := &ListNode{Val: nums[i]}
		node.Next, res = res, node
	}

	return res
}

func printList(node *ListNode) {
	h := node
	for h != nil {
		fmt.Printf("%d --> ", h.Val)
		h = h.Next
	}
	fmt.Println("nil")
}

func main() {
	list := buildList([]int{0, 4, 2, 1, 3})
	printList(list)
	res := reverseEvenLengthGroups(list)
	printList(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(list *ListNode, aimHead *ListNode, num int) *ListNode {
	if list == nil {
		return nil
	}

	// 7 - 3 - 8 - 4
	// 4 - 8 - 3 - 7
	head := list
	next := list.Next
	tail := list
	for num != 1 {
		tail, next, next.Next = next, next.Next, tail
		num--
	}

	head.Next = aimHead
	return tail
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	num := 0
	//reverse := false
	aimNum := 1
	h := &ListNode{Next: head}
	preH := h
	for num != aimNum {
		nextH := preH.Next

		for nextH != nil && num != aimNum {
			nextH = nextH.Next

			num++
		}

		if num%2 == 0 {
			preH.Next = reverseList(preH.Next, nextH, num)
		}
		if num != aimNum || nextH == nil {
			break
		}
		nextPreH := preH
		for num > 0 {
			nextPreH = nextPreH.Next
			num--
		}

		preH = nextPreH
		//reverse = !reverse
		num = 0
		aimNum = aimNum + 1
	}

	return h.Next
}
