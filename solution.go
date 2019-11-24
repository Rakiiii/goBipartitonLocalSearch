package bipartitonlocalsearchlib

import (
	"math"

	pairlib "github.com/Rakiiii/goPair"
)

//Solution is struct that represent the solution of graph bipartition
//contains graph it self and vector with solution with value
type Solution struct {
	Value  int64
	Vector []bool
	Gr     *Graph
	mark   []pairlib.IntPair
}

//Init initialized solution for graph
func (s *Solution) Init(g *Graph) {
	s.Value = -1
	s.Gr = g
	s.Vector = make([]bool, g.AmountOfVertex())
}

//CountParameter return amount of edges between groups for dependent grpah vartex
func (s *Solution) CountParameter() int64 {
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
func (s *Solution) SetDependentAsBinnary(num int64) {
	for i := len(s.Vector) - 1; i >= s.Gr.GetAmountOfIndependent(); i-- {
		if num%2 == 1 {
			s.Vector[i] = true
		} else {
			s.Vector[i] = false
		}
	}
}

//CountMark returns low mark of possible partition of independent vertrexes
func (s *Solution) CountMark() int64 {
	s.mark = make([]pairlib.IntPair, s.Gr.GetAmountOfIndependent())
	var mark int64 = 0
	for i := 0; i < s.Gr.GetAmountOfIndependent(); i++ {
		var fg int64 = 0
		var sg int64 = 0

		for _, v := range s.Gr.GetEdges(i) {
			switch {
			case v >= s.Gr.GetAmountOfIndependent():
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
		s.mark[i].Second = int(fg - sg)
	}

	return mark + s.CountParamForDependent()
}

//PartIndependent parts independent vertexes in hungry form
func (s *Solution) PartIndependent(groupSize int) bool {
	s.mark = pairlib.QuicksortIntPairSecond(s.mark)

	groupTwoSize := 0

	for _, flag := range s.Vector {
		if flag {
			groupTwoSize++
		}
	}

	if math.Abs(float64(groupSize-groupTwoSize)) >= float64(len(s.mark)) {
		return false
	}

	for i := 0; i < groupSize-groupTwoSize; i++ {
		s.Vector[s.mark[i].First] = true
	}
	return true
}

//CountParamForDependent counting param for subgraph containted dependent vertexes
func (s *Solution) CountParamForDependent() int64 {
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
