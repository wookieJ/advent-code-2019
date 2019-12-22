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

func TestShouldParseInputIntoRange(t *testing.T) {
	// given
	wrongInput1 := "1:99"
	wrongInput2 := "1a-99"
	wrongInput3 := "1-99a"
	correctInput := "1-99"

	// when
	_, err1 := parseInputToPasswordRange(wrongInput1)
	_, err2 := parseInputToPasswordRange(wrongInput2)
	_, err3 := parseInputToPasswordRange(wrongInput3)
	passwordRange, err := parseInputToPasswordRange(correctInput)

	// then
	assert.NotNil(t, err1)
	assert.NotNil(t, err2)
	assert.NotNil(t, err3)
	assert.Nil(t, err)
	assert.Equal(t, Range{1, 99}, passwordRange)
}

	func TestShouldComputePossiblePasswordsNumber(t *testing.T) {
		// given
		range1 := Range{111111, 111121}
		range2 := Range{111111, 111131}
		range3 := Range{123444, 123460}

		// when
		passwordsNumber1 := computePossilbePasswordsNumber(range1, false)
		passwordsNumber2 := computePossilbePasswordsNumber(range2, false)
		passwordsNumber3 := computePossilbePasswordsNumber(range3, true)

		// then
		assert.Equal(t, 9, passwordsNumber1)
		assert.Equal(t, 17, passwordsNumber2)
		assert.Equal(t, 6, passwordsNumber3)
	}

func TestShouldCheckIfNumberIsEligible(t *testing.T) {
	// given
	number1 := 223450
	number2 := 123789
	number3 := 111111
	number4 := 122345

	// when
	result1 := isNumberEligible(number1, false)
	result2 := isNumberEligible(number2, false)
	result3 := isNumberEligible(number3, false)
	result4 := isNumberEligible(number4, false)

	// then
	assert.False(t, result1)
	assert.False(t, result2)
	assert.True(t, result3)
	assert.True(t, result4)
}

func TestShouldCheckIfNumberContainsTwoAdjacentDigits(t *testing.T) {
	// given
	number1 := 223450
	number2 := 123789
	number3 := 111111
	number4 := 122345
	number5 := 112233
	number6 := 123444
	number7 := 111122
	number8 := 122333

	// when
	result1 := containsTwoAdjacentDigits(number1, false)
	result2 := containsTwoAdjacentDigits(number2, false)
	result3 := containsTwoAdjacentDigits(number3, false)
	result4 := containsTwoAdjacentDigits(number4, false)
	result5 := containsTwoAdjacentDigits(number5, true)
	result6 := containsTwoAdjacentDigits(number6, true)
	result7 := containsTwoAdjacentDigits(number7, true)
	result8 := containsTwoAdjacentDigits(number8, true)

	// then
	assert.True(t, result1)
	assert.False(t, result2)
	assert.True(t, result3)
	assert.True(t, result4)
	assert.True(t, result5)
	assert.False(t, result6)
	assert.True(t, result7)
	assert.True(t, result8)
}

func TestShouldCheckIfNumberContainsNeverDecreaseDigits(t *testing.T) {
	// given
	number1 := 223450
	number2 := 123789
	number3 := 111111
	number4 := 122345

	// when
	result1 := containsNeverDecreaseDigits(number1)
	result2 := containsNeverDecreaseDigits(number2)
	result3 := containsNeverDecreaseDigits(number3)
	result4 := containsNeverDecreaseDigits(number4)

	// then
	assert.False(t, result1)
	assert.True(t, result2)
	assert.True(t, result3)
	assert.True(t, result4)
}
