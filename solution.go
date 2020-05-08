package lspartitioninglib

import (
	pairlib "github.com/Rakiiii/goPair"
)

//Solution is struct that represent the solution of graph bipartition
//contains graph it self and vector with solution with value
type Solution struct {
	Value  int64
	Vector []bool
	Gr     IGraph
	mark   []pairlib.IntPair
}

//Init initialized solution for graph
func (s *Solution) Init(g IGraph) {
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
		num /= 2
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
			//fmt.Println("vertex:", v)
			switch {
			case v < s.Gr.GetAmountOfIndependent():
				//fmt.Println("out of index")
			case !s.Vector[v]:
				fg++
			case s.Vector[v]:
				sg++
			}
		}
		//fmt.Println("fg:", fg, " sg:", sg)
		if fg > sg {
			mark += sg
		} else {
			mark += fg
		}
		s.mark[i].First = i
		s.mark[i].Second = int(sg - fg)
	}

	return mark + s.CountParamForDependent()
}

//PartIndependent parts independent vertexes in hungry form
func (s *Solution) PartIndependent(groupSize int) bool {

	groupTwoSize := 0

	for _, flag := range s.Vector {
		if flag {
			groupTwoSize++
		}
	}

	if groupSize-groupTwoSize < 0 || groupSize-groupTwoSize > s.Gr.GetAmountOfIndependent() {
		return false
	}

	s.mark = pairlib.ReversIntPairSlice(pairlib.QuicksortIntPairSecond(s.mark))
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

//TranslateResultVector renum posirions of elems in @vec with new order @ord
func TranslateResultVector(vec []bool, ord []int) []bool {
	formatedRes := make([]bool, len(ord))
	for i, num := range ord {
		if vec[i] {
			formatedRes[num] = true
		} else {
			formatedRes[num] = false
		}
	}
	return formatedRes
}

//TranslateResultVector renum posirions of elems in @vec with new order @ord
func TranslateResultVectorToOut(vec []bool, ord []int) []int {
	formatedRes := make([]int, len(ord))
	for i, num := range ord {
		if vec[i] {
			formatedRes[num] = 1
		} else {
			formatedRes[num] = 0
		}
	}
	return formatedRes
}
