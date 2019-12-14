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
const noun = 12
const verb = 2
const Part2OutputValue = 19690720

type Pair struct {
	noun, verb int
}

func main() {
	fmt.Println("--- Day 2: 1202 Program Alarm ---")
	pwd, _ := os.Getwd()
	input, err := getInput(pwd + path)
	if err != nil {
		log.Fatal(err)
	}
	intCodePart1 := loadInputIntoTable(input)
	intCodePart1[1] = noun
	intCodePart1[2] = verb
	computingIntCode := computeIntCode(intCodePart1)
	fmt.Println(fmt.Sprintf("Part 1 >> %d", computingIntCode[0]))

	intCodePart2 := loadInputIntoTable(input)
	nounVerbPair := computeNounAndVerb(intCodePart2, Part2OutputValue)
	fmt.Println(fmt.Sprintf("Part 2 >> %d", 100*nounVerbPair.noun+nounVerbPair.verb))
}

func getInput(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func loadInputIntoTable(input string) []int {
	input = strings.TrimSpace(input)
	stringArray := strings.Split(input, ",")
	var intArray []int
	for _, it := range stringArray {
		intValue, err := strconv.Atoi(it)
		if err != nil {
			log.Fatal(err)
		}
		intArray = append(intArray, intValue)
	}
	return intArray
}

func computeIntCode(intCode []int) []int {
	for i := 0; i < len(intCode); i += 4 {
		code := intCode[i]
		if code == 1 {
			intCode = computeAdding(intCode, i)
		} else if code == 2 {
			intCode = computeMultiplying(intCode, i)
		} else if code == 99 {
			break
		}
	}
	return intCode
}

func computeAdding(intCode []int, i int) []int {
	if i+3 >= len(intCode) {
		log.Fatal()
	}
	intCode[intCode[i+3]] = intCode[intCode[i+1]] + intCode[intCode[i+2]]
	return intCode
}

func computeMultiplying(intCode []int, i int) []int {
	if i+3 >= len(intCode) {
		log.Fatal()
	}
	intCode[intCode[i+3]] = intCode[intCode[i+1]] * intCode[intCode[i+2]]
	return intCode
}

func computeNounAndVerb(intCode []int, outputValue int) Pair {
	for i := 0; i < len(intCode); i++ {
		for j := 0; j < len(intCode); j++ {
			testIntCode := make([]int, len(intCode))
			copy(testIntCode, intCode)
			testIntCode[1] = i
			testIntCode[2] = j
			if computeIntCode(testIntCode)[0] == outputValue {
				return Pair{i, j}
			}
		}
	}
	return Pair{-1, -1}
}
