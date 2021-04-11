package lspartitioninglib

import (
	"fmt"
	"log"
	"testing"
)

func TestLSPartiotionAlgorithmNonRec(t *testing.T) {
	fmt.Println("Start TestLSPartiotionAlgorithmNonRec")
	graph := *NewGraph()
	if err := graph.ParseGraph("testgraph"); err != nil {
		log.Println(err)
		return
	}

	groupSize := graph.AmountOfVertex() / 2

	graph.HungryNumIndependent()

	res := LSPartiotionAlgorithmNonRec(&graph, nil, groupSize)

	if res.Value != 5 {
		t.Error("wrong value for LSPartiotionAlgorithmNonRec:", res.Value)
	} else {
		fmt.Println("TestLSPartiotionAlgorithmNonRec=[ok]")
	}
}

func TestLSPartiotionAlgorithm(t *testing.T) {
	fmt.Println("Start TestLSPartiotionAlgorithm")
	graph := *NewGraph()
	if err := graph.ParseGraph("testgraph"); err != nil {
		log.Println(err)
		return
	}

	groupSize := graph.AmountOfVertex() / 2

	graph.HungryNumIndependent()
	//sol.CountParameter()

	res := LSPartiotionAlgorithm(&graph, nil, groupSize, 0)

	if res.Value != 5 {
		t.Error("wrong value for LSPartiotionAlgorithm:", res.Value)
	} else {
		fmt.Println("TestLSPartiotionAlgorithm=[ok]")
	}
}

func TestLSPartiotionAlgorithmNonRecFast(t *testing.T) {
	fmt.Println("Start TestLSPartiotionAlgorithmNonRecFast")
	graph := *NewGraph()
	if err := graph.ParseGraph("testgraph"); err != nil {
		log.Println(err)
		return
	}

	groupSize := graph.AmountOfVertex() / 2

	graph.HungryNumIndependent()

	res := LSPartiotionAlgorithmNonRecFast(&graph, nil, groupSize)

	if res.Value != 5 {
		t.Error("wrong value for LSPartiotionAlgorithmNonRecFast:", res.Value)
	} else {
		fmt.Println("result:", res.Vector)
		fmt.Println("TestLSPartiotionAlgorithmNonRecFast=[ok]")
	}
}

func TestThreeLevelPartiotionAlgorithmNonRec(t *testing.T) {
	fmt.Println("Start ThreeLevelPartiotionAlgorithmNonRec")
	graph := *NewGraph()
	if err := graph.ParseGraph("testgraph"); err != nil {
		log.Println(err)
		return
	}

	groupSize := graph.AmountOfVertex() / 2

	graph.HungryNumIndependent()

	res := ThreeLevelPartiotionAlgorithmNonRec(&graph, nil, groupSize, 1)

	if res.GetValue() != 5 {
		fmt.Println("result:", TranslateResultVectorToOut(res.GetVector(), ord))
		t.Error("wrong value for ThreeLevelPartiotionAlgorithmNonRec:", res.GetValue())
	} else {
		fmt.Println("result:", TranslateResultVectorToOut(res.GetVector(), ord))
		fmt.Println("ThreeLevelPartiotionAlgorithmNonRec=[ok]")
	}
}
