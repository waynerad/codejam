// This is my solution to Googe Code Jam Round 1A 2017 problem A "Alphabet Cake"
// Read the problem description here: https://code.google.com/codejam/contest/5304486/dashboard

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

func byteToString(b byte) string {
	byteArray := make([]byte, 1)
	byteArray[0] = b
	return string(byteArray)
}

type point struct {
	x int
	y int
}

func alphabetcake(inputfile string, outputfile string) {
	currentTime := time.Now()
	currentUnix := currentTime.UnixNano()
	rand.Seed(currentUnix)
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
		stuff := strings.Split(inLine, " ")
		rows := toInteger(stuff[0])
		cols := toInteger(stuff[1])
		grid := make([][]byte, rows)
		originalLines := make([]string, rows)
		for i := 0; i < rows; i++ {
			grid[i] = make([]byte, cols)
		}
		// we have to save the original lines in case we need to start over
		for i := 0; i < rows; i++ {
			originalLines[i] = readline(fhIn)
		}
		// set to true so it will do initial round -- plus it's true, there are ?s
		foundqs := true
		for foundqs {
			for i := 0; i < rows; i++ {
				inLine = originalLines[i]
				for j := 0; j < cols; j++ {
					grid[i][j] = inLine[j]
				}
			}
			ptlist := make(map[byte][]point)
			// order & count are for randomization
			order := make([]byte, 0)
			count := 0
			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					if grid[i][j] != 63 { // ?
						kid := grid[i][j]
						_, ok := ptlist[kid]
						if !ok {
							ptlist[kid] = make([]point, 1)
							ptlist[kid][0].x = j
							ptlist[kid][0].y = i
							order = append(order, kid)
							count++
						} else {
							var newpt point
							newpt.x = j
							newpt.y = i
							ptlist[kid] = append(ptlist[kid], newpt)
						}
					}
				}
			}
			// randomize order
			for z := 0; z < count; z++ {
				n := int(rand.Int31n(int32(count)))
				// swap
				k := order[z]
				order[z] = order[n]
				order[n] = k
			}
			// for kid, starting := range ptlist { -- we're replacing this walk through the list with our randomized order
			for z := 0; z < count; z++ {
				kid := order[z]
				starting := ptlist[kid]
				box := make([]point, 2)
				first := true
				for _, pt := range starting {
					if first {
						// initialize to both corners same point
						box[0].x = pt.x
						box[1].x = pt.x
						box[0].y = pt.y
						box[1].y = pt.y
						first = false
					} else {
						if pt.x < box[0].x {
							box[0].x = pt.x
						}
						if pt.x > box[1].x {
							box[1].x = pt.x
						}
						if pt.y < box[0].y {
							box[0].y = pt.y
						}
						if pt.y > box[1].y {
							box[1].y = pt.y
						}
					}
				}
				keepGoing := true
				for keepGoing {
					keepGoing = false
					// can we go up?
					if box[0].y > 0 {
						allQs := true
						for i := box[0].x; i <= box[1].x; i++ {
							if grid[box[0].y-1][i] != 63 {
								allQs = false
							}
						}
						if allQs {
							box[0].y--
							keepGoing = true
						}
					}
					// can we go left?
					if box[0].x > 0 {
						allQs := true
						for i := box[0].y; i <= box[1].y; i++ {
							if grid[i][box[0].x-1] != 63 {
								allQs = false
							}
						}
						if allQs {
							box[0].x--
							keepGoing = true
						}
					}
					// can we go down
					if box[1].y < rows-1 {
						allQs := true
						for i := box[0].x; i <= box[1].x; i++ {
							if grid[box[1].y+1][i] != 63 {
								allQs = false
							}
						}
						if allQs {
							box[1].y++
							keepGoing = true
						}
					}
					// can we go right
					if box[1].x < cols-1 {
						allQs := true
						for i := box[0].y; i <= box[1].y; i++ {
							if grid[i][box[1].x+1] != 63 {
								allQs = false
							}
						}
						if allQs {
							box[1].x++
							keepGoing = true
						}
					}
				}
				// fill
				for i := box[0].y; i <= box[1].y; i++ {
					for j := box[0].x; j <= box[1].x; j++ {
						grid[i][j] = kid
					}
				}
			}
			// check
			foundqs = false
			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					if grid[i][j] == 63 {
						foundqs = true
					}
				}
			}
			if foundqs {
				// _, err = fmt.Fprintf(fhOut, "Case #%d: blew up because of a question mark\n", testNumber)
				// panic("found a question mark!!!")
			}
		}
		_, err = fmt.Fprintf(fhOut, "Case #%d:\n", testNumber)
		for i := 0; i < rows; i++ {
			str := ""
			for j := 0; j < cols; j++ {
				str += byteToString(grid[i][j])
			}
			fmt.Fprintln(fhOut, str)
		}
		if err != nil {
			fmt.Println("Error on output")
		}
	}
	fmt.Println("Done")
}

func main() {
	fileBase := "A-large-practice"
	inputfile := fileBase + ".in"
	outputfile := fileBase + ".out"
	alphabetcake(inputfile, outputfile)
}
