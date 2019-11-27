package bipartitonlocalsearchlib

import graphlib "github.com/Rakiiii/goGraph"
import gotuple "github.com/Rakiiii/goTuple"

//Graph if specialization of grpah for biderectional local search
type Graph struct {
	graphlib.Graph
	amountOfIndependent int
}

//GetAmountOfIndependent returns amount of independent vertex in the graph
func (g *Graph) GetAmountOfIndependent() int {
	return g.amountOfIndependent
}

//NumIndependent renumberring vertex with some independetn subset of vertex
func (g *Graph) NumIndependent() []int {
	colorSet := make([]bool, g.AmountOfVertex())
	amountOfIndependent := 0

	for i := 0; i < g.AmountOfVertex(); i++ {
		if !colorSet[i] {
			amountOfIndependent++
			for _, v := range g.GetEdges(i) {
				colorSet[v] = true
			}
		}
	}

	newOrd := make([]int, g.AmountOfVertex())
	p1 := 0
	p2 := amountOfIndependent
	for i, flag := range colorSet {
		if flag {
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
func (g *Graph)HungryNumIndependent()[]int{
	sortedOrd := make([]goTuple.IntTuple,g.AmountOfVertex())

	for i := 0 ; i < g.AmountOfVertex() ; i ++{
		newOrd[i].First = i
		newOrd[i].Second = len(g.GetEdges(i))
		newOrd[i].Third = 0
	}

	sortedOrd = gotuple.QuicksortIntTupleSecond(sortedOrd)

	optPointers := make([]int,g.AmountOfVertex())
	for i,num := range sortedOrd{
		optPointers[num.First] = i
	}

	newOrd := make([]int,g.AmountOfVertex())
	counter := 0 
	 
	for i, vertex := range sortedOrd{
		if vertex.Third == 0{
			newOrd[counter] = vertex.First
			counter ++
			for _,vertex := range g.GetEdges(vertex.First){
				sortedOrd[optPointers[vertex]].Third = 1
			}
		}
	}

	g.AmountOfVertex = counter

	for i,vertex := range sortedOrd{
		if vertex.Third == 1{
			newOrd[counter] = vertex.First
		}
	}

	g.RenumVertex(newOrd)

	return newOrd

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
