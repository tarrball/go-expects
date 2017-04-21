package goexpectations

import (
	"testing"

	mocks "github.com/user/goexpectations/mocks"
)

func TestToBeFalseFailsWithTrue(t *testing.T) {
	actual := true
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected 'true' to be 'false'")

	expectation := expectationBool{actual, expectation{mock}}

	expectation.toBeFalse()
}

func TestToBeFalsePassesWithFalse(t *testing.T) {
	actual := false

	expectation := expectationBool{actual, expectation{t}}

	expectation.toBeFalse()
}
