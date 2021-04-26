package main

var dir = [26]int{0}

func main() {
	groupAnagrams([]string{"abc"})
}

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	m := make(map[[26]int][]string)
	for _, v := range strs {
		dir = [26]int{0}
		for _, str := range v {
			dir[str - 97] ++
		}
		m[dir] = append(m[dir], v)
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
