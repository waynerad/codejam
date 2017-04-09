package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func asciiToString(ascii int) string {
	byteArray := make([]byte, 1)
	byteArray[0] = byte(ascii)
	return string(byteArray)
}

func tidy(inputfile string, outputfile string) {
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
		x := len(inLine)
		digits := make([]int, x)
		for i := 0; i < x; i++ {
			digits[i] = int(inLine[i]) - 48
		}
		position := x - 2
		for position >= 0 {
			if digits[position] > digits[position+1] {
				digits[position]--
				for k := position + 1; k < x; k++ {
					digits[k] = 9
				}
				position = x - 1
			}
			position--
		}
		result := ""
		started := false
		for i := 0; i < x; i++ {
			if digits[i] == 0 {
				if started {
					result += asciiToString(48 + digits[i])
				}
			} else {
				started = true
				result += asciiToString(48 + digits[i])
			}
		}
		_, err = fmt.Fprintf(fhOut, "Case #%d: %s\n", testNumber, result)
		if err != nil {
			fmt.Println("Error on output")
		}
	}
	fmt.Println("Done")
}

func main() {
	fileBase := "B-large"
	inputfile := fileBase + ".in"
	outputfile := fileBase + ".out"
	tidy(inputfile, outputfile)
}
