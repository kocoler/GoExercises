package main

//RankType 按秩合并类型
type RankType int

//按秩合并类型
const (
	RankNone   RankType = iota //不按秩合并
	RankHeight                 //按树的高度合并
	RankSize                   //按树的节点数合并
)

//UnionFind 并查集
type UnionFind interface {
	Union(x, y int) bool
	Find(x int) int
	Rank(x int) int
	Count() int
	IsConnected(x, y int) bool
}

//MapUnionFind 并查集（map存储）
type MapUnionFind struct {
	parent   map[int]int
	rank     map[int]int
	count    int
	rankType RankType
}

//NewMapUnionFind 初始化并查集
func NewMapUnionFind() *MapUnionFind {
	uf := &MapUnionFind{
		parent: make(map[int]int, 0),
	}
	return uf
}

//NewMapUnionFindWithRank 初始化带秩的并查集
func NewMapUnionFindWithRank(n int, rankType RankType) *MapUnionFind {
	uf := &MapUnionFind{
		parent:   make(map[int]int, n),
		rank:     make(map[int]int, n),
		rankType: rankType,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = -1
		uf.rank[i] = 1
	}
	return uf
}

//Union 合并两个节点,按秩合并
func (s *MapUnionFind) Union(x, y int) bool {
	rootX, rootY := s.Find(x), s.Find(y)
	if rootX != rootY {
		switch s.rankType {
		case RankNone:
			s.parent[rootX] = rootY
		case RankHeight:
			if s.rank[rootX] == s.rank[rootY] {
				s.rank[rootY]++
			} else if s.rank[rootX] > s.rank[rootY] {
				rootX, rootY = rootY, rootX
			}
			s.parent[rootX] = rootY
		case RankSize:
			if s.rank[rootX] >= s.rank[rootY] {
				rootX, rootY = rootY, rootX
			}
			s.parent[rootX] = rootY
			s.rank[rootY] += s.rank[rod f dotX]
		}
		s.count--
		return true
	}
	return false
}

//Find 路径压缩
func (s *MapUnionFind) Find(x int) int {
	if _, ok := s.parent[x]; !ok {
		s.parent[x] = x
		s.count++
	}
	if x != s.parent[x] {
		s.parent[x] = s.Find(s.parent[x])
	}
	return s.parent[x]
}

//IsConnected 节点是否连通
func (s *MapUnionFind) IsConnected(x, y int) bool {
	return s.Find(x) == s.Find(y)
}

//Count ..
func (s *MapUnionFind) Count() int {
	return s.count
}

//Rank ..
func (s *MapUnionFind) Rank(x int) int {
	return s.rank[x]
}

