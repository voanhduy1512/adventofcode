package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day1() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := int64(0)
	for scanner.Scan() {
		in, _ := strconv.Atoi(scanner.Text())
		sum = sum + calcFuel(int64(in))
	}

	fmt.Println(sum)
}
func calcFuel(in int64) int64 {
	in = int64(in/3) - 2
	if in <= 0 {
		return 0
	}
	return calcFuel(in) + in
}
