package bipartitonlocalsearchlib

import graphlib "github.com/Rakiiii/goGraph"

type Graph struct {
	graphlib.Graph
	amountOfIndependent int
}

func (g *Graph) GetAmountOfIndependent() int {
	return g.amountOfIndependent
}

func (g *Graph) NumIndependent() []int {
	colorSet := make([]bool, g.AmountOfVertex())
	amountOfIndependent := 0

	for i := 0; i < g.AmountOfVertex(); i++ {
		if !colorSet[i] {
			amountOfIndependent++
			for _, v := range g.GetEdges(i) {
				colorSet[v] = true
			}
		}
	}

	newOrd := make([]int, g.AmountOfVertex())
	p1 := 0
	p2 := amountOfIndependent
	for i, flag := range colorSet {
		if flag {
			newOrd[p2] = i
			p2++

		} else {
			newOrd[p1] = i
			p1++
		}
	}

	g.amountOfIndependent = amountOfIndependent

	g.RenumVertex(newOrd)
	return newOrd
}

func (g *Graph) ParseGraph(path string) error {
	var parser graphlib.Parser

	gr, err := parser.ParseUnweightedUndirectedGraphFromFile(path)
	if err != nil {
		return err
	}

	for i := 0; i < gr.AmountOfVertex(); i++ {
		g.AddVertexWithEdges(gr.GetEdges(i))
	}
	return nil
}
