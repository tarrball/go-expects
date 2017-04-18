package goexpectations

import "testing"

func TestToBeIntMatchPasses(t *testing.T) {
	expect := expectationInt{1, expectation{t}}
	expect.toBe(1)
}
