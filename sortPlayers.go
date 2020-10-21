package main

import (
	"fmt"
	"sort"
)

type player struct {
	Nome    string
	Points  float64
	Assists float64
	Rebound float64
}

//Points
type ordenaPorPoints []player

func (x ordenaPorPoints) Len() int {
	return len(x)
}
func (x ordenaPorPoints) Less(i, j int) bool {
	return x[i].Points > x[j].Points
}

func (x ordenaPorPoints) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

//Assists
type ordenaPorAssists []player

func (y ordenaPorAssists) Len() int {
	return len(y)
}

func (y ordenaPorAssists) Less(k, l int) bool {
	return y[k].Assists > y[l].Assists
}

func (y ordenaPorAssists) Swap(k, l int) {
	y[k], y[l] = y[l], y[k]
}

//Rebound
type ordenaPorRebound []player

func (z ordenaPorRebound) Len() int {
	return len(z)
}

func (z ordenaPorRebound) Less(m, n int) bool {
	return z[m].Rebound > z[n].Rebound
}

func (z ordenaPorRebound) Swap(m, n int) {
	z[m], z[n] = z[n], z[m]
}

func main() {
	players := []player{
		player{"Karl Malone", 36928, 5248, 14968},
		player{"LeBron James", 34241, 9346, 9405},
		player{"Kareem Abdul-Jabbar", 38387, 5660, 17440},
	}

	fmt.Println(players)
	fmt.Println("")

	sort.Sort(ordenaPorPoints(players))
	fmt.Println(players)

	sort.Sort(ordenaPorAssists(players))
	fmt.Println(players)

	sort.Sort(ordenaPorRebound(players))
	fmt.Println(players)
}
