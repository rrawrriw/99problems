package main

import (
	"fmt"
	"sort"
)

type (
	freqStats map[int]int
	sortData  struct {
		data  [][]string
		stats freqStats
	}
)

func main() {
	in := [][]string{{"a", "b", "c"}, {"d", "e"}, {"f", "g", "h"}, {"d", "e"}, {"i", "j", "k", "l"}, {"m", "n"}, {"o"}}
	sortByFreq(in)
	fmt.Println(in)

}

func sortByFreq(in [][]string) {
	freqStats := newFreq(in)
	runSort(in, freqStats)
}

func newFreq(in [][]string) freqStats {
	d := make(freqStats)
	for _, e := range in {
		if _, ok := d[len(e)]; !ok {
			d[len(e)] = 1
			continue
		}

		d[len(e)] += 1
	}

	return d
}

func runSort(in [][]string, st freqStats) {
	d := sortData{
		data:  in,
		stats: st,
	}
	sort.Sort(d)
}

func (d sortData) Len() int {
	return len(d.data)
}

func (d sortData) Less(i, j int) bool {
	if d.stats[len(d.data[i])] < d.stats[len(d.data[j])] {
		return true
	}

	return false
}

func (d sortData) Swap(i, j int) {
	tmp := d.data[j]
	d.data[j] = d.data[i]
	d.data[i] = tmp
}
