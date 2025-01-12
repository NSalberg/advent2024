package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func soln1(input string) int {
	var blockString []int

	//Create block string
	for i, s := range input {
		l := int(s - '0')

		for range l {
			if i%2 == 0 {
				blockString = append(blockString, (i / 2))
			} else {
				blockString = append(blockString, -1)
			}
		}

	}
	//fmt.Println(blockString)

	l, r := 0, len(blockString)-1

	total := 0
	for l <= r {
		if blockString[r] == -1 {
			r--
			continue
		}

		if blockString[l] == -1 {
			blockString[l] = blockString[r]
			blockString[r] = -1
			r--

		}
		val := blockString[l]
		total += l * val
		l++
	}

	return total
}

func soln2(input string) int {
	input = strings.TrimSpace(input)
	var blockString []int
	var heaps [10]IntHeap

	for i := range heaps {
		heap.Init(&heaps[i])
	}

	//Create block string
	for i, s := range input {
		l := int(s - '0')

		if i%2 == 0 {
			for range l {
				blockString = append(blockString, (i / 2))
			}
		} else {
			heap.Push(&heaps[l], len(blockString))
			for range l {
				blockString = append(blockString, -1)
			}
		}

	}

	//fmt.Println(heaps)
	//fmt.Println(blockString)

	r := len(blockString) - 1
	for 0 < r {
		fileID := blockString[r]
		if blockString[r] == -1 {
			r--
			continue
		}

		fileWidth := 0

		//fmt.Print(blockString[r], " ")
		for r >= 0 && blockString[r] == fileID {
			fileWidth += 1
			//fmt.Print(r)
			r--
		}
		//fmt.Println(fileID, fileWidth, r)

		if fileWidth >= 1 {
			minIdx := len(blockString)
			minH := -1
			for i := fileWidth; i <= 9; i++ {
				if len(heaps[i]) > 0 && heaps[i][0] < minIdx {
					minH = i
					minIdx = heaps[i][0]
				}
			}
			//fmt.Println("minH, minH-fileWidth, heaps: ", minH, minH-fileWidth, heaps)

			if minH != -1 && minIdx < r {
				minIdx = heap.Pop(&heaps[minH]).(int)
				for k := range fileWidth {
					blockString[minIdx] = fileID
					blockString[r+k+1] = -1
					minIdx++
				}
				heap.Push(&heaps[minH-fileWidth], minIdx)
				//fmt.Println(fileID, fileWidth, heaps[minH], minH-fileWidth, heaps)

				//fmt.Println(blockString)
			}
		}
	}

	total := 0
	for i, n := range blockString {
		if n != -1 {
			total += i * n
		}

	}
	return total
}

func main() {

	file, err := os.ReadFile("../../inputs/day9.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println(soln1(input))
	fmt.Println(soln2(input))

}
