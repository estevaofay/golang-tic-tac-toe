package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type coordinateTest struct {
	arg1, expectedx, expectedy int
}

var coordinateTests = []coordinateTest{
	{1, 0, 0},
	{2, 0, 1},
	{3, 0, 2},
	{4, 1, 0},
	{5, 1, 1},
	{6, 1, 2},
	{7, 2, 0},
	{8, 2, 1},
	{9, 2, 2},
}

func TestCoordinates(t *testing.T) {
	for _, test := range coordinateTests {
		outputX, outputY := GetCoordinates(test.arg1)
		assert.Equal(t, test.expectedx, outputX)
		assert.Equal(t, test.expectedy, outputY)
	}
}
