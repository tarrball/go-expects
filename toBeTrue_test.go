package goexpectations

import (
	"testing"

	mocks "github.com/user/goexpectations/mocks"
)

func TestToBeTrueFailsWithFalse(t *testing.T) {
	actual := false
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected 'false' to be 'true'")

	expectation := expectationBool{actual, expectation{mock}}

	expectation.toBeTrue()
}

func TestToBeTruePassesWithTrue(t *testing.T) {
	actual := true

	expectation := expectationBool{actual, expectation{t}}

	expectation.toBeTrue()
}
