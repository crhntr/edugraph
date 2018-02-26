package edugraph

import (
	"encoding/json"
	"math"
	"testing"
)

func Test(t *testing.T) {
	g := Graph{}
	err := json.Unmarshal([]byte(simpleGraphJSON), &g)
	if err != nil {
		t.Error(err)
	}

	EdgesSortedByIncreasingCost(g.Edges).Sort()

	t.Log(g)

	vm := g.Link()

	if vm["B"].Cost("C") != math.Inf(1) {
		t.Fail()
	}
	if vm["A"].Cost("B") != 1 {
		t.Fail()
	}
}

var simpleGraphJSON = `{
  "vertices": ["A", "B", "C", "D"],
  "edges": [
    {"v": ["A", "B"], "cost": 1},
    {"v": ["B", "D"], "cost": 4},
    {"v": ["D", "C"], "cost": 2},
    {"v": ["C", "A"], "cost": 2},
    {"v": ["A", "D"], "cost": 8}
  ]
}`
