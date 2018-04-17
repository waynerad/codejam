package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readline(r *bufio.Reader) string {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not read line.")
		fmt.Fprintln(os.Stderr, err)
		panic("could not read line")
	}
	return string(ln)
}

func toInt(numstr string) int {
	numbr, err := strconv.ParseInt(numstr, 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not convert to integer")
		fmt.Fprintln(os.Stderr, err)
		panic("could not convert to integer")
	}
	return int(numbr)
}

func intToString(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func checkArrays(array1 []int, array2 []int) int {
	l1 := len(array1)
	l2 := len(array2)

	idx := 0
	for jj := 0; jj < l2; jj++ {
		// step 1
		if array1[jj] > array2[jj] {
			return idx
		}
		idx++
		// step 2
		if (jj + 1) < l1 {
			if array2[jj] > array1[jj+1] {
				return idx
			}
		}
		idx++
	}
	return -1
}

func troubleSort(fhOriginal *os.File, fhOut *os.File) {
	fhIn := bufio.NewReader(fhOriginal)
	inLine := readline(fhIn)
	testCases := toInt(inLine)
	for testNumber := 1; testNumber <= testCases; testNumber++ {
		inLine = readline(fhIn)
		numberOfValues := toInt(inLine)
		inLine = readline(fhIn)
		stuff := strings.Split(inLine, " ")
		j := 0
		k := 0
		m := 0
		s := (numberOfValues / 2) + 2
		array1 := make([]int, s)
		array2 := make([]int, s)
		for i := 0; i < numberOfValues; i++ {
			if m == 0 {
				array1[j] = toInt(stuff[i])
				j++
			} else {
				array2[k] = toInt(stuff[i])
				k++
			}
			m = 1 - m
		}
		array1 = array1[0:j]
		array2 = array2[0:k]
		sort.Sort(sort.IntSlice(array1))
		sort.Sort(sort.IntSlice(array2))

		answer := checkArrays(array1, array2)

		result := intToString(answer)
		if answer < 0 {
			result = "OK"
		}
		_, err := fmt.Fprintf(fhOut, "Case #%d: %s\n", testNumber, result)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error on output")
		}
	}
}

func main() {
	stdin := os.Stdin
	stdout := os.Stdout
	troubleSort(stdin, stdout)
}
