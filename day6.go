package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Object type
type Object struct {
	body string
	next *Object
}

func day6() {
	file, err := os.Open("input6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mercury := make(map[string]*Object)
	for scanner.Scan() {
		ins := strings.Split(scanner.Text(), ")")
		var objs Object
		if _, ok := mercury[ins[0]]; !ok {
			mercury[ins[0]] = &Object{ins[0], nil}
		}
		temp := mercury[ins[0]]
		if xxx, ok := mercury[ins[1]]; !ok {
			objs = Object{ins[1], temp}
			mercury[ins[1]] = &objs
		} else {
			xxx.next = temp
			mercury[ins[1]] = xxx
		}
	}
	sum := 0
	san := make([]string, 0)
	you := make([]string, 0)
	for _, m := range mercury {
		// fmt.Println(m.body)
		current := m.body
		for m.next != nil {
			// fmt.Printf("next= %s,", m.next.body)
			if current == "YOU" {
				you = append(you, m.next.body)
			} else if current == "SAN" {
				san = append(san, m.next.body)
			}
			sum = sum + 1
			// fmt.Printf("next= %s sum=%d,", m.next.body, sum)
			m = m.next

		}

		// fmt.Println()
	}
	fmt.Println(you)
	fmt.Println(san)
	// fmt.Println(mercury)
	fmt.Println(sum)
	transferNumber := 0
	for i, y := range you {
		for j, s := range san {
			if y == s {
				transferNumber = i + j
				break
			}
		}
		if transferNumber > 0 {
			break
		}
	}
	fmt.Println(transferNumber)
}
