package bipartitonlocalsearchlib

import graphlib "github.com/Rakiiii/goGraph"

type Graph struct {
	graphlib.Graph
	renomberedSet       []int
	amountOfIndependetn []int
}
