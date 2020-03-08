package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func day2() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	ins := strings.Split(scanner.Text(), ",")

	originalIns := make([]int, len(ins))
	for i := range originalIns {
		originalIns[i], _ = strconv.Atoi(ins[i])
	}
	copyIns := make([]int, len(originalIns))
	for i := 0; i < len(originalIns); i++ {
		for j := 0; j < len(originalIns); j++ {
			copy(copyIns, originalIns)
			copyIns[1], copyIns[2] = i, j
			process(copyIns, os.Stdin, os.Stdout)
			if copyIns[0] == 19690720 {
				fmt.Println(copyIns)
				return
			}
		}
	}
}
func process(ins []int, r io.Reader, w io.Writer) {
	x, y, step := 0, 0, 0
	done := false
	for {
		if done {
			return
		}
		x = y
		opcode := getOpcode(ins[x])

		if opcode == 1 || opcode == 2 || opcode == 7 || opcode == 8 {
			step = 4
		} else if opcode == 5 || opcode == 6 {
			step = 3
		} else {
			step = 2
		}
		y = y + step
		if y > len(ins) {
			y = len(ins)
		}
		done = processIntcode(ins, ins[x:y], &y, r, w)
	}
}
func getOpcode(in int) int {
	return int(math.Mod(float64(in), float64(100)))
}
func getPositionMode(in int) (out []int) {
	postModeStr := strconv.Itoa(in / 100)
	for i := len(postModeStr) - 1; i >= 0; i-- {
		tmp, _ := strconv.Atoi(string(postModeStr[i]))
		out = append(out, tmp)
	}
	return
}
func processIntcode(ins, ins4 []int, nextPoint *int, r io.Reader, w io.Writer) bool {
	// fmt.Println(ins)
	// fmt.Println(ins4)
	opcode := getOpcode(ins4[0])
	positionModes := getPositionMode(ins4[0])
	positions := make([]int, 0)
	vals := make([]int, 1)
	vals[0] = 0
	switch opcode {
	case 1:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			if i-1 < len(positionModes) {
				if positionModes[i-1] == 0 {
					vals = append(vals, ins[positions[i]])
				} else {
					vals = append(vals, positions[i])
				}
			} else {
				vals = append(vals, ins[positions[i]])
			}
		}
		ins[positions[3]] = vals[1] + vals[2]
	case 2:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			if i-1 < len(positionModes) {
				if positionModes[i-1] == 0 {
					vals = append(vals, ins[positions[i]])
				} else {
					vals = append(vals, positions[i])
				}
			} else {
				vals = append(vals, ins[positions[i]])
			}
		}
		ins[positions[3]] = vals[1] * vals[2]
	case 3:
		pos1 := ins4[1]
		for {
			var i int
			n, _ := fmt.Fscan(r, &i)
			if n == 0 {
				continue
			}
			fmt.Printf("in=%d\n", i)
			ins[pos1] = i
			break
		}
	case 4:
		pos1 := ins4[1]
		var val int
		if positionModes[0] == 1 {
			val = pos1
		} else {
			val = ins[pos1]
		}
		for {
			fmt.Printf("out=%d\n", val)
			n, err := fmt.Fprintf(w, "%d", val)
			if n == 0 || err != nil {
				continue
			}
			break
		}
	case 5:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			if i-1 < len(positionModes) {
				if positionModes[i-1] == 0 {
					vals = append(vals, ins[positions[i]])
				} else {
					vals = append(vals, positions[i])
				}
			} else {
				vals = append(vals, ins[positions[i]])
			}
		}
		if vals[1] != 0 {
			*nextPoint = vals[2]
		}
	case 6:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			if i-1 < len(positionModes) {
				if positionModes[i-1] == 0 {
					vals = append(vals, ins[positions[i]])
				} else {
					vals = append(vals, positions[i])
				}
			} else {
				vals = append(vals, ins[positions[i]])
			}
		}
		if vals[1] == 0 {
			*nextPoint = vals[2]
		}
	case 7:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			if i-1 < len(positionModes) {
				if positionModes[i-1] == 0 {
					vals = append(vals, ins[positions[i]])
				} else {
					vals = append(vals, positions[i])
				}
			} else {
				vals = append(vals, ins[positions[i]])
			}
		}
		if vals[1] < vals[2] {
			ins[positions[3]] = 1
		} else {
			ins[positions[3]] = 0
		}
	case 8:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			if i-1 < len(positionModes) {
				if positionModes[i-1] == 0 {
					vals = append(vals, ins[positions[i]])
				} else {
					vals = append(vals, positions[i])
				}
			} else {
				vals = append(vals, ins[positions[i]])
			}
		}
		if vals[1] == vals[2] {
			ins[positions[3]] = 1
		} else {
			ins[positions[3]] = 0
		}
	case 99:
		return true
	default:
		return true
	}
	return false
}
