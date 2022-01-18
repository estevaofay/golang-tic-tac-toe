package main

import "testing"

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

func TestAdd(t *testing.T) {

	for _, test := range coordinateTests {
		if outputX, outputY := GetCoordinates(test.arg1); outputX != test.expectedx && outputY != test.expectedy {
			t.Errorf("Output %q not equal to expected %q and/or output %q not equal to expected output %q", outputX, test.expectedx, outputY, test.expectedy)
		}
	}
}
