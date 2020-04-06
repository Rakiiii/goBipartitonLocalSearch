package lspartitioninglib

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

func TestTranslateResultVector(t *testing.T){
	fmt.Println("Start TestTranslateResultVector")
	old := []bool{true,true,true,true,false,false,false,false}
	new := []bool{false,false,false,false,true,true,true,true}
	ord := []int{4,5,6,7,0,1,2,3}
	old = TranslateResultVector(old,ord)
	checkFlag := true
	for i,j := range old{
		if j != new[i] {
			t.Error("Wrong at pos:",i)
			checkFlag = false
		}
	}
	if checkFlag{
		fmt.Println("TestTranslateResultVector=[ok]")
	}
}

func TestTranslateResultVectorToOut(t *testing.T){
	fmt.Println("Start TestTranslateResultVectorToOut")
	old := []bool{true,true,true,true,false,false,false,false}
	new := []int{0,0,0,0,1,1,1,1}
	ord := []int{4,5,6,7,0,1,2,3}
	sub := TranslateResultVectorToOut(old,ord)
	checkFlag := true
	for i,j := range sub{
		if j != new[i] {
			t.Error("Wrong at pos:",i)
			checkFlag = false
		}
	}
	if checkFlag{
		fmt.Println("TestTranslateResultVectorToOut=[ok]")
	}
}
