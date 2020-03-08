package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Point type
type Point struct {
	x, y int
}

// Intersection type
type Intersection struct {
	p     Point
	wire1 int
	wire2 int
}

var centralPoint = Point{x: 1, y: 1}

func day3() {
	file, err := os.Open("input3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	in1s := strings.Split(scanner.Text(), ",")

	scanner.Scan()
	in2s := strings.Split(scanner.Text(), ",")

	// fmt.Println(in1s)
	// fmt.Println(in2s)
	smallestDistance := findNearestPoint(in1s, in2s)

	fmt.Println(smallestDistance)
	zz := findFewestStep(in1s, in2s)
	fmt.Println(zz)
}
func findNearestPoint(in1s, in2s []string) int {
	grid := make(map[Point]int)
	crossPoints := make([]Point, 0)
	intersetions := make([]Intersection, 0)

	gogo(1, in1s, grid, &crossPoints, &intersetions)
	// fmt.Println(crossPoints)
	gogo(2, in2s, grid, &crossPoints, &intersetions)
	// fmt.Println(crossPoints)
	var nearestPoint Point
	for i, currentPoint := range crossPoints {
		if i == 0 {
			nearestPoint = currentPoint
		}
		if distanceFromCentralPoint(currentPoint) < distanceFromCentralPoint(nearestPoint) {
			nearestPoint = currentPoint
		}
	}
	return distanceFromCentralPoint(nearestPoint)
}
func findFewestStep(in1s, in2s []string) int {
	grid := make(map[Point]int)
	crossPoints := make([]Point, 0)
	intersetions := make([]Intersection, 0)

	gogo(1, in1s, grid, &crossPoints, &intersetions)
	// fmt.Println(crossPoints)
	gogo(2, in2s, grid, &crossPoints, &intersetions)
	// fmt.Println(intersetions)
	var smallestStep int
	for i, intersetion := range intersetions {
		sum := intersetion.wire1 + intersetion.wire2
		if i == 0 {
			smallestStep = sum
		}
		if sum < smallestStep {
			smallestStep = sum
		}
	}
	return smallestStep
}
func distanceFromCentralPoint(p Point) int {
	return int(math.Abs(float64(p.x-centralPoint.x)) + math.Abs(float64(p.y-centralPoint.y)))
}
func gogo(wireth int, steps []string, grid map[Point]int, crossPoints *[]Point, intersetions *[]Intersection) {
	// centralPointX, centralPointY := 1, 1
	currentPoint := Point{1, 1}
	// currentX, currentY := centralPointX, centralPointY
	for i := 0; i < len(steps); i++ {
		command := string(steps[i][0])
		numstep, _ := strconv.Atoi(string(steps[i][1:]))
		// fmt.Println(numstep)

		wire := 0
		for k := 0; k < i; k++ {
			tmp, _ := strconv.Atoi(string(steps[k][1:]))
			wire += tmp
		}
		wire1, wire2 := 0, 0
		switch command {
		case "R":
			for j := 1; j <= numstep; j++ {
				currentPoint.y++
				if wirethStept, ok := grid[currentPoint]; ok {
					if wirethStept != 0 && wireth == 2 {
						*crossPoints = append(*crossPoints, currentPoint)
						wire1 = wirethStept
						wire2 = wire + j
						*intersetions = append(*intersetions, Intersection{currentPoint, wire1, wire2})
						// fmt.Println(steps[i])
						// fmt.Println(intersetions)
					}

					// fmt.Println(crossPoints)
				}
				if wireth == 1 && grid[currentPoint] == 0 {
					grid[currentPoint] = wire + j
				}
			}
			// fmt.Println(currentPoint)
			// fmt.Println(grid)

		case "L":
			for j := 1; j <= numstep; j++ {
				currentPoint.y--
				if wirethStept, ok := grid[currentPoint]; ok {
					if wirethStept != 0 && wireth == 2 {
						*crossPoints = append(*crossPoints, currentPoint)
						wire1 = wirethStept
						wire2 = wire + j
						*intersetions = append(*intersetions, Intersection{currentPoint, wire1, wire2})
					}
				}
				if wireth == 1 && grid[currentPoint] == 0 {
					grid[currentPoint] = wire + j
				}
			}
		case "U":
			for j := 1; j <= numstep; j++ {
				currentPoint.x++
				if wirethStept, ok := grid[currentPoint]; ok {
					if wirethStept != 0 && wireth == 2 {
						*crossPoints = append(*crossPoints, currentPoint)
						wire1 = wirethStept
						wire2 = wire + j
						*intersetions = append(*intersetions, Intersection{currentPoint, wire1, wire2})
					}
				}
				if wireth == 1 && grid[currentPoint] == 0 {
					grid[currentPoint] = wire + j
				}
			}
		case "D":
			for j := 1; j <= numstep; j++ {
				currentPoint.x--
				if wirethStept, ok := grid[currentPoint]; ok {
					if wirethStept != 0 && wireth == 2 {
						*crossPoints = append(*crossPoints, currentPoint)
						wire1 = wirethStept
						wire2 = wire + j
						*intersetions = append(*intersetions, Intersection{currentPoint, wire1, wire2})
					}
				}
				if wireth == 1 && grid[currentPoint] == 0 {
					grid[currentPoint] = wire + j
				}
			}
		}

	}
}
