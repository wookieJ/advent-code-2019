package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestShouldGetCorrectIntersectionConfiguration(t *testing.T) {
	// given
	segment1 := Segment{Point{-5, 1}, Point{3, 1}}
	segment2 := Segment{Point{1, 12}, Point{1, -8}}

	segment3 := Segment{Point{0, 0}, Point{10, 0}}
	segment4 := Segment{Point{5, 5}, Point{5, -8}}

	segment5 := Segment{Point{7, 2}, Point{7, -6}}
	segment6 := Segment{Point{10, -1}, Point{5, -1}}

	// when
	intersectionConfiguration1 := getSegmentsIntersectionConfiguration(segment1, segment2)
	intersectionConfiguration2 := getSegmentsIntersectionConfiguration(segment3, segment4)
	intersectionConfiguration3 := getSegmentsIntersectionConfiguration(segment5, segment6)

	// then
	assert.NotNil(t, intersectionConfiguration1)
	assert.NotNil(t, intersectionConfiguration2)
	assert.NotNil(t, intersectionConfiguration3)
	assert.Equal(t, IntersectionConfiguration{1, 1, -5, 3, -8, 12}, intersectionConfiguration1)
	assert.Equal(t, IntersectionConfiguration{0, 5, 0, 10, -8, 5}, intersectionConfiguration2)
	assert.Equal(t, IntersectionConfiguration{-1, 7, 5, 10, -6, 2}, intersectionConfiguration3)
}

func TestShouldCheckIfSegmentsIntersectsBasedOnConfig(t *testing.T) {
	// given
	internConfig1 := IntersectionConfiguration{1, 1, -5, 3, -8, 12}
	internConfig2 := IntersectionConfiguration{3, 3, 1, 3, 1, 2}
	internConfig3 := IntersectionConfiguration{0, 5, 0, 10, -8, 5}

	// when
	result1 := areSegmentsIntersectsBasedOnConfig(internConfig1)
	result2 := areSegmentsIntersectsBasedOnConfig(internConfig2)
	result3 := areSegmentsIntersectsBasedOnConfig(internConfig3)

	// then
	assert.True(t, result1)
	assert.False(t, result2)
	assert.True(t, result3)
}

func TestShouldComputeManhattanDistance(t *testing.T) {
	// given
	point1 := Point{1, 5}
	point2 := Point{10, -5}
	point3 := Point{-1, -51}
	point4 := Point{0, 3}
	point5 := Point{7, 3}
	point6 := Point{0, -3}

	// when
	dist1 := manhattanDistanceFromZero(point1)
	dist2 := manhattanDistanceFromZero(point2)
	dist3 := manhattanDistanceFromZero(point3)
	dist4 := manhattanDistanceFromZero(point5)
	dist5 := manhattanDistance(point4, point5)
	dist6 := manhattanDistance(point4, point6)

	// then
	assert.Equal(t, 6, dist1)
	assert.Equal(t, 15, dist2)
	assert.Equal(t, 52, dist3)
	assert.Equal(t, 10, dist4)
	assert.Equal(t, 7, dist5)
	assert.Equal(t, 6, dist6)
}

func TestShouldFatalIfInputFileError(t *testing.T) {
	// when
	_, err := getInput("/not/existing/file")

	// then
	assert.Error(t, err)
}

func TestShouldGetInputFromFile(t *testing.T) {
	// given
	content := []byte("Test file")
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	// when
	input, err := getInput(tmpfile.Name())

	// then
	assert.Nil(t, err)
	assert.Equal(t, "Test file", input)
}

func TestShouldComputeWireCoordinates(t *testing.T) {
	// given
	instructionsOne := "R5,U5"
	instructionsTwo := "U7,R6,D4,L4"

	// when
	coordsOne := computeWireCoords(instructionsOne)
	coordsTwo := computeWireCoords(instructionsTwo)

	// then
	assert.NotNil(t, coordsOne)
	assert.NotNil(t, coordsTwo)
	assert.Equal(t, []Point{{0, 0}, {5, 0}, {5, 5}}, coordsOne)
	assert.Equal(t, []Point{{0, 0}, {0, 7}, {6, 7}, {6, 3}, {2, 3}}, coordsTwo)
}

func TestShouldComputeSegmentsIntersectionPoint(t *testing.T) {
	// given
	segment1 := Segment{Point{-5, 1}, Point{3, 1}}
	segment2 := Segment{Point{1, 12}, Point{1, -8}}

	segment3 := Segment{Point{0, 0}, Point{10, 0}}
	segment4 := Segment{Point{5, 5}, Point{5, -8}}

	segment5 := Segment{Point{1, 5}, Point{3, 5}}
	segment6 := Segment{Point{3, 1}, Point{3, 2}}

	segment7 := Segment{Point{7, 2}, Point{7, -6}}
	segment8 := Segment{Point{10, -1}, Point{5, -1}}

	// when
	result1 := computeSegmentsIntersectionPoint(segment1, segment2)
	result2 := computeSegmentsIntersectionPoint(segment3, segment4)
	result3 := computeSegmentsIntersectionPoint(segment5, segment6)
	result4 := computeSegmentsIntersectionPoint(segment7, segment8)

	// then
	assert.Equal(t, &Point{1, 1}, result1)
	assert.Equal(t, &Point{5, 0}, result2)
	assert.Nil(t, result3)
	assert.Equal(t, &Point{7, -1}, result4)
}

func TestShouldCheckIfIntersectionExistsBasedOnConfig(t *testing.T) {
	// given
	config1 := IntersectionConfiguration{-1, 7, 5, 10, -6, 2}

	// when
	result1 := areSegmentsIntersectsBasedOnConfig(config1)

	// then
	assert.True(t, result1)

}

func TestShouldCheckIfSegmentIsHorizontal(t *testing.T) {
	// given
	segment1 := Segment{Point{-5, 1}, Point{3, 1}}
	segment2 := Segment{Point{1, 12}, Point{1, -8}}
	segment3 := Segment{Point{0, 0}, Point{10, 0}}

	// when
	result1 := isHorizontal(segment1)
	result2 := isHorizontal(segment2)
	result3 := isHorizontal(segment3)

	// then
	assert.True(t, result1)
	assert.False(t, result2)
	assert.True(t, result3)
}

func TestShouldCheckIfSegmentArePerpendicular(t *testing.T) {
	// given
	segment1 := Segment{Point{-5, 1}, Point{3, 1}}
	segment2 := Segment{Point{1, 12}, Point{1, -8}}
	segment3 := Segment{Point{0, 0}, Point{10, 0}}
	segment4 := Segment{Point{5, 5}, Point{5, -8}}

	// when
	result1 := areSegmentsPerpendicular(segment1, segment2)
	result2 := areSegmentsPerpendicular(segment2, segment2)
	result3 := areSegmentsPerpendicular(segment2, segment3)
	result4 := areSegmentsPerpendicular(segment3, segment4)

	// then
	assert.True(t, result1)
	assert.False(t, result2)
	assert.True(t, result3)
	assert.True(t, result4)
}

func TestShouldComputeClosestIntersectionPointBasedOnInstruction(t *testing.T) {
	// given
	firstWireInstruction1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	secondWireInstruction1 := "U62,R66,U55,R34,D71,R55,D58,R83"

	firstWireInstruction2 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	secondWireInstruction2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

	firstWireInstruction3 := "R10,U10,L7"
	secondWireInstruction3 := "U20,R5,D15"

	firstWireInstruction4 := "U7,R6,D4,L4"
	secondWireInstruction4 := "R8,U5,L5,D3"

	// when
	result1 := computeClosestIntersectionDistanceAndPathLength(firstWireInstruction1, secondWireInstruction1)
	result2 := computeClosestIntersectionDistanceAndPathLength(firstWireInstruction2, secondWireInstruction2)
	result3 := computeClosestIntersectionDistanceAndPathLength(firstWireInstruction3, secondWireInstruction3)
	result4 := computeClosestIntersectionDistanceAndPathLength(firstWireInstruction4, secondWireInstruction4)

	// then
	assert.Equal(t, Pair{159, 610}, result1)
	assert.Equal(t, Pair{135, 410}, result2)
	assert.Equal(t, Pair{15, 60}, result3)
	assert.Equal(t, Pair{6, 30}, result4)
}
