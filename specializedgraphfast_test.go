package lspartitioninglib

import (
	"fmt"
	"testing"

	graphlib "github.com/Rakiiii/goGraph"
)

func TestParseGraphFast(t *testing.T) {
	fmt.Println("Parse test FastGraph")
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

	graph := *NewGraphFast()
	graph.ParseGraph("test_gr3")
	for i := 0; i < 14; i++ {
		fmt.Println(i, " | ", graph.GetEdges(i))
	}

	fmt.Println("am of ver:", graph.AmountOfVertex())

	g := *NewGraphFast()
	g.ParseGraph("test_grSmall")
	for i := 0; i < g.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", g.GetEdges(i))
	}
}

func TestHungryNumIndependentFast(t *testing.T) {

	fmt.Println("HungryNumIndependentFastGraphTest")
	graph := *NewGraphFast()
	graph.ParseGraph("test_grSmall")
	fmt.Println(graph.HungryNumIndependent(), " ", graph.GetAmountOfIndependent())
	fmt.Println()
	for i := 0; i < graph.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", graph.GetEdges(i))
	}
}

func TestNumIndependentFast(t *testing.T) {

	fmt.Println("NumIndependentFastGraphTest")
	graph := *NewGraphFast()
	graph.ParseGraph("test_grSmall")
	fmt.Println(graph.NumIndependent(), " ", graph.GetAmountOfIndependent())
	fmt.Println()
	for i := 0; i < graph.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", graph.GetEdges(i))
	}
}

func TestGetDependentGraph1Fast(t *testing.T) {

	fmt.Println("GetDependentGraphFastGraphTestWithHI")
	graph := *NewGraphFast()
	graph.ParseGraph("test_grSmall")
	for i := 0; i < graph.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", graph.GetEdges(i))
	}
	fmt.Println(graph.HungryNumIndependent(), " ", graph.GetAmountOfIndependent())
	fmt.Println()
	for i := 0; i < graph.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", graph.GetEdges(i))
	}
	newgr := graph.GetDependentGraph()
	for i := 0; i < newgr.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", newgr.GetEdges(i))
	}
}

func TestGetDependentGraph2Fast(t *testing.T) {

	fmt.Println("GetDependentGraphFastGraphTestWithHI")
	graph := *NewGraphFast()
	graph.ParseGraph("test_grSmall")
	for i := 0; i < graph.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", graph.GetEdges(i))
	}
	fmt.Println(graph.NumIndependent(), " ", graph.GetAmountOfIndependent())
	fmt.Println()
	for i := 0; i < graph.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", graph.GetEdges(i))
	}
	newgr := graph.GetDependentGraph()
	for i := 0; i < newgr.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", newgr.GetEdges(i))
	}
}
