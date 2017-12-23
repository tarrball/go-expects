package test

import (
	"testing"

	"github.com/tarrball/gophertest/mocks"
)

func TestToBeFalseFailsWithTrue(t *testing.T) {
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected 'true' to be 'false'.")

	sut := SUT{testContext: mock, actual: true}
	sut.ToBeFalse()
}

func TestToBeFalsePassesWithFalse(t *testing.T) {
	sut := SUT{testContext: t, actual: false}
	sut.ToBeFalse()
}
