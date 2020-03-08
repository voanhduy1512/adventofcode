package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestProcessExtendMulti(t *testing.T) {
	tests := []struct {
		inputSignal     []int
		settingSequence []string
		expectedOutput  int
	}{
		{[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, []string{"4", "3", "2", "1", "0"}, 43210},
		{[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23,
			1, 24, 23, 23, 4, 23, 99, 0, 0}, []string{"0", "1", "2", "3", "4"}, 54321},
		{[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
			1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, []string{"1", "0", "4", "3", "2"}, 65210},
	}
	for _, tt := range tests {
		temp1, temp2 := &bytes.Buffer{}, &bytes.Buffer{}
		temp2.WriteString("0")
		for _, sq := range tt.settingSequence {
			temp1.WriteString(sq)
			temp1.WriteString("\n")

			var i []byte
			fmt.Fscanln(temp2, &i)
			temp1.Write(i)

			process(tt.inputSignal, temp1, temp2)
			// fmt.Println(temp2.String())
		}
		result, _ := strconv.Atoi(temp2.String())
		if result != tt.expectedOutput {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", result, tt.expectedOutput)
		}
	}
}

func TestProcessExtendMultiLoop(t *testing.T) {
	tests := []struct {
		inputSignal     []int
		settingSequence []string
		expectedOutput  int
	}{
		{[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26,
			27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}, []string{"9", "8", "7", "6", "5"}, 139629729},
		// {[]int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
		// 	-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
		// 	53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}, []string{"9", "7", "8", "5", "6"}, 18216},
	}

	for _, tt := range tests {
		done := make(chan int)
		donez := make(chan int)
		ins := []*bytes.Buffer{&bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}}
		outs := []*bytes.Buffer{&bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}}
		var wg sync.WaitGroup
		for i, sq := range tt.settingSequence {
			wg.Add(1)
			copyIns := make([]int, len(tt.inputSignal))
			copy(copyIns, tt.inputSignal)
			if i == 0 {
				ins[i].WriteString(sq + "\n0")
			} else {
				ins[i].WriteString(sq + "\n")
			}
			// fmt.Println(i, ins[i], outs[i])
			go func(ii int, sq string) {
				defer wg.Done()
				process(copyIns, ins[ii], outs[ii])
			}(i, sq)
		}
		go func() {
			i := 0
			for {
				select {
				case <-done:
					// fmt.Println("return", i)
					donez <- i
					return
				default:
					var v int
					if n, _ := fmt.Fscan(outs[i], &v); n == 0 {
						continue
					}
					if i = i + 1; i == len(tt.settingSequence) {
						i = 0
					}
					for {
						fmt.Printf("(%d, %d, %s)\n", v, i, ins[i].String())
						n, err := ins[i].WriteString(strconv.Itoa(v))
						if err != nil || n == 0 {
							continue
						}
						break
					}
					fmt.Printf("((%d, %s, %s))\n", i, ins[i], outs[i])
				}
			}
		}()
		wg.Wait()
		done <- 1
		xx := <-donez
		result := 0
		if xx == 4 {
			result, _ = strconv.Atoi(strings.Trim(outs[len(tt.settingSequence)-1].String(), "\n"))
		} else {
			result, _ = strconv.Atoi(strings.Trim(ins[0].String(), "\n"))
		}
		if result != tt.expectedOutput {
			fmt.Println(ins, outs)
			t.Errorf("Sum was incorrect, got: %d, want: %d.", result, tt.expectedOutput)
		}
	}
}
