package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_sortWord(t *testing.T) {
	
	testTable := []struct {
		name string
		input string
		output string
	} {
		{
			name: "OK",
			input: "столик",
			output: "иклост",
		},
		{
			name: "Empty",
			input: "",
			output: "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res := sortWord(testCase.input)
			assert.Equal(t, testCase.output, res)
		})
	}

}

func TestMain_sortStringSlice(t *testing.T) {
	
	testTable := []struct {
		name string
		input []string
		output []string
	} {
		{
			name: "OK",
			input: []string{"столик", "слоник", "дом"},
			output: []string{"дом", "слоник", "столик"},
		},
		{
			name: "Empty",
			input: []string{""},
			output: []string{""},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res := sortStringSlice(testCase.input)
			assert.Equal(t, testCase.output, res)
		})
	}
	
}

func TestMain_removeDuplicates(t *testing.T) {
	
	testTable := []struct {
		name string
		input []string
		output []string
	} {
		{
			name: "OK",
			input: []string{"столик", "столик", "слоник", "дом"},
			output: []string{"столик", "слоник", "дом"},
		},
		{
			name: "Empty",
			input: []string{""},
			output: []string{""},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res := removeDuplicates(testCase.input)
			assert.Equal(t, testCase.output, res)
		})
	}
	
}


func TestMain_SearchAnagrams(t *testing.T) {

	outputTrue := make(map[string][]string)
	outputTrue["листок"] = append(outputTrue["листок"], "листок")
	outputTrue["листок"] = append(outputTrue["листок"], "столик")
	
	testTable := []struct {
		name string
		input []string
		output map[string][]string
	} {
		{
			name: "OK",
			input: []string{"листок", "столик", "столик", "слоник", "дом"},
			output: outputTrue,
		},
		{
			name: "Empty",
			input: []string{""},
			output: make(map[string][]string),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res := SearchAnagrams(testCase.input)
			assert.Equal(t, testCase.output, res)
		})
	}
	
}

