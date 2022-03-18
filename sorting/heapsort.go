package sorting

import (
	"algopractice/datastructures"
	"algopractice/helpers"
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func HeapSort() {
	h := &datastructures.IntHeap{}
	heap.Init(h)

	fp := "input.txt"
	f, err := os.Open(fp)
	helpers.Check(err)

	defer f.Close()

	r := bufio.NewScanner(f)

	for r.Scan() {
		i, err := strconv.Atoi(r.Text())
		helpers.Check(err)
		heap.Push(h, i)
	}

	out_fp := "sorted_input.txt"
	of, err := os.Create(out_fp)
	helpers.Check(err)

	defer of.Close()

	for h.Len() > 0 {
		_, err := of.WriteString(fmt.Sprintf("%d\n", heap.Pop(h)))
		helpers.Check(err)
	}
}
