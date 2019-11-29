package bipartitonlocalsearchlib

import (
	"fmt"
	"testing"

	graphlib "github.com/Rakiiii/goGraph"
)

func TestParseGraph(t *testing.T) {
	fmt.Println("Parse test")
	var parser graphlib.Parser
	gr, err := parser.ParseUnweightedUndirectedGraphFromFile("test_gr3")
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < gr.AmountOfVertex(); i++ {
		fmt.Println(gr.GetEdges(i))
	}

	fmt.Println("am of ver:", gr.AmountOfVertex())

	fmt.Println("newGraph")

	var graph Graph
	graph.ParseGraph("test_gr3")
	for i := 0; i < 14; i++ {
		fmt.Println(i," | ",graph.GetEdges(i))
	}

	fmt.Println("am of ver:", graph.AmountOfVertex())

	var g Graph
	g.ParseGraph("test_grSmall")
	for i := 0; i < g.AmountOfVertex(); i++ {
		fmt.Println(i," | ",g.GetEdges(i))
	}
}

func TestHungryNumIndependent(t *testing.T){

	fmt.Println("HungryNumIndependentTest")
	var graph Graph
	graph.ParseGraph("test_grSmall")
	fmt.Println(graph.HungryNumIndependent()," ",graph.GetAmountOfIndependent())
	fmt.Println()
	for i := 0; i < graph.AmountOfVertex(); i++ {
		fmt.Println(i," | ",graph.GetEdges(i))
	}
}
