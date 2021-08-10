package lcof

// 总之就是 append 开销太大，要第一次统计出来长度开辟空间比较好

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	root := head

	count := 0
	for head != nil {
		count ++
		head = head.Next
	}

	res := make([]int, count)
	head = root

	for head != nil {
		count --
		res[count] = head.Val
		head = head.Next
	}

	return res
}
