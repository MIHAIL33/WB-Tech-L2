package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_Unpacking(t *testing.T) {
	
	testTable := []struct {
		name string
		input string
		output string
		err bool
	} {
		{
			name: "OK",
			input: "a4bc2d5e",
			output: "aaaabccddddde",
			err: false,
		},
		{
			name: "OK",
			input: "abcd",
			output: "abcd",
			err: false,
		},
		{
			name: "Bad string",
			input: "45",
			output: "",
			err: true,
		},
		{
			name: "OK",
			input: "",
			output: "",
			err: false,
		},
		{
			name: "OK",
			input: "qwe\\4\\5",
			output: "qwe45",
			err: false,
		},
		{
			name: "OK",
			input: "qwe\\45",
			output: "qwe44444",
			err: false,
		},
		{
			name: "OK",
			input: "qwe\\\\5",
			output: "qwe\\\\\\\\\\",
			err: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := Unpacking(testCase.input)
			if testCase.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.output, *res)
			}
		})
	} 

}