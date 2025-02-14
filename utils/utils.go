package utils

import (
	"cmp"
	"container/heap"
	"iter"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func InputByLines(file string) iter.Seq[string] {
	file = filepath.FromSlash(file)

	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Unable to read file: %s", err)
	}
	return strings.Lines(string(b))
}

// New generic type that accounts for all int types in Go
type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Abs[T Int](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

type MinHeap[T cmp.Ordered] struct {
	minHeap[T]
}

// New Pop so we can have a type returned instead of any which is required by heap.Interface
func (h *MinHeap[T]) Pop() T {
	return heap.Pop(&h.minHeap).(T)
}

// New Push so we can have a type returned instead of any which is required by heap.Interface
func (h *MinHeap[T]) Push(x T) {
	heap.Push(&h.minHeap, x)
}

type minHeap[T cmp.Ordered] []T

var _ heap.Interface = (*minHeap[int])(nil)

func (h minHeap[T]) Len() int {
	return len(h)
}

func (h minHeap[T]) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *minHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
