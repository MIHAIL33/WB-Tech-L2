package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_CheckURL(t *testing.T) {
	
	testTable := []struct {
		name string
		input string
		output bool
	} {
		{
			name: "OK",
			input: "/domain",
			output: true,
		},
		{
			name: "do not belong",
			input: "http://domain.com/domain",
			output: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res := CheckURL(testCase.input)
			assert.Equal(t, testCase.output, res)
		})
	}
	
}


func TestMain_ParseURL(t *testing.T) {
	
	testTable := []struct {
		name string
		inputURL string
		inputDomain bool
		output string
		isErr bool
	} {
		{
			name: "OK",
			inputURL: "http://domain.com/domain",
			inputDomain: true,
			output: "http://domain.com",
			isErr: false,
		},
		{
			name: "OK",
			inputURL: "http://domain.com/domain",
			inputDomain: false,
			output: "/domain",
			isErr: false,
		},
		{
			name: "Error",
			inputURL: "",
			inputDomain: false,
			output: "",
			isErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := ParseURL(testCase.inputURL, testCase.inputDomain)
			if testCase.isErr {
				assert.NotNil(t, err)
				assert.Equal(t, testCase.output, res)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.output, res)
			}
		})
	}
	
}