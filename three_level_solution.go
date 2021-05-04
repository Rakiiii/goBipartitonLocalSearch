package lspartitioninglib

import (
	"math"

	pairlib "github.com/Rakiiii/goPair"
)

//Solution is struct that represent the solution of graph bipartition
//contains graph it self and vector with solution with value
type ThreeLevelSolution struct {
	Value           int64
	Vector          []bool
	Gr              IGraph
	subgraph        IGraph
	levelSize       int
	mark            []pairlib.IntPair
	subMark         int64
	lastVertex      []pairlib.IntPair
	lastVertexСache []pairlib.IntPair
	lastPair        []pairlib.IntPair
}

//GetValue implements ISolution
func (s *ThreeLevelSolution) GetValue() int64 {
	return s.Value
}

//GetVector implements ISolution
func (s *ThreeLevelSolution) GetVector() []bool {
	return s.Vector
}

//GetGraph implements ISolution
func (s *ThreeLevelSolution) GetGraph() IGraph {
	return s.Gr
}

func (s *ThreeLevelSolution) SetSolution(sol ISolution) {
	s.Vector = make([]bool, len(sol.GetVector()))
	for i := 0; i < len(sol.GetVector()); i++ {
		s.Vector[i] = sol.GetVector()[i]
	}
	s.Value = sol.GetValue()
}

//Init initialized solution for graph
func (s *ThreeLevelSolution) Init(g IGraph, levelSize int) {
	s.Value = -1
	s.Gr = g
	s.Vector = make([]bool, g.AmountOfVertex())
	s.levelSize = levelSize
	s.mark = make([]pairlib.IntPair, s.Gr.GetAmountOfIndependent()-s.Gr.GetThirdLevelSize())
	s.subMark = -1
	s.lastVertex = make([]pairlib.IntPair, s.Gr.GetThirdLevelSize())
	s.lastVertexСache = make([]pairlib.IntPair, s.Gr.GetThirdLevelSize())
	s.lastPair = make([]pairlib.IntPair, s.Gr.GetThirdLevelSize())
}

//Init initialized solution for graph
func (s *ThreeLevelSolution) ReInit() {
	s.Value = -1
	for i := 0; i < len(s.Vector); i++ {
		s.Vector[i] = false
	}
	for i := 0; i < len(s.mark); i++ {
		s.mark[i] = pairlib.IntPair{}
	}
	for i := 0; i < len(s.lastVertex); i++ {
		s.lastVertex[i] = pairlib.IntPair{}
	}
	for i := 0; i < len(s.lastVertexСache); i++ {
		s.lastVertexСache[i] = pairlib.IntPair{}
	}
	for i := 0; i < len(s.lastPair); i++ {
		s.lastPair[i] = pairlib.IntPair{}
	}
	s.subMark = -1
}

func (s *ThreeLevelSolution) ClearSub() {
	for i := 0; i < len(s.lastVertexСache); i++ {
		s.lastVertex[i] = s.lastVertexСache[i]
	}
	for i := 0; i < len(s.lastPair); i++ {
		s.lastPair[i] = pairlib.IntPair{}
	}
	for i := 0; i < s.Gr.GetAmountOfIndependent(); i++ {
		s.Vector[i] = false
	}
	s.Value = -1
}

//Init initialized solution for graph
func (s *ThreeLevelSolution) cacheSubSetMark() {
	var mark int64 = 0
	for i := 0; i < s.Gr.GetAmountOfIndependent()-s.Gr.GetThirdLevelSize(); i++ {
		var fg int64 = 0
		var sg int64 = 0

		for _, v := range s.Gr.GetEdges(i) {
			switch {
			case v < s.Gr.GetAmountOfIndependent():
			case !s.Vector[v]:
				fg++
			case s.Vector[v]:
				sg++
			}
		}
		if fg > sg {
			mark += sg
		} else {
			mark += fg
		}
		s.mark[i].First = i
		s.mark[i].Second = int(sg - fg)
	}
	s.subMark = mark
	s.mark = pairlib.ReversIntPairSlice(pairlib.QuicksortIntPairSecond(s.mark))
	offset := s.Gr.GetAmountOfIndependent() - s.Gr.GetThirdLevelSize()
	for i := offset; i < s.Gr.GetAmountOfIndependent(); i++ {
		var fg int = 0
		var sg int = 0

		for _, v := range s.Gr.GetEdges(i) {
			switch {
			case v < s.Gr.GetAmountOfIndependent() || v >= s.Gr.AmountOfVertex()-s.levelSize:
			case !s.Vector[v]:
				fg++
			case s.Vector[v]:
				sg++
			}
		}

		s.lastVertex[i-offset].First = fg
		s.lastVertex[i-offset].Second = sg

		s.lastVertexСache[i-offset].First = fg
		s.lastVertexСache[i-offset].Second = sg
	}
}

//CountParameter return amount of edges between groups for dependent grpah vartex
func (s *ThreeLevelSolution) CountParameter() int64 {
	var param int64 = 0
	for i, group := range s.Vector {
		for _, v := range s.Gr.GetEdges(i) {
			if group != s.Vector[v] {
				param++
			}
		}
	}
	s.Value = param / 2
	return param / 2
}

//SetDependentAsBinnary setting partition of dependent vertexes as binnary form of number
func (s *ThreeLevelSolution) SetDependentAsBinnarySecondLevel(num int64) {
	i := len(s.Vector) - 1 - s.levelSize
	for ; i >= s.Gr.GetAmountOfIndependent(); i-- {
		if num%2 == 1 {
			s.Vector[i] = true
		} else {
			s.Vector[i] = false
		}
		num /= 2
	}
}

//SetDependentAsBinnary setting partition of dependent vertexes as binnary form of number
func (s *ThreeLevelSolution) SetDependentAsBinnaryThirdLevel(num int64) {
	i := len(s.Vector) - 1
	for ; i > len(s.Vector)-1-s.levelSize; i-- {
		if num%2 == 1 {
			s.Vector[i] = true
		} else {
			s.Vector[i] = false
		}
		num /= 2
	}
}

func (s *ThreeLevelSolution) constructSubGraph() {
	s.subgraph = NewGraphFast()
	minNumberOfNotCounted := s.Gr.GetAmountOfIndependent() - s.Gr.GetThirdLevelSize()
	i := len(s.Vector) - 1
	for ; i > len(s.Vector)-1-s.levelSize; i-- {
		edges := make([]int, 0)
		for _, v := range s.Gr.GetEdges(i) {
			switch {
			case v >= s.Gr.GetAmountOfIndependent() || v < minNumberOfNotCounted:
			default:
				edges = append(edges, v)
			}
		}
		s.subgraph.AddVertexWithEdges(edges)
	}
}

func (s *ThreeLevelSolution) OverCountThirdLevelWithSubgraph() {
	minNumberOfNotCounted := s.Gr.GetAmountOfIndependent() - s.Gr.GetThirdLevelSize()
	offset := s.Gr.AmountOfVertex() - s.levelSize
	for i := len(s.Vector) - 1; i < len(s.Vector)-1-s.levelSize; i-- {
		for _, v := range s.subgraph.GetEdges(i - offset) {
			if !s.Vector[i] {
				s.lastVertex[v-minNumberOfNotCounted].First++
			} else {
				s.lastVertex[v-minNumberOfNotCounted].Second++
			}
		}
	}
}

func (s *ThreeLevelSolution) OverCountThirdLevel() {
	minNumberOfNotCounted := s.Gr.GetAmountOfIndependent() - s.Gr.GetThirdLevelSize()
	i := len(s.Vector) - 1
	for ; i > len(s.Vector)-1-s.levelSize; i-- {
		for _, v := range s.Gr.GetEdges(i) {
			switch {
			case v >= s.Gr.GetAmountOfIndependent() || v < minNumberOfNotCounted:
			case !s.Vector[i]:
				s.lastVertex[v-minNumberOfNotCounted].First++
			case s.Vector[i]:
				s.lastVertex[v-minNumberOfNotCounted].Second++
			}
		}
	}
}

func (s *ThreeLevelSolution) CountMark() int64 {
	return s.subMark + s.CountParamForDependent() + s.OverCountMark()
}

func (s *ThreeLevelSolution) OverCountMark() int64 {
	mark := 0
	for _, i := range s.lastVertex {
		if i.First > i.Second {
			mark += i.Second
		} else {
			mark += i.First
		}
	}
	return int64(mark)
}

//PartIndependent parts independent vertexes in hungry form
func (s *ThreeLevelSolution) PartIndependent(groupSize int) bool {

	groupTwoSize := 0

	for _, flag := range s.Vector {
		if flag {
			groupTwoSize++
		}
	}

	if groupSize-groupTwoSize < 0 || groupSize-groupTwoSize > s.Gr.GetAmountOfIndependent() {
		return false
	}

	lastMap := make(map[int]int, 0)
	minNumberOfNotCounted := s.Gr.GetAmountOfIndependent() - s.Gr.GetThirdLevelSize()
	for i, p := range s.lastVertex {
		lastMap[minNumberOfNotCounted+i] = p.Second - p.First
	}
	var res int = 0
	position := 0
	for i := 0; i < groupSize-groupTwoSize; i++ {
		if res != -1 {
			if len(s.mark) == 0 {
				s.Vector[getSmallest(lastMap)] = true
			} else {
				if res = checkContains(lastMap, s.mark[i].Second); res != -1 {
					s.Vector[res] = true
				} else {
					s.Vector[s.mark[position].First] = true
					position++
				}
			}
		} else {
			if position < len(s.mark) {
				s.Vector[s.mark[position].First] = true
				position++
			} else {
				s.Vector[getSmallest(lastMap)] = true
			}
		}
	}
	return true
}

func checkContains(m map[int]int, val int) int {
	for key, value := range m {
		if value > val {
			delete(m, key)
			return key
		}
	}
	return -1
}

func getSmallest(m map[int]int) int {
	minValue := math.MaxInt32
	dKey := -1
	for key, value := range m {
		if value < minValue {
			minValue = value
			dKey = key
		}
	}
	delete(m, dKey)
	return dKey
}

//CountParamForDependent counting param for subgraph containted dependent vertexes
func (s *ThreeLevelSolution) CountParamForDependent() int64 {
	var param int64 = 0
	for i := s.Gr.GetAmountOfIndependent(); i < len(s.Vector); i++ {
		for _, v := range s.Gr.GetEdges(i) {
			if s.Vector[i] != s.Vector[v] && v > s.Gr.GetAmountOfIndependent() {
				param++
			}
		}
	}
	return param / 2
}

func (s *ThreeLevelSolution) String() string {
	var res string
	for _, b := range s.Vector {
		if b {
			res += "1 "
		} else {
			res += "0 "
		}
	}
	return res
}
