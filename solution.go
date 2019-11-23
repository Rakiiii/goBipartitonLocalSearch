package bipartitonlocalsearchlib

import pairlib "github.com/Rakiiii/goPair"

//Solution is struct that represent the solution of graph bipartition
//contains graph it self and vector with solution with value
type Solution struct {
	Value  int
	Vector []bool
	Gr     *Graph
	mark   []pairlib.IntPair
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

	return mark + s.CountParameter()
}

//PartIndependent parts independent vertexes in hungry form
func (s *Solution) PartIndependent(groupSize int) {
	s.mark = pairlib.QuicksortIntPairSecond(s.mark)

	groupTwoSize := 0

	for _, flag := range s.Vector {
		if flag {
			groupTwoSize++
		}
	}

	for i := 0; i < groupSize-groupTwoSize; i++ {
		s.Vector[s.mark[i].First] = true
	}

}
