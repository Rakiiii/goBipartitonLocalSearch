package lspartitioninglib

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

	graph := *NewGraph()
	graph.ParseGraph("test_gr3")
	for i := 0; i < 14; i++ {
		fmt.Println(i, " | ", graph.GetEdges(i))
	}

	fmt.Println("am of ver:", graph.AmountOfVertex())

	g := *NewGraph()
	g.ParseGraph("test_grSmall")
	for i := 0; i < g.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", g.GetEdges(i))
	}
}

func TestHungryNumIndependent(t *testing.T) {

	fmt.Println("HungryNumIndependentTest")
	graph := *NewGraph()
	graph.ParseGraph("test_grSmall")
	fmt.Println(graph.HungryNumIndependent(), " ", graph.GetAmountOfIndependent())
	fmt.Println()
	for i := 0; i < graph.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", graph.GetEdges(i))
	}
}

func TestNumIndependent(t *testing.T) {

	fmt.Println("NumIndependentTest")
	graph := *NewGraph()
	graph.ParseGraph("test_grSmall")
	fmt.Println(graph.NumIndependent(), " ", graph.GetAmountOfIndependent())
	fmt.Println()
	for i := 0; i < graph.AmountOfVertex(); i++ {
		fmt.Println(i, " | ", graph.GetEdges(i))
	}
}

func TestGetDependentGraph1(t *testing.T) {

	fmt.Println("GetDependentGraphTestWithHI")
	graph := *NewGraph()
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

func TestGetDependentGraph2(t *testing.T) {

	fmt.Println("GetDependentGraphTestWithHI")
	graph := *NewGraph()
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
