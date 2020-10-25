package lspartitioninglib

import (
	"errors"
	"log"
	"math"
	"os"
	"sync"
)

type SafeValue struct {
	Value int64
	Mux   sync.Mutex
}

//LSPartiotionAlgorithm return optimum of bipartion of graph @gr with biggest group size @groupSize start @number must be 0 if if you want to check all diapaaon
func LSPartiotionAlgorithm(gr IGraph, sol *Solution, groupSize int, number int64) *Solution {
	log.Println("Check number:", number)

	if float64(number) >= math.Pow(2, float64(gr.AmountOfVertex()-gr.GetAmountOfIndependent())) {
		log.Println("finish:", math.Pow(2, float64(gr.AmountOfVertex()-gr.GetAmountOfIndependent())))
		return sol
	}
	var newSol Solution

	log.Println("solution constructed")

	newSol.Init(gr)
	newSol.SetDependentAsBinnary(number)
	mark := newSol.CountMark()
	log.Println("mark:", mark)

	if sol == nil {
		log.Println("nil solution removed")
		if flag := newSol.PartIndependent(groupSize); flag {
			log.Println("better param:", newSol.CountParameter())
			return LSPartiotionAlgorithm(gr, &newSol, groupSize, number+1)
		} else {
			log.Println("invalid disb for:", number)
			return LSPartiotionAlgorithm(gr, nil, groupSize, number+1)
		}
	}
	if mark < sol.CountParameter() {
		log.Println("better mark for :", sol.Vector)
		if flag := newSol.PartIndependent(groupSize); flag {
			if newSol.CountParameter() < sol.CountParameter() {
				log.Println("better param:", newSol.Value)
				return LSPartiotionAlgorithm(gr, &newSol, groupSize, number+1)
			} else {
				log.Println("low param for:", number, " new param:", newSol.Value, " old param:", sol.Value)
			}
		} else {
			log.Println("invalid disb for:", number)
		}
	} else {
		log.Println("low mark for:", number)
	}
	return LSPartiotionAlgorithm(gr, sol, groupSize, number+1)

}

//CheckPartition checking best value async
func CheckPartition(graph IGraph, sol *Solution, groupSize int, number int64) *Solution {
	log.Println("Check number:", number)

	var newSol Solution

	log.Println("solution constructed")

	newSol.Init(graph)
	newSol.SetDependentAsBinnary(number)
	mark := newSol.CountMark()
	log.Println("mark:", mark)

	if sol == nil {
		log.Println("nil solution removed")
		if flag := newSol.PartIndependent(groupSize); flag {
			log.Println("better param:", newSol.CountParameter())
			sol = &newSol
		} else {
			log.Println("invalid disb for:", number)
			sol = nil
		}
	} else {
		if mark < sol.CountParameter() {
			log.Println("better mark for :", sol.Vector)
			if flag := newSol.PartIndependent(groupSize); flag {
				if newSol.CountParameter() < sol.CountParameter() {
					log.Println("better param:", newSol.Value)
					sol = &newSol
				} else {
					log.Println("low param for:", number, " new param:", newSol.Value, " old param:", sol.Value)
				}
			} else {
				log.Println("invalid disb for:", number)
			}
		} else {
			log.Println("low mark for:", number)
		}
	}
	return sol
}

//AsyncCheckPartitionInRange checkes best baption param in range from @start to @end
func AsyncCheckPartitionInRange(start int64, end int64, val *SafeValue, wg *sync.WaitGroup, ch chan *Solution,
	graph IGraph, groupSize int) {
	log.Println("start new goroutine")
	defer wg.Done()

	var sol *Solution

	for start <= end {

		sol = CheckPartition(graph, sol, groupSize, start)
		start++

		if sol != nil {

			val.Mux.Lock()

			if sol.Value < val.Value {
				val.Value = sol.Value
			}

			val.Mux.Unlock()
		}

	}
	ch <- sol

}

func CheckPartitionInRange(start int64, end int64, graph IGraph, groupSize int) *Solution {
	var sol *Solution
	for start <= end {
		sol = CheckPartition(graph, sol, groupSize, start)
		start++
	}
	return sol
}

//LSPartiotionAlgorithmNonRec non recursive variation of LSpartition algorithm
func LSPartiotionAlgorithmNonRec(gr IGraph, sol *Solution, groupSize int) *Solution {

	var it int64

	for it = 0; it < int64(math.Pow(2, float64(gr.AmountOfVertex()-gr.GetAmountOfIndependent()))); it++ {
		newSol := new(Solution)
		log.Println("Check number:", it)

		newSol.Init(gr)
		log.Println("solution constructed")
		newSol.SetDependentAsBinnary(it)
		mark := newSol.CountMark()

		log.Println("mark:", mark)

		if sol == nil {
			log.Println("nil solution removed")
			if flag := newSol.PartIndependent(groupSize); flag {
				log.Println("better param:", newSol.CountParameter())
				sol = newSol
				continue
			} else {
				log.Println("invalid disb for:", it)
				continue
			}
		}
		if mark < sol.CountParameter() {
			log.Println("better mark for :", sol.Vector)
			if flag := newSol.PartIndependent(groupSize); flag {
				if newSol.CountParameter() < sol.CountParameter() {
					log.Println("better param:", newSol.Value)
					sol = newSol
					continue
				} else {
					log.Println("low param for:", it, " new param:", newSol.Value, " old param:", sol.Value)
				}
			} else {
				log.Println("invalid disb for:", it)
			}
		} else {
			log.Println("low mark for:", it)
		}
	}
	return sol
}

//LSPartiotionAlgorithmCountStatistic only for statistic counting
func LSPartiotionAlgorithmCountStatistic(gr IGraph, sol *Solution, groupSize int, best int64) *Solution {
	f, _ := os.Create("trash")
	log.SetOutput(f)
	var it int64

	Statistic.m[AmountOfItterations] += int64(math.Pow(2, float64(gr.AmountOfVertex()-gr.GetAmountOfIndependent())))

	for it = 0; it < int64(math.Pow(2, float64(gr.AmountOfVertex()-gr.GetAmountOfIndependent()))); it++ {
		newSol := new(Solution)
		log.Println("Check number:", it)

		newSol.Init(gr)
		log.Println("solution constructed")
		newSol.SetDependentAsBinnary(it)
		mark := newSol.CountMark()

		log.Println("mark:", mark)

		if sol == nil {
			log.Println("nil solution removed")
			if flag := newSol.PartIndependent(groupSize); flag {
				log.Println("better param:", newSol.CountParameter())
				sol = newSol
				continue
			} else {
				log.Println("invalid disb for:", it)
				continue
			}
		}
		Statistic.m[AmountOfMarkCount]++
		if mark < sol.CountParameter() {
			Statistic.m[AmountOfFalseMark]++
			log.Println("better mark for :", sol.Vector)
			if flag := newSol.PartIndependent(groupSize); flag {
				if newSol.CountParameter() < sol.CountParameter() {
					log.Println("better param:", newSol.Value)
					sol = newSol
					continue
				} else {
					Statistic.m[AmountOfParamFail]++
					log.Println("low param for:", it, " new param:", newSol.Value, " old param:", sol.Value)
				}
			} else {
				Statistic.m[AmountOfDisbalanceFail]++
				log.Println("invalid disb for:", it)
			}
		} else {
			der := mark - sol.Value
			if der <= 1 {
				Statistic.m[MarkOneDerivative]++
			}
			derFromBest := mark - best
			Statistic.m[OverallMarkDerivativeFromBest] += derFromBest
			if derFromBest <= 5 {
				Statistic.m[AmountOfMarkDerivativeIn0To5]++
			}
			Statistic.m[OverallMarkDerivative] += der
			Statistic.m[AmountOfTrueMark]++
			log.Println("low mark for:", it)
		}
	}
	return sol
}

//LSPartiotionAlgorithmNonRecFast fastest non rec variation
func LSPartiotionAlgorithmNonRecFast(gr IGraph, sol *Solution, groupSize int) *Solution {
	var it int64

	if sol != nil && sol.Value == -1 {
		log.Panic(errors.New("Value is not inited in start solution"))
	}

	for it = 0; it < int64(math.Pow(2, float64(gr.AmountOfVertex()-gr.GetAmountOfIndependent()))); it++ {
		newSol := new(Solution)

		newSol.Init(gr)
		newSol.SetDependentAsBinnary(it)
		mark := newSol.CountMark()

		if sol == nil {
			if flag := newSol.PartIndependent(groupSize); flag {
				sol = newSol
				sol.CountParameter()
				continue
			} else {
				continue
			}
		}
		if mark < sol.Value {
			if flag := newSol.PartIndependent(groupSize); flag {
				if newSol.CountParameter() < sol.Value {
					sol = newSol
				}
			}
		}
	}
	return sol
}
