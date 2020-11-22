package lspartitioninglib

import (
	"fmt"
	"log"
	"testing"
)

const (
	graph_bench_1     = "graph_bench_1"
	graph_bench_1_opt = int64(14)
	graph_bench_2     = "graph_bench_2"
	graph_bench_2_opt = int64(14)
)

func BenchmarkLSPartiotionAlgorithmNonRec(b *testing.B) {
	graph := *NewGraph()
	if err := graph.ParseGraph(graph_bench_1); err != nil {
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
	graph := *NewGraph()
	if err := graph.ParseGraph(graph_bench_1); err != nil {
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
	for _, i := range res.Vector {
		if i {
			fmt.Print("1 ")
		} else {
			fmt.Print("0 ")
		}
	}
}

func BenchmarkLSPartiotionAlgorithmNonRecStatistic(b *testing.B) {
	b.Skip()
	graph := *NewGraph()
	if err := graph.ParseGraph(graph_bench_2); err != nil {
		log.Println(err)
		return
	}

	groupSize := graph.AmountOfVertex() / 2

	graph.HungryNumIndependent()
	for j := 0; j < b.N; j++ {
		LSPartiotionAlgorithmCountStatistic(&graph, nil, groupSize, graph_bench_2_opt)
	}

	fmt.Println("persent of true mark:", float64(Statistic.m[AmountOfTrueMark])/float64(Statistic.m[AmountOfMarkCount]))
	fmt.Println("persent of false mark:", float64(Statistic.m[AmountOfFalseMark])/float64(Statistic.m[AmountOfMarkCount]))
	fmt.Println("persent of disb fail:", float64(Statistic.m[AmountOfDisbalanceFail])/float64(Statistic.m[AmountOfItterations]))
	fmt.Println("persent of param fail:", float64(Statistic.m[AmountOfParamFail])/float64(Statistic.m[AmountOfItterations]))
	fmt.Println("midlle Mark Derivative:", float64(Statistic.m[OverallMarkDerivative])/float64(Statistic.m[AmountOfTrueMark]))
	fmt.Println("persent of one-s mark derivative:", float64(Statistic.m[MarkOneDerivative])/float64(Statistic.m[AmountOfMarkCount]))
	fmt.Println("persent of derivative in 0 to 4 from best:", float64(Statistic.m[AmountOfMarkDerivativeIn0To5])/float64(Statistic.m[AmountOfMarkCount]))
	fmt.Println("midlle mark derivative from best:", float64(Statistic.m[OverallMarkDerivativeFromBest])/float64(Statistic.m[AmountOfMarkCount]))
}
