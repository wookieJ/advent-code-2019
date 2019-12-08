package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

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

func TestShouldSplitInputIntoMasses(t *testing.T) {
	// given
	input := "1234\n9999\n1\n888"

	// when
	masses := getIntModulesMassesFromInput(input)

	// then
	assert.NotNil(t, masses)
	assert.Equal(t, 4, len(masses))
	assert.Equal(t, []int{1234, 9999, 1, 888}, masses)
}

func TestShouldComputeFuelMassFromModuleMass(t *testing.T) {
	// given
	mass1 := 14
	mass2 := 1969
	mass3 := 100756
	mass4 := 1

	// when
	fuelMass1 := computeFuelMassFromModuleMass(mass1)
	fuelMass2 := computeFuelMassFromModuleMass(mass2)
	fuelMass3 := computeFuelMassFromModuleMass(mass3)
	fuelMass4 := computeFuelMassFromModuleMass(mass4)

	// then
	assert.Equal(t, 2, fuelMass1)
	assert.Equal(t, 966, fuelMass2)
	assert.Equal(t, 50346, fuelMass3)
	assert.Equal(t, 0, fuelMass4)
}

func TestShouldComputeFuelSum(t *testing.T) {
	// given
	input := []int{30, 20, 1969}

	// when
	result := computeFuelSum(input)

	// then
	assert.Equal(t, 978, result)
}
