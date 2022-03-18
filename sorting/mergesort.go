package main

import (
	"algopractice/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	helpers.GenerateInput()
	fp := "input.txt"
	f, err := os.Open(fp)
	helpers.Check(err)

	r := bufio.NewScanner(f)
	var input = make([]int, 0)

	for r.Scan() {
		i, err := strconv.Atoi(r.Text())
		helpers.Check(err)
		input = append(input, i)
	}
	MergeSort(input, 0, 99999)
	fmt.Println(input)
}

func MergeSort(input []int, low int, high int) {
	if low < high {
		middle := (low + high) / 2
		MergeSort(input, low, middle)
		MergeSort(input, middle+1, high)

		merge(input, low, middle, high)
	}
}

func merge(input []int, low int, middle int, high int) {
	var i int

	var buf1 = make([]int, 0)
	var buf2 = make([]int, 0)

	for i = low; i <= middle; i++ {
		buf1 = append(buf1, input[i])
	}

	for i = middle + 1; i <= high; i++ {
		buf2 = append(buf2, input[i])
	}

	i = low

	for !(len(buf1) == 0 || len(buf2) == 0) {
		if buf1[0] <= buf2[0] {
			input[i] = buf1[0]
			i++
			buf1 = buf1[1:]
		} else {
			input[i] = buf2[0]
			i++
			buf2 = buf2[1:]
		}
	}

	for len(buf1) != 0 {
		input[i] = buf1[0]
		i++
		buf1 = buf1[1:]
	}

	for len(buf2) != 0 {
		input[i] = buf2[0]
		i++
		buf2 = buf2[1:]
	}

}
