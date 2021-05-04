package lspartitioninglib

//Solution is struct that represent the solution of graph bipartition
//contains graph it self and vector with solution with value
type FourLevelSolution struct {
	ThreeLevelSolution
	vectorDependendentCache []bool
	isSetted                bool
}

//Init initialized solution for graph
func (s *FourLevelSolution) Init(g IGraph, levelSize int) {
	s.ThreeLevelSolution.Init(g, levelSize)
	s.isSetted = false
	s.vectorDependendentCache = make([]bool, g.AmountOfVertex()-g.GetAmountOfIndependent())
}

//Init initialized solution for graph
func (s *FourLevelSolution) ReInit() {
	s.ThreeLevelSolution.ReInit()
	s.isSetted = false
}

func (s *FourLevelSolution) ClearSub() {
	for i := 0; i < s.Gr.GetAmountOfIndependent(); i++ {
		s.Vector[i] = false
	}
	s.Value = -1
}

func (s *FourLevelSolution) cacheVector() {
	offset := s.Gr.GetAmountOfIndependent()
	for i := offset; i < len(s.Vector)-1; i++ {
		s.vectorDependendentCache[i-offset] = s.Vector[i]
	}
	s.isSetted = true
}

func (s *FourLevelSolution) countDiffSize() int {
	diff := 0
	offset := s.Gr.AmountOfVertex() - s.levelSize
	for i := offset; i < s.Gr.AmountOfVertex(); i++ {
		if s.vectorDependendentCache[i-offset] != s.Vector[i] {
			diff++
		}
	}
	return diff
}

func (s *FourLevelSolution) updateDiffBySize(diffSize int) {
	minNumberOfNotCounted := s.Gr.GetAmountOfIndependent() - s.Gr.GetThirdLevelSize()
	offset := len(s.Vector) - 1 - diffSize
	for i := len(s.Vector) - 1; i > offset; i-- {
		for _, v := range s.subgraph.GetEdges(i - offset - 1) {
			if !s.Vector[i] {
				s.lastVertex[v-minNumberOfNotCounted].First++
				s.lastVertex[v-minNumberOfNotCounted].Second--
			} else {
				s.lastVertex[v-minNumberOfNotCounted].Second++
				s.lastVertex[v-minNumberOfNotCounted].First--
			}
		}
	}
}
