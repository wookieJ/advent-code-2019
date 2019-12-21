package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const path = "/src/data/input"
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type Point struct {
	x, y int
}

type Segment struct {
	p1, p2 Point
}

type Pair struct {
	a, b int
}

type IntersectionConfiguration struct {
	constHor, constVer, horizontalMin, horizontalMax, verticalMin, verticalMax int
}

func main() {
	fmt.Println("--- Day 3: Crossed Wires ---")
	pwd, _ := os.Getwd()
	input, err := getInput(pwd + path)
	if err != nil {
		log.Fatal(err)
	}
	wireInstructions := strings.Split(input, "\n")
	if len(wireInstructions) != 2 {
		log.Fatal(fmt.Sprintf("Incorrect input wire instructions length (expected 2, found %v",
			len(wireInstructions)))
	}
	closestDistance := computeClosestIntersectionDistanceAndPathLength(wireInstructions[0], wireInstructions[1])
	fmt.Println(fmt.Sprintf("Part 1 >> %d", closestDistance.a))
	fmt.Println(fmt.Sprintf("Part 2 >> %d", closestDistance.b))
}

func computeClosestIntersectionDistanceAndPathLength(firstWireInstructions string, secondWireInstructions string) Pair {
	firstWireCoords := computeWireCoords(firstWireInstructions)
	secondWireCoords := computeWireCoords(secondWireInstructions)
	minDistance := MaxInt
	minPathLength := MaxInt
	firstPath := 0
	for i := 0; i < len(firstWireCoords)-1; i++ {
		segment1 := Segment{firstWireCoords[i], firstWireCoords[i+1]}
		firstPath += manhattanDistance(segment1.p1, segment1.p2)
		secondPath := firstPath
		for j := 0; j < len(secondWireCoords)-1; j++ {
			segment2 := Segment{secondWireCoords[j], secondWireCoords[j+1]}
			secondPath += manhattanDistance(segment2.p1, segment2.p2)
			if intersectionPoint := computeSegmentsIntersectionPoint(segment1, segment2); intersectionPoint != nil {
				manhattanDistanceFromZero := manhattanDistanceFromZero(*intersectionPoint)
				secondPath -= manhattanDistance(segment1.p2, *intersectionPoint)
				secondPath -= manhattanDistance(segment2.p2, *intersectionPoint)
				if manhattanDistanceFromZero < minDistance {
					minDistance = manhattanDistanceFromZero
				}
				if secondPath < minPathLength {
					minPathLength = secondPath
				}
				break
			}
		}
	}
	return Pair{minDistance, minPathLength}
}

func getInput(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func computeWireCoords(instructions string) []Point {
	points := []Point{{0, 0}}
	instructionsTable := strings.Split(instructions, ",")
	for idx, instruction := range instructionsTable {
		value, err := strconv.Atoi(string(instruction[1:]))
		if err != nil {
			log.Fatal(err)
		}
		var oldPoint Point
		if idx > 0 {
			oldPoint = points[idx]
		}
		switch string(instruction[0]) {
		case "R":
			oldPoint.x += value
		case "L":
			oldPoint.x -= value
		case "U":
			oldPoint.y += value
		case "D":
			oldPoint.y -= value
		}
		points = append(points, oldPoint)
	}
	return points
}

func computeSegmentsIntersectionPoint(segment1, segment2 Segment) *Point {
	if !areSegmentsPerpendicular(segment1, segment2) {
		return nil
	}
	interConfig := getSegmentsIntersectionConfiguration(segment1, segment2)
	if areSegmentsIntersectsBasedOnConfig(interConfig) {
		return &Point{interConfig.constVer, interConfig.constHor}
	}
	return nil
}

func areSegmentsIntersectsBasedOnConfig(configuration IntersectionConfiguration) bool {
	return configuration.horizontalMin < configuration.constVer &&
		configuration.horizontalMax > configuration.constVer &&
		configuration.verticalMin < configuration.constHor &&
		configuration.verticalMax > configuration.constHor
}

func isHorizontal(segment Segment) bool {
	return segment.p1.y == segment.p2.y
}

func areSegmentsPerpendicular(segment1, segment2 Segment) bool {
	return isHorizontal(segment1) && !isHorizontal(segment2) || isHorizontal(segment2) && !isHorizontal(segment1)
}

func getSegmentsIntersectionConfiguration(segment1 Segment, segment2 Segment) IntersectionConfiguration {
	var constHor int
	var constVer int
	var horMin int
	var horMax int
	var verMin int
	var verMax int
	if isHorizontal(segment1) {
		constHor = segment1.p1.y
		constVer = segment2.p1.x
		if horMin = segment1.p1.x; segment1.p1.x <= segment1.p2.x {
			horMax = segment1.p2.x
		} else {
			horMin = segment1.p2.x
			horMax = segment1.p1.x
		}
		if verMin = segment2.p1.y; segment2.p1.y <= segment2.p2.y {
			verMax = segment2.p2.y
		} else {
			verMin = segment2.p2.y
			verMax = segment2.p1.y
		}
		return IntersectionConfiguration{constHor, constVer, horMin, horMax, verMin, verMax}
	}
	constHor = segment2.p1.y
	constVer = segment1.p1.x
	if horMin = segment2.p1.x; segment2.p1.x <= segment2.p2.x {
		horMax = segment2.p2.x
	} else {
		horMin = segment2.p2.x
		horMax = segment2.p1.x
	}
	if verMin = segment1.p1.y; segment1.p1.y <= segment1.p2.y {
		verMax = segment1.p2.y
	} else {
		verMin = segment1.p2.y
		verMax = segment1.p1.y
	}
	return IntersectionConfiguration{constHor, constVer, horMin, horMax, verMin, verMax}
}

func manhattanDistanceFromZero(point Point) int {
	x := point.x
	y := point.y
	if x < 0 {
		x = -point.x
	}
	if y < 0 {
		y = -point.y
	}
	return x + y
}

func manhattanDistance(point1, point2 Point) int {
	x := point1.x- point2.x
	y := point1.y - point2.y
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}
