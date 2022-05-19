package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMain_GetTime(t *testing.T) {
	
	testTable := []struct {
		name string
		ntpServer string
		err bool
	} {
		{
			name: "OK",
			ntpServer: "0.beevik-ntp.pool.ntp.org",
			err: false,
		},
		{
			name: "Bad url of ntp server",
			ntpServer: "BAD URL",
			err: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			time, err := GetTime(testCase.ntpServer)
			if testCase.err {
				assert.Nil(t, time)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, time)
			}
		})
	} 
}