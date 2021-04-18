package main

type BSTNode struct {
	val int
	left *BSTNode
	right *BSTNode
}

//     root
//  left   right
// left < root < right

func search(v int, head *BSTNode) *BSTNode {
	t := head
	for t != nil {
		if t.val == v {
			return t
		} else if t.val > v {
			t = t.left
		} else {
			t = t.right
		}
	}

	return nil
}

func insert(v int, head *BSTNode) *BSTNode {
	for t := head; t != nil; {
		if t.val <= v {
			if t.left != nil {
				t = t.left
			} else {
				t.left = &BSTNode{
					val:   v,
				}
				return t.left
			}
		} else {
			// t.val > v
			if t.right != nil {
				t = t.right
			} else {
				t.right = &BSTNode{
					val: v,
				}
				return t.right
			}
		}
	}

	return nil
}

func successor(head *BSTNode) int {
	t := head.right
	for t.left != nil {
		t = t.left
	}
	return t.val
}

func predecessor(head *BSTNode) int {
	t := head.left
	for t.right != nil {
		t = t.right
	}
	return t.val
}

func del(v int, head *BSTNode) *BSTNode {
	if head == nil {
		return nil
	}

	if head.val < v {
		head.right = del(v, head.right)
	} else if head.val > v {
		head.left = del(v, head.left)
	} else {
		if head.left == nil && head.right == nil {
			return nil
		}
		if head.right != nil {
			head.val = successor(head)
			head.right = del(head.val, head.right)
		} else {
			head.val = predecessor(head)
			head.left = del(head.val, head.left)
		}
	}

	return head
}
