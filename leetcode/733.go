package main

// dfs
// 学到了最后变颜色

var dis = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

var nc int
var lenr int
var lenc int
var visited [51][51]bool

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	visited = [51][51]bool{}
	nc = newColor
	lenr = len(image)
	lenc = len(image[0])

	dfs(image, sr, sc)
	// image[sr][sc] = newColor

	return image
}

func dfs(image [][]int, sr, sc int) {
	if visited[sr][sc] {
		return
	}
	// fmt.Println(sr, sc)
	visited[sr][sc] = true

	for i := 0; i < 4; i++ {
		nsr := sr + dis[i][0]
		nsc := sc + dis[i][1]
		if nsr > -1 && nsc > -1 && nsr < lenr && nsc < lenc && image[sr][sc] == image[nsr][nsc] {
			dfs(image, nsr, nsc)
			// image[nsr][nsc] = nc
		}
	}
	image[sr][sc] = nc
}

