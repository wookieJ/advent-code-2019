package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const PATH = "/Day-01-Rocket-Equation/src/data/input"

func main() {
	fmt.Println("--- Day 1: The Tyranny of the Rocket Equation ---")
	pwd, _ := os.Getwd()
	input, err := getInput(pwd + PATH)
	if err != nil {
		log.Fatal(err)
	}
	masses := getIntModulesMassesFromInput(input)
	result := computeFuelSum(masses)
	fmt.Println(fmt.Sprintf(">> %d", result))
}

func getInput(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func getIntModulesMassesFromInput(input string) []int {
	massesSplit := strings.Split(input, "\n")
	var massesInt []int
	for _, mass := range massesSplit {
		tmp, err := strconv.Atoi(mass)
		if err != nil {
			log.Fatal(fmt.Sprintf("Unexpected input string %v", mass))
		}
		massesInt = append(massesInt, tmp)
	}
	return massesInt
}

func computeFuelMassFromModuleMass(mass int) int {
	result := 0
	fuelMass := (mass / 3) - 2
	if fuelMass > 0 {
		result += fuelMass
		result += computeFuelMassFromModuleMass(fuelMass)
	}
	return result
}

func computeFuelSum(masses []int) interface{} {
	result := 0
	for _, mass := range masses {
		result += computeFuelMassFromModuleMass(mass)
	}
	return result
}
