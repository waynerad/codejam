package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println("could not read line.")
		fmt.Println(err)
		panic("could not read line")
	}
	return string(ln)
}

func toInteger(numstr string) int {
	numbr, err := strconv.ParseInt(numstr, 10, 64)
	if err != nil {
		fmt.Println("could not convert to integer")
		fmt.Println(err)
		os.Exit(1)
	}
	return int(numbr)
}

func intToString(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func stalls(inputfile string, outputfile string) {
	fhOut, err := os.Create(outputfile)
	if err != nil {
		fmt.Printf("Could not open outputfile: %s", outputfile)
		return
	}
	defer fhOut.Close()
	fhOriginal, err := os.Open(inputfile)
	if err != nil {
		fmt.Printf("Not found: %s", inputfile)
		return
	}
	defer fhOriginal.Close()
	fhIn := bufio.NewReader(fhOriginal)
	inLine := readline(fhIn)
	testCases := toInteger(inLine)
	for testNumber := 1; testNumber <= testCases; testNumber++ {
		inLine = readline(fhIn)
		nandk := strings.Split(inLine, " ")
		numStalls := toInteger(nandk[0])
		numUsers := toInteger(nandk[1])
		clumps := make(map[int]int)
		clumps[numStalls] = 1
		maxIdx := numStalls
		cumulativeUsers := 0
		maxLRanswer := 0
		minLRanswer := 0
		for maxIdx > 0 {
			numOfClumps := clumps[maxIdx]
			if (cumulativeUsers < numUsers) && (cumulativeUsers+numOfClumps) >= numUsers {
				if (maxIdx & 1) == 1 {
					// odd rule: both (N - 1) / 2
					maxLRanswer = (maxIdx - 1) / 2
					minLRanswer = maxLRanswer
				} else {
					// even rule: (N / 2) and (N / 2) - 1
					maxLRanswer = maxIdx / 2
					minLRanswer = maxLRanswer - 1
				}
			}
			cumulativeUsers += numOfClumps
			if (maxIdx & 1) == 1 {
				// for odd: 2 of (N - 1) / 2
				if maxIdx > 1 {
					twice := (maxIdx - 1) / 2
					_, ok := clumps[twice]
					if ok {
						clumps[twice] += (2 * numOfClumps)
					} else {
						clumps[twice] = (2 * numOfClumps)
					}
				}
			} else {
				// for even: N/2 and N/2 - 1
				first := maxIdx / 2
				second := first - 1
				_, ok := clumps[first]
				if ok {
					clumps[first] += numOfClumps
				} else {
					clumps[first] = numOfClumps
				}
				if second > 0 {
					_, ok = clumps[second]
					if ok {
						clumps[second] += numOfClumps
					} else {
						clumps[second] = numOfClumps
					}
				}
			}
			delete(clumps, maxIdx)
			maxIdx = 0
			for x, _ := range clumps {
				if x > maxIdx {
					maxIdx = x
				}
			}
		}
		result := intToString(maxLRanswer) + " " + intToString(minLRanswer)
		_, err = fmt.Fprintf(fhOut, "Case #%d: %s\n", testNumber, result)
		if err != nil {
			fmt.Println("Error on output")
		}
	}
	fmt.Println("Done")
}

func main() {
	fileBase := "C-large"
	inputfile := fileBase + ".in"
	outputfile := fileBase + ".out"
	stalls(inputfile, outputfile)
}
