package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const path = "/src/data/input"

type Range struct {
	min, max int
}

func main() {
	fmt.Println("--- Day 4: Secure Container ---")
	pwd, _ := os.Getwd()
	input, err := getInput(pwd + path)
	if err != nil {
		log.Fatal(err)
	}
	passwordRange, err := parseInputToPasswordRange(input)
	if err != nil {
		log.Fatal(err)
	}
	possiblePasswordsNumber := computePossilbePasswordsNumber(passwordRange, false)
	possiblePasswordsNumberWithCondition := computePossilbePasswordsNumber(passwordRange, true)
	fmt.Println(fmt.Sprintf("Part 1 >> %d", possiblePasswordsNumber))
	fmt.Println(fmt.Sprintf("Part 2 >> %d", possiblePasswordsNumberWithCondition))
}

func computePossilbePasswordsNumber(passwordRange Range, areAdjacentDigitsInGroup bool) int {
	result := 0
	for i := passwordRange.min; i <= passwordRange.max; i++ {
		if isNumberEligible(i, areAdjacentDigitsInGroup) {
			result++
		}
	}
	return result
}

func isNumberEligible(number int, areAdjacentDigitsInGroup bool) bool {
	return containsTwoAdjacentDigits(number, areAdjacentDigitsInGroup) && containsNeverDecreaseDigits(number)
}

func containsNeverDecreaseDigits(number int) bool {
	previousNumber := number % 10
	number /= 10
	for number != 0 {
		actualNumber := number % 10
		if actualNumber > previousNumber {
			return false
		}
		previousNumber = actualNumber
		number /= 10
	}
	return true
}

func containsTwoAdjacentDigits(number int, areAdjacentDigitsInGroup bool) bool {
	previousNumber := number % 10
	number /= 10
	numberOfAdjacentDigits := make(map[int]int)
	for number != 0 {
		actualNumber := number % 10
		if actualNumber == previousNumber {
			if !areAdjacentDigitsInGroup {
				return true
			} else {
				numberOfAdjacentDigits[actualNumber] = numberOfAdjacentDigits[actualNumber] + 1
			}
		}
		previousNumber = actualNumber
		number /= 10
	}
	for i := range numberOfAdjacentDigits {
		if numberOfAdjacentDigits[i] == 1 {
			return true
		}
	}
	return false
}

func parseInputToPasswordRange(input string) (Range, error) {
	passwordRangeArray := strings.Split(input, "-")
	if len(passwordRangeArray) != 2 {
		return Range{}, errors.New(fmt.Sprintf("Wrong input format. Found %v, required x-x", input))
	}
	min, err := strconv.Atoi(passwordRangeArray[0])
	if err != nil {
		return Range{}, errors.New(fmt.Sprintf("Cannot parse %v into number", passwordRangeArray[0]))
	}
	max, err := strconv.Atoi(passwordRangeArray[1])
	if err != nil {
		return Range{}, errors.New(fmt.Sprintf("Cannot parse %v into number", passwordRangeArray[1]))
	}
	return Range{min, max}, nil
}

func getInput(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
