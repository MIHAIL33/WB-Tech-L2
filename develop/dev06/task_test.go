package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_delimiter(t *testing.T) {

	strs := []string{"Winter: white: snow: frost",
					"Spring: green: grass: warm",
					"Summer: colorful: blossom: hot",
					"Autumn: yellow: leaves: cool",
					"asdsadasdasd"}

	res1 := []string{"Winter",
					"Spring",
					"Summer",
					"Autumn",}

	res2 := []string{"Winter",
					"Spring",
					"Summer",
					"Autumn",
					"asdsadasdasd"}
	
	testTable := []struct {
		name string
		input []string
		output []string
		field int
		delim string
		sep bool
	} {
		{
			name: "All options",
			input: strs,
			output: res1,
			field: 1,
			delim: ":",
			sep: true,
		},
		{
			name: "Default delimiter",
			input: strs,
			output: strs,
			field: 1,
			delim: "\t",
			sep: false,
		},
		{
			name: "No separated",
			input: strs,
			output: res2,
			field: 1,
			delim: ":",
			sep: false,
		}, 
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res, _ := delimiter(testCase.input, testCase.field, testCase.delim, testCase.sep)
			assert.Equal(t, testCase.output, res)
		})
	} 

}