package lspartitioninglib

import (
	"sync"
	"log"
	"math"
)

type SafeValue struct {
	Value int64
	Mux   sync.Mutex
}

func LSPartiotionAlgorithm(gr *Graph, sol *Solution, groupSize int, number int64) *Solution {
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

func CheckPartition(graph *Graph, sol *Solution, groupSize int, number int64) *Solution {
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

func AsyncCheckPartitionInRange(start int64, end int64, val *SafeValue, wg *sync.WaitGroup, ch chan *Solution,
	graph *Graph, groupSize int) {
		log.Println("start new goroutine")
	defer wg.Done()
	
	var sol *Solution 
	
	for start <= end {

		sol = CheckPartition(graph, sol, groupSize, start)
		start++

		if sol != nil{
	
			val.Mux.Lock()

		if sol.Value < val.Value {
			val.Value = sol.Value
		}

		val.Mux.Unlock()
	}

	}
	ch <- sol

}

func CheckPartitionInRange(start int64, end int64, graph *Graph, groupSize int) *Solution {
	var sol *Solution
	for start <= end {
		sol = CheckPartition(graph, sol, groupSize, start)
		start++
	}
	return sol
}

func LSPartiotionAlgorithmNonRec(gr *Graph, sol *Solution, groupSize int) *Solution {
	log.Println("Check number:", number)

	var newSol Solution
	solFlag := false

	log.Println("solution constructed")

	for start < end;start++{

		newSol.Init(gr)
		newSol.SetDependentAsBinnary(start)
		mark := newSol.CountMark()
		log.Println("mark:", mark)

		if solFlag{
			log.Println("nil solution removed")
			if flag := newSol.PartIndependent(groupSize); flag {
				log.Println("better param:", newSol.CountParameter())
				solFlag = true
				sol = &newSol
				continue
			} else {
				log.Println("invalid disb for:", number)
				continue
			}
		}
		if mark < sol.CountParameter() {
			log.Println("better mark for :", sol.Vector)
			if flag := newSol.PartIndependent(groupSize); flag {
				if newSol.CountParameter() < sol.CountParameter() {
					log.Println("better param:", newSol.Value)
					sol = &newSol
					continue
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
}