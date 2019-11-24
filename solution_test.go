package bipartitonlocalsearchlib

import (
	"fmt"
	"testing"
)

func TestSetDependentAsBinnary(t *testing.T) {
	var graph Graph
	graph.ParseGraph("test_grSmall")
	fmt.Println(graph.NumIndependent())

	var sol Solution
	sol.Init(&graph)

	sol.SetDependentAsBinnary(9)
	fmt.Println("param:", sol.CountParamForDependent(), "| mark:", sol.CountMark())

}
