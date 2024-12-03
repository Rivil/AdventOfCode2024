package puzzle1

import (
	"bufio"
	"os"
	"slices"
	"strconv"
)

func Puzzle1() int64 {
	list1, list2 := getLists()
	var total int64 = 0
	for i := 0; i < len(list1); i++ {
		num := list1[i] - list2[i]
		if num < 0 {
			num = num * -1
		}
		total += num
	}
	return total
}

func Puzzle2() int64 {
	list1, list2 := getLists()
	var total int64 = 0
	for i := 0; i < len(list1); i++ {
		iter := 0
		pos, found := slices.BinarySearch(list2, list1[i])
		if found {
			for j := pos; j < len(list2); j++ {
				if list2[j] == list1[i] {
					iter++
				} else {
					break
				}
			}
			total += list1[i] * int64(iter)
		}
	}
	return total
}

func getLists() ([]int64, []int64) {
	file := "puzzle1/input.txt"
	open, err := os.Open(file)
	if err != nil {
		panic(err)
		return nil, nil
	}
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			panic(err)
		}
	}(open)
	scanner := bufio.NewScanner(open)
	scanner.Split(bufio.ScanWords)
	list1 := make([]int64, 0)
	list2 := make([]int64, 0)
	word := 0
	for scanner.Scan() {
		word++
		parseInt, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
			return nil, nil
		}
		if word%2 == 1 {
			list1 = append(list1, parseInt)
		} else {
			list2 = append(list2, parseInt)
		}
	}
	slices.Sort(list1)
	slices.Sort(list2)
	return list1, list2
}
