package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func day5() {
	file, err := os.Open("input5.txt")
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
	process(originalIns, os.Stdin, os.Stdout)
}
