package test

import (
	"testing"

	"github.com/tarrball/gophertest/mocks"
)

func TestIsTrueFailsWithFalse(t *testing.T) {
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected 'false' to be 'true'.")

	sut := SUT{testContext: mock, actual: false}
	sut.IsTrue()
}

func TestIsTruePassesWithTrue(t *testing.T) {
	sut := SUT{testContext: t, actual: true}
	sut.IsTrue()
}
