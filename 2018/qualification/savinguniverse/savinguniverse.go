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
		fmt.Fprintln(os.Stderr, "could not read line.")
		fmt.Fprintln(os.Stderr, err)
		panic("could not read line")
	}
	return string(ln)
}

func toInteger(numstr string) int {
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

func removeTrailingCs(programString string) string {
	lx := len(programString)
	b := programString[lx-1]
	if b != 67 {
		return programString
	}
	result := programString
	for b == 67 {
		result = result[0 : lx-1]
		lx = len(result)
		if lx != 0 {
			b = result[lx-1]
		} else {
			b = 0
		}
	}
	return result
}

func swapLastC(programString string) (string, bool) {
	last := -1
	lx := len(programString)
	for i := 0; i < lx; i++ {
		if programString[i] == 67 {
			last = i
		}
	}
	if last == -1 {
		return programString, false // we failed to swap
	}
	result := programString[0:last] + programString[last+1:last+2] + programString[last:last+1] + programString[last+2:]
	return result, true
}

func calcValue(programString string) int {
	total := 0
	adder := 1
	for _, ch := range programString {
		if ch == 83 {
			// shoot
			total += adder
		}
		if ch == 67 {
			// charge
			if adder < 1073741824 {
				// we were guaranteed in the problem description D would be <= 10^9
				adder = adder * 2
			}
		}
	}
	return total
}

func savingTheUniverseAgain(fhOriginal *os.File, fhOut *os.File) {
	fhIn := bufio.NewReader(fhOriginal)
	inLine := readline(fhIn)
	testCases := toInteger(inLine)
	for testNumber := 1; testNumber <= testCases; testNumber++ {
		inLine = readline(fhIn)
		stuff := strings.Split(inLine, " ")
		damage := toInteger(stuff[0])
		programString := stuff[1] // convert to byte array?
		programString = removeTrailingCs(programString)
		numberOfSwaps := 0
		impossible := false
		for calcValue(programString) > damage {
			var didit bool
			programString, didit = swapLastC(programString)
			programString = removeTrailingCs(programString)
			numberOfSwaps++
			if !didit {
				impossible = true
				programString = "" // force exit of loop
			}
		}
		result := intToString(numberOfSwaps)
		if impossible {
			result = "IMPOSSIBLE"
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
	savingTheUniverseAgain(stdin, stdout)
}
