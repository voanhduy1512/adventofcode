package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	file, err := os.Open("input7.txt")
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
	max := 0
	for i := 0; i <= 4; i = i + 1 {
		for j := 0; j <= 4; j = j + 1 {
			for k := 0; k <= 4; k = k + 1 {
				for l := 0; l <= 4; l = l + 1 {
					for m := 0; m <= 4; m = m + 1 {
						if i != j && i != k && i != l && i != m && j != k && j != l && j != m && k != l && k != m && l != m {
							done := make(chan int)
							donez := make(chan int)
							check := 0
							mutex := &sync.Mutex{}
							ins := []*bytes.Buffer{&bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}}
							outs := []*bytes.Buffer{&bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}}
							var wg sync.WaitGroup
							settingSequence := []string{strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(k), strconv.Itoa(l), strconv.Itoa(m)}
							// fmt.Println(settingSequence)
							for _, sq := range settingSequence {
								wg.Add(1)
								go func(ii int, sq string) {
									defer wg.Done()

									copyIns := make([]int, len(originalIns))
									copy(copyIns, originalIns)
									// ins[ii].WriteString(sq + "\n")
									if ii == 0 {
										ins[ii].WriteString(sq + "\n0")
									} else {
										ins[ii].WriteString(sq + "\n")
									}
									mutex.Lock()
									check++
									mutex.Unlock()

									// fmt.Println("check", check)
									// fmt.Println(ii, ins[ii], outs[ii])
									process(copyIns, ins[ii], outs[ii])
								}(i, sq)
							}
							go func() {
								i := 0
								for {
									// fmt.Println("checkxx", check)
									if check < len(settingSequence)-1 {
										continue
									}
									select {
									case <-done:
										// fmt.Println("return", i)
										donez <- 1
										return
									default:
										var v int
										n, err := fmt.Fscan(outs[i], &v)

										if err != nil && n == 0 {
											continue
										}
										i = i + 1
										if i == len(settingSequence) {
											i = 0
										}
										// fmt.Println(strconv.Itoa(v) + "\n")
										ins[i].WriteString(strconv.Itoa(v))
										continue
									}
								}
							}()
							wg.Wait()
							done <- 1
							<-donez
							// result, _ := strconv.Atoi(strings.Trim(outs[len(settingSequence)-1].String(), "\n"))
							result1, _ := strconv.Atoi(strings.Trim(ins[0].String(), "\n"))
							if max < result1 {
								fmt.Println(settingSequence)
								max = result1
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(max)
}

func part1() {
	file, err := os.Open("input7.txt")
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
	max := 0
	for i := 0; i <= 4; i = i + 1 {
		for j := 0; j <= 4; j = j + 1 {
			for k := 0; k <= 4; k = k + 1 {
				for l := 0; l <= 4; l = l + 1 {
					for m := 0; m <= 4; m = m + 1 {
						if i != j && i != k && i != l && i != m && j != k && j != l && j != m && k != l && k != m && l != m {
							temp1, temp2 := &bytes.Buffer{}, &bytes.Buffer{}
							temp2.WriteString("0")
							settingSequence := []string{strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(k), strconv.Itoa(l), strconv.Itoa(m)}
							// fmt.Println(settingSequence)
							for _, sq := range settingSequence {
								temp1.WriteString(sq)
								temp1.WriteString("\n")

								var ii []byte
								fmt.Fscanln(temp2, &ii)
								temp1.Write(ii)

								process(originalIns, temp1, temp2)
								// fmt.Println(temp2.String())

								result, _ := strconv.Atoi(temp2.String())
								if max < result {
									fmt.Println(settingSequence)
									max = result
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(max)
}
