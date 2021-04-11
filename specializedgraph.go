package lspartitioninglib

import (
	graphlib "github.com/Rakiiii/goGraph"
	pairlib "github.com/Rakiiii/goPair"
	gotuple "github.com/Rakiiii/goTuple"
)

type IGraph interface {
	//must implement default graph interface
	graphlib.IGraph
	//must return amount of independet vertex in grpah
	GetAmountOfIndependent() int
	//must return third level size in independent subset
	GetThirdLevelSize() int
	NumDependentOptimalThirdLevel() []int
	//must set amount of independent vertex in graph
	SetAmountOfIndependent(int)
	//must renum vertex in graph in way that firstly must be independent vertx
	NumIndependent() []int
	HungryNumIndependent() []int
	//must renum vertex in graph in way that firstly must be independet subset and at the end of subset is vertexes that connected with last n vertexes of dependent set
	NumThreeLevel(int) []int
	//must parse grpah from @param file of metis format
	ParseGraph(string) error
	//must return graph inisiated with dependent vertex set
	GetDependentGraph() Graph
}

//Graph if specialization of grpah for biderectional local search
type Graph struct {
	graphlib.IGraph
	amountOfIndependent int
	sizeOgThirdLevel    int
}

//NewGraphFromBase returns new lspartionable graph that decorates baseGraph
func NewGraphFromBase(baseGraph graphlib.IGraph) *Graph {
	return &Graph{IGraph: baseGraph, amountOfIndependent: 0}
}

//NewGraph returns new empty lspartitionable graph that decorates graphlib.Graph
func NewGraph() *Graph {
	baseGraph := &graphlib.Graph{}
	baseGraph.Init(0, 0)
	return NewGraphFromBase(baseGraph)
}

//NewGraph returns new empty lspartitionable graph that decorates graphlib.FastGraph
func NewGraphFast() *Graph {
	baseGraph := &graphlib.FastGraph{}
	baseGraph.Init(0, 0)
	return NewGraphFromBase(baseGraph)
}

//GetAmountOfIndependent returns amount of independent vertex in the graph
func (g *Graph) GetAmountOfIndependent() int {
	return g.amountOfIndependent
}

//SetAmountOfIndependent sets amount of indepndent vertex in the graph
func (g *Graph) SetAmountOfIndependent(am int) {
	g.amountOfIndependent = am
}

//NumIndependent renumberring vertex with some independetn subset of vertex
func (g *Graph) NumIndependent() []int {

	colorSet := make([]int, g.AmountOfVertex())
	amountOfIndependent := 0

	for i := 0; i < g.AmountOfVertex(); i++ {
		if colorSet[i] == 0 {
			amountOfIndependent++
			colorSet[i] = 1
			for _, v := range g.GetEdges(i) {
				colorSet[v] = 2
			}
		}
	}

	newOrd := make([]int, g.AmountOfVertex())
	p1 := 0
	p2 := amountOfIndependent
	for i, flag := range colorSet {
		if flag == 2 {
			newOrd[p2] = i
			p2++

		} else {
			newOrd[p1] = i
			p1++
		}
	}

	g.amountOfIndependent = amountOfIndependent

	g.RenumVertex(newOrd)
	return newOrd
}

//HungryNumIndependent realization of hugnry search algorithm of max independent vertex set with renumbering prev set of vertex
func (g *Graph) HungryNumIndependent() []int {
	sortedOrd := make([]gotuple.IntTuple, g.AmountOfVertex())

	for i := 0; i < g.AmountOfVertex(); i++ {
		sortedOrd[i].First = i
		sortedOrd[i].Second = len(g.GetEdges(i))
		sortedOrd[i].Third = 0
	}

	sortedOrd = gotuple.QuicksortIntTupleSecond(sortedOrd)

	/*for _,i := range sortedOrd{
		fmt.Println(i.First," ",i.Second," ",i.Third)
	}*/

	optPointers := make([]int, g.AmountOfVertex())
	for i, num := range sortedOrd {
		optPointers[num.First] = i
	}

	newOrd := make([]int, g.AmountOfVertex())
	counter := 0

	for _, vertex := range sortedOrd {
		if vertex.Third == 0 {
			newOrd[counter] = vertex.First
			counter++
			for _, vertex := range g.GetEdges(vertex.First) {
				sortedOrd[optPointers[vertex]].Third = 1
			}
		}
	}

	g.amountOfIndependent = counter

	for _, vertex := range sortedOrd {
		if vertex.Third == 1 {
			newOrd[counter] = vertex.First
			counter++
		}
	}

	g.RenumVertex(newOrd)

	return newOrd

}

func (g *Graph) NumThreeLevel(size int) []int {
	subSet := make(map[int]bool, 0)
	for i := g.AmountOfVertex() - 1; i > g.AmountOfVertex()-1-size; i-- {
		for _, vertex := range g.GetEdges(i) {
			if vertex < g.amountOfIndependent {
				subSet[vertex] = true
			}
		}
	}
	newOrder := make([]int, g.AmountOfVertex())
	position := 0
	for i := 0; i < g.amountOfIndependent; i++ {
		if subSet[i] == false {
			newOrder[position] = i
			position++
		}
	}
	for key, _ := range subSet {
		newOrder[position] = key
		position++
	}
	for ; position < g.AmountOfVertex(); position++ {
		newOrder[position] = position
	}
	g.RenumVertex(newOrder)
	g.sizeOgThirdLevel = len(subSet)
	return newOrder
}

func (g *Graph) NumDependentOptimalThirdLevel() []int {
	dependent := make([]pairlib.IntPair, g.AmountOfVertex()-g.amountOfIndependent)
	for i := 0; i < len(dependent); i++ {
		dependent[i] = pairlib.IntPair{First: i + g.amountOfIndependent, Second: len(g.GetEdges(i + g.amountOfIndependent))}
	}
	dependent = pairlib.ReversIntPairSlice(pairlib.QuicksortIntPairSecond(dependent))
	newOrd := make([]int, g.AmountOfVertex())
	for i := 0; i < len(newOrd); i++ {
		if i < g.amountOfIndependent {
			newOrd[i] = i
		} else {
			newOrd[i] = dependent[i-g.amountOfIndependent].First
		}
	}
	g.RenumVertex(newOrd)
	return newOrd
}

func (g *Graph) GetThirdLevelSize() int {
	return g.sizeOgThirdLevel
}

//ParseGraph parsing graph of metis format from file @path returning errors of reading file
func (g *Graph) ParseGraph(path string) error {
	var parser graphlib.Parser

	gr, err := parser.ParseUnweightedUndirectedGraphFromFile(path)
	if err != nil {
		return err
	}

	for i := 0; i < gr.AmountOfVertex(); i++ {
		g.AddVertexWithEdges(gr.GetEdges(i))
	}
	return nil
}

//GetDependentGraph returns type Graph that contains graph of denpendent vertex
func (g *Graph) GetDependentGraph() Graph {
	newGraph := *NewGraph()
	if g.amountOfIndependent <= 0 {
		return *g
	}

	for i := g.GetAmountOfIndependent(); i < g.AmountOfVertex(); i++ {
		newSet := make([]int, 0)
		for _, v := range g.GetEdges(i) {
			if v > g.GetAmountOfIndependent() {
				newSet = append(newSet, v-g.GetAmountOfIndependent())
			}
		}
		newGraph.AddVertexWithEdges(newSet)
	}

	return newGraph
}
