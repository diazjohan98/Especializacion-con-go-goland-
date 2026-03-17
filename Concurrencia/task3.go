package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortPartition(arr []int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Goroutine %d is sorting subarray: %v\n", id, arr)
	sort.Ints(arr)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	fmt.Println("Please enter a series of integers separated by spaces:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var numbers []int
	for _, s := range strings.Fields(input) {
		n, err := strconv.Atoi(s)
		if err == nil {
			numbers = append(numbers, n)
		}
	}

	if len(numbers) == 0 {
		fmt.Println("No valid integers provided.")
		return
	}

	n := len(numbers)
	chunkSize := n / 4
	remainder := n % 4

	var partitions [][]int
	start := 0
	for i := 0; i < 4; i++ {
		end := start + chunkSize
		if remainder > 0 {
			end++
			remainder--
		}
		partitions = append(partitions, numbers[start:end])
		start = end
	}

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go sortPartition(partitions[i], &wg, i+1)
	}

	wg.Wait()

	merged1 := merge(partitions[0], partitions[1])
	merged2 := merge(partitions[2], partitions[3])
	finalSorted := merge(merged1, merged2)

	fmt.Println("Final merged and sorted array:", finalSorted)
}
