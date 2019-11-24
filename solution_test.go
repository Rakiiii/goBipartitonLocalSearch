package bipartitonlocalsearchlib

import "testing"

func TestSetDependentAsBinnary(t *testing.T) {
	var graph Graph
	graph.ParseGraph("test_gr3")
	graph.NumIndependent()

	var sol Solution
	sol.Init(&graph)

}
