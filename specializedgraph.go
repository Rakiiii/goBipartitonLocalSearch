package lspartitioninglib

import graphlib "github.com/Rakiiii/goGraph"
import gotuple "github.com/Rakiiii/goTuple"

type IGraph interface{
	//must implement default graph interface
	graphlib.IGraph
	//must return amount of independet vertex in grpah
	GetAmountOfIndependent()int
	//must set amount of independent vertex in graph
	SetAmountOfIndependent(int)
	//must renum vertex in graph in way that firstly must be independent vertx
	NumIndependent() []int
	HungryNumIndependent()[]int
	//must parse grpah from @param file of metis format
	ParseGraph(string) error
	//must return graph inisiated with dependent vertex set
	GetDependentGraph()Graph
}

//Graph if specialization of grpah for biderectional local search
type Graph struct {
	graphlib.Graph
	amountOfIndependent int
}

//GetAmountOfIndependent returns amount of independent vertex in the graph
func (g *Graph) GetAmountOfIndependent() int {
	return g.amountOfIndependent
}

//SetAmountOfIndependent sets amount of indepndent vertex in the graph
func (g *Graph)SetAmountOfIndependent(am int){
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
func (g *Graph)HungryNumIndependent()[]int{
	sortedOrd := make([]gotuple.IntTuple,g.AmountOfVertex())

	for i := 0 ; i < g.AmountOfVertex() ; i ++{
		sortedOrd[i].First = i
		sortedOrd[i].Second = len(g.GetEdges(i))
		sortedOrd[i].Third = 0
	}

	sortedOrd = gotuple.QuicksortIntTupleSecond(sortedOrd)

	/*for _,i := range sortedOrd{
		fmt.Println(i.First," ",i.Second," ",i.Third)
	}*/

	optPointers := make([]int,g.AmountOfVertex())
	for i,num := range sortedOrd{
		optPointers[num.First] = i
	}

	newOrd := make([]int,g.AmountOfVertex())
	counter := 0 
	 
	for _, vertex := range sortedOrd{
		if vertex.Third == 0{
			newOrd[counter] = vertex.First
			counter ++
			for _,vertex := range g.GetEdges(vertex.First){
				sortedOrd[optPointers[vertex]].Third = 1
			}
		}
	}

	g.amountOfIndependent = counter

	for _,vertex := range sortedOrd{
		if vertex.Third == 1{
			newOrd[counter] = vertex.First
			counter++
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

//GetDependentGraph returns type Graph that contains graph of denpendent vertex
func (g *Graph)GetDependentGraph()Graph{
	var newGraph Graph
	if(g.amountOfIndependent <= 0){
		return *g
	}

	for i := g.GetAmountOfIndependent();i < g.AmountOfVertex(); i++{
		newSet := make([]int,0)
		for _,v := range g.GetEdges(i){
			if v > g.GetAmountOfIndependent(){
				newSet = append(newSet,v-g.GetAmountOfIndependent())
			}
		}
	
		newGraph.AddVertexWithEdges(newSet)
	}

	return newGraph
}



