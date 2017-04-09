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

func stringToByte(st string) byte {
	return st[0]
}

func intToString(x int) string {
	return strconv.FormatInt(int64(x), 10)
}

func uintToBinary(i uint64) string {
	return strconv.FormatUint(i, 2)
}

func flippancakes(inputfile string, outputfile string) {
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
	plus := stringToByte("+")
	var bignum uint64
	bignum = 1 << 63
	for testNumber := 1; testNumber <= testCases; testNumber++ {
		inLine = readline(fhIn)
		lineStr := strings.Split(inLine, " ")
		rowOfPanckaes := lineStr[0]
		conseq := toInteger(lineStr[1])
		// we actually load the pancakes backwards because forwards and backwards seem to be mathematically equivalent
		x := len(rowOfPanckaes)
		s := (x / 64) + 1
		var bit uint64
		panface := make([]uint64, s)
		for i := 0; i < s; i++ {
			panface[i] = 0
		}
		mask := make([]uint64, s)
		for i := 0; i < s; i++ {
			mask[i] = 0
		}
		bit = 1
		byte := 0
		for i := 0; i < x; i++ {
			if rowOfPanckaes[i] != plus {
				// 1 means NOT happy -- algorithm's goal is to get to all 0s
				panface[byte] |= bit
			}
			if i < conseq {
				mask[byte] |= bit
			}
			bit <<= 1
			if bit == 0 {
				byte++
				bit = 1
			}
		}
		// reset bit to use as test bit going forward
		bit = 1
		byte = 0
		steps := x - conseq
		maskApplications := 0
		for i := 0; i <= steps; i++ {
			// step 1 - test
			// first we have to rewind the test bit because it starts out of alignment with the mask
			// test test bit
			if (panface[byte] & bit) != 0 {
				// apply mask
				for k := 0; k < s; k++ {
					if mask[k] != 0 {
						panface[k] ^= mask[k]
					}
				}
				maskApplications++
			}
			// shift test bit
			if bit == bignum {
				byte++
				bit = 1
			} else {
				bit <<= 1
			}
			// shift mask
			carry := false
			for k := 0; k < s; k++ {
				nextCarry := false
				z := mask[k]
				if (z & bignum) != 0 {
					nextCarry = true
				}
				z <<= 1
				if carry {
					z |= 1
				}
				carry = nextCarry
				mask[k] = z
			}
		}
		impossible := false
		for i := 0; i < s; i++ {
			if panface[i] != 0 {
				impossible = true
			}
		}
		var result string
		if impossible {
			result = "IMPOSSIBLE"
		} else {
			result = intToString(maskApplications)
		}
		_, err = fmt.Fprintf(fhOut, "Case #%d: %s\n", testNumber, result)
		if err != nil {
			fmt.Println("Error on output")
		}
	}
	fmt.Println("Done")
}

func main() {
	fileBase := "A-large"
	inputfile := fileBase + ".in"
	outputfile := fileBase + ".out"
	flippancakes(inputfile, outputfile)
}
