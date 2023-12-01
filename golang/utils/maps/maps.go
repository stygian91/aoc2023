package maps

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type ord = constraints.Ordered

type sortable[K ord] struct {
	Data []K
}

func (this sortable[K]) Len() int {
	return len(this.Data)
}

func (this sortable[K]) Less(i, j int) bool {
	return this.Data[i] < this.Data[j]
}

func (this *sortable[K]) Swap(i, j int) {
	swap := this.Data[i]
	this.Data[i] = this.Data[j]
	this.Data[j] = swap
}

func KeysInOrder[K ord, V any](table map[K]V) []K {
	sortable := sortable[K]{Data: []K{}}

	for k := range table {
		sortable.Data = append(sortable.Data, k)
	}

	sort.Sort(&sortable)

	return sortable.Data
}

func IterInOrder[K ord, V any](table map[K]V, cb func(K, V)) {
	keys := KeysInOrder(table)

	for _, k := range keys {
		v := table[k]
		cb(k, v)
	}
}
