package edugraph

import (
	"math"
	"sort"
)

type Graph struct {
	Directed bool     `json:"directed"`
	Vertices []string `json:"vertices"`
	Edges    []Edge   `json:"edges"`
}

type Edge struct {
	Vertices [2]string `json:"v"`
	Cost     float64   `json:"cost"`
}

type Vertex struct {
	Name      string
	Neighbors []string

	costs map[string]float64
}

func (g *Graph) Link() map[string]*Vertex {
	vertexMap := make(map[string]*Vertex)

	for _, name := range g.Vertices {
		vertexMap[name] = &Vertex{
			Name:  name,
			costs: make(map[string]float64),
		}
	}

	for _, edge := range g.Edges {
		vertexMap[edge.Vertices[0]].Neighbors = append(vertexMap[edge.Vertices[0]].Neighbors, edge.Vertices[1])
		vertexMap[edge.Vertices[0]].costs[edge.Vertices[1]] = edge.Cost

		if !g.Directed {
			vertexMap[edge.Vertices[1]].Neighbors = append(vertexMap[edge.Vertices[1]].Neighbors, edge.Vertices[0])
			vertexMap[edge.Vertices[1]].costs[edge.Vertices[0]] = edge.Cost
		}
	}
	return vertexMap
}

func (v Vertex) Cost(to string) float64 {
	cost, ok := v.costs[to]
	if !ok {
		return math.Inf(1)
	}
	return cost
}

type EdgesSortedByIncreasingCost []Edge

func (ed EdgesSortedByIncreasingCost) Sort() {
	sort.SliceStable(ed, func(i int, j int) bool {
		return ed[i].Cost < ed[j].Cost
	})
}
