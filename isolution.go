package lspartitioninglib

type ISolution interface {
	GetValue() int64
	GetVector() []bool
	GetGraph() IGraph
	SetSolution(ISolution)
}
