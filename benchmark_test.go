package lspartitioninglib

import (
	"fmt"
	"log"
	"testing"
)

const (
	graph_bench = "graph_bench"
)

func BenchmarkLSPartiotionAlgorithmNonRec(b *testing.B) {
	var graph Graph
	if err := graph.ParseGraph(graph_bench); err != nil {
		log.Println(err)
		return
	}

	groupSize := graph.AmountOfVertex() / 2

	var res *Solution

	graph.HungryNumIndependent()
	for j := 0; j < b.N; j++ {
		res = LSPartiotionAlgorithmNonRec(&graph, nil, groupSize)
	}
	fmt.Println(res.Value)
	fmt.Println(res.Vector)
}

func BenchmarkLSPartiotionAlgorithmNonRecFast(b *testing.B) {
	var graph Graph
	if err := graph.ParseGraph(graph_bench); err != nil {
		log.Println(err)
		return
	}

	groupSize := graph.AmountOfVertex() / 2

	var res *Solution
	graph.HungryNumIndependent()
	for j := 0; j < b.N; j++ {
		res = LSPartiotionAlgorithmNonRecFast(&graph, nil, groupSize)
	}
	fmt.Println(res.Value)
	fmt.Println(res.Vector)
}

func BenchmarkLSPartiotionAlgorithmNonRecStatistic(b *testing.B) {
	var graph Graph
	if err := graph.ParseGraph(graph_bench); err != nil {
		log.Println(err)
		return
	}

	groupSize := graph.AmountOfVertex() / 2

	graph.HungryNumIndependent()
	for j := 0; j < b.N; j++ {
		LSPartiotionAlgorithmCountStatistic(&graph, nil, groupSize)
	}

	fmt.Println("persent of true mark:", float64(Statistic.m[AmountOfTrueMark])/float64(Statistic.m[AmountOfMarkCount]))
	fmt.Println("persent of false mark:", float64(Statistic.m[AmountOfFalseMark])/float64(Statistic.m[AmountOfMarkCount]))
	fmt.Println("persent of disb fail:", float64(Statistic.m[AmountOfDisbalanceFail])/float64(Statistic.m[AmountOfItterations]))
	fmt.Println("persent of param fail:", float64(Statistic.m[AmountOfParamFail])/float64(Statistic.m[AmountOfItterations]))
	fmt.Println("midlle Mark Derivative:", float64(Statistic.m[OverallMarkDerivative])/float64(Statistic.m[AmountOfTrueMark]))
	fmt.Println("persent of one-s mark derivative:", float64(Statistic.m[MarkOneDerivative])/float64(Statistic.m[AmountOfMarkCount]))
}
