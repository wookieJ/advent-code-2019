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

func TestShouldLoadStringInputIntoArray(t *testing.T) {
	// given
	input := "1,9,10,3,2,3,11,0,99,30,40,50"

	// when
	commands := loadInputIntoTable(input)

	// then
	assert.NotNil(t, commands)
	assert.Equal(t, 12, len(commands))
	assert.Equal(t, []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, commands)
}

func TestShouldComputeAdding(t *testing.T) {
	// given
	intCode := []int{1, 0, 2, 3, 99}

	// when
	resultArray := computeAdding(intCode, 0)

	// then
	assert.NotNil(t, resultArray)
	assert.Equal(t, len(intCode), len(resultArray))
	assert.Equal(t, []int{1, 0, 2, 3, 99}, resultArray)
}

func TestShouldComputeMultiplying(t *testing.T) {
	// given
	intCode := []int{2, 0, 2, 3, 99}

	// when
	resultArray := computeAdding(intCode, 0)

	// then
	assert.NotNil(t, resultArray)
	assert.Equal(t, len(intCode), len(resultArray))
	assert.Equal(t, []int{2, 0, 2, 4, 99}, resultArray)
}

func TestShouldComputeAlarmIntCode(t *testing.T) {
	// given
	intCode := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	// when
	resultArray := computeIntCode(intCode)

	// then
	assert.NotNil(t, resultArray)
	assert.Equal(t, len(intCode), len(resultArray))
	assert.Equal(t, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, resultArray)
}

func TestShouldComputeNounAndVerb(t *testing.T) {
	// given
	intCode := []int{1, 0, 0, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	// when
	nounAndVerb1 := computeNounAndVerb(intCode, 3500)

	// then
	assert.NotNil(t, nounAndVerb1)
	assert.Equal(t, Pair{9, 10}, nounAndVerb1)
}
