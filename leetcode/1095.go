package main

// 可以提出来做函数

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	l := 0
	// mid := mountainArr.get(mountainArr.length()/2)
	r := mountainArr.length() - 1

	// 1 2 3 4 3 2 1
	for l < r {
		mid := l + (r-l)/2
		// 在下降
		if mountainArr.get(mid) > mountainArr.get(mid+1) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if mountainArr.get(l) == target {
		return l
	}

	// fmt.Println(l)
	// 先搜索左边
	r = l
	l = 0
	for l <= r {
		mid := l + (r-l)/2
		val := mountainArr.get(mid)
		if val == target {
			return mid
		}

		// 在上升
		if val > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	r = mountainArr.length() - 1
	for l <= r {
		mid := l + (r-l)/2
		val := mountainArr.get(mid)
		if val == target {
			return mid
		}

		// 在下降
		if val < target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
