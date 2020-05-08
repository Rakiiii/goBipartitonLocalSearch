package lspartitioninglib

import (
	"fmt"
	"log"
	"testing"
)

func TestLSPartiotionAlgorithmNonRec(t *testing.T) {
	fmt.Println("Start TestLSPartiotionAlgorithmNonRec")
	var graph Graph
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
	var graph Graph
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
	var graph Graph
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
		fmt.Println("TestLSPartiotionAlgorithmNonRecFast=[ok]")
	}
}
