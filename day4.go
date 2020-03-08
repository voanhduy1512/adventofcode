package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func day4() {
	file, err := os.Open("input4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	rangeFrom, rangeTo := 0, 0
	fmt.Sscanf(scanner.Text(), "%d-%d", &rangeFrom, &rangeTo)
	// fmt.Println(rangeFrom, rangeTo)

	// fmt.Println(countValidPassword(rangeFrom, rangeTo))
	// fmt.Println(countValidPasswordExtraRule(rangeFrom, rangeTo))

	re := regexp.MustCompile(`((.){2})`)
	fmt.Println(re.FindAllStringSubmatch("111981", -1))
}
func countValidPassword(from, to int) int {
	num := 0
	for i := from; i <= to; i++ {
		if isValidPassword(i) {
			num++
		}
	}
	return num
}
func isValidPassword(in int) bool {
	inStr := strconv.Itoa(in)

	c0, c1, c2, c3, c4, c5 := 0, 0, 0, 0, 0, 0
	fmt.Sscanf(inStr, "%1d%1d%1d%1d%1d%1d", &c0, &c1, &c2, &c3, &c4, &c5)
	if (c0 == c1 || c1 == c2 || c2 == c3 || c3 == c4 || c4 == c5) &&
		(c0 <= c1 && c1 <= c2 && c2 <= c3 && c3 <= c4 && c4 <= c5) {
		return true
	}
	return false
}
func countValidPasswordExtraRule(from, to int) int {
	num := 0
	for i := from; i <= to; i++ {
		if isValidPasswordExtraRule(i) {
			num++
		}
	}
	return num
}
func isValidPasswordExtraRule(in int) bool {
	inStr := strconv.Itoa(in)

	c0, c1, c2, c3, c4, c5 := 0, 0, 0, 0, 0, 0
	fmt.Sscanf(inStr, "%1d%1d%1d%1d%1d%1d", &c0, &c1, &c2, &c3, &c4, &c5)
	if (c0 == c1 || c1 == c2 || c2 == c3 || c3 == c4 || c4 == c5) &&
		(c0 <= c1 && c1 <= c2 && c2 <= c3 && c3 <= c4 && c4 <= c5) {
		if c0 == c1 && c1 < c2 || c0 < c1 && c1 == c2 && c2 < c3 || c1 < c2 && c2 == c3 && c3 < c4 || c2 < c3 && c3 == c4 && c4 < c5 || c3 < c4 && c4 == c5 {
			return true
		}
	}
	return false
}
