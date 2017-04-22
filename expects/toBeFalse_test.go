package expects

import (
	"testing"

	"github.com/tarrball/go-expects/mocks"
)

func TestToBeFalseFailsWithTrue(t *testing.T) {
	actual := true
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected 'true' to be 'false'")

	expectation := ExpectationBool{actual, Expectation{mock}}

	expectation.toBeFalse()
}

func TestToBeFalsePassesWithFalse(t *testing.T) {
	actual := false

	expectation := ExpectationBool{actual, Expectation{t}}

	expectation.toBeFalse()
}
