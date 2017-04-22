package expects

import (
	"testing"

	"github.com/tarrball/go-expects/mocks"
)

func TestToBeTrueFailsWithFalse(t *testing.T) {
	actual := false
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected 'false' to be 'true'")

	expectation := ExpectationBool{actual, Expectation{mock}}

	expectation.toBeTrue()
}

func TestToBeTruePassesWithTrue(t *testing.T) {
	actual := true

	expectation := ExpectationBool{actual, Expectation{t}}

	expectation.toBeTrue()
}
