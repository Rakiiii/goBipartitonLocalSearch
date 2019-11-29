package bipartitonlocalsearchlib

import (
	"fmt"
	"testing"
)

func TestSetDependentAsBinnary(t *testing.T) {
	fmt.Println("test set dependent as binnary")
	var graph Graph
	graph.ParseGraph("test_grSmall")
	fmt.Println(graph.NumIndependent())

	var sol Solution
	sol.Init(&graph)

	sol.SetDependentAsBinnary(0)
	fmt.Println("param:", sol.CountParamForDependent(), "| mark:", sol.CountMark())

	sol.PartIndependent(4)
	fmt.Println(sol.Vector, " ", sol.CountParameter())

}

func TestCountParameter(t *testing.T){
	fmt.Println("count parameter test")

	var graph Graph
	graph.ParseGraph("test_grSmall")

	/*ord :=*/ graph.HungryNumIndependent()

	var sol Solution
	sol.Init(&graph)

	sol.Vector[0] = true
	sol.Vector[2] = true
	sol.Vector[3] = true
	sol.Vector[6] = true

	fmt.Println("vector:",sol.Vector," parameter:",sol.CountParameter())
}
