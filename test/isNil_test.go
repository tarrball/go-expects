package test

import "testing"
import "github.com/tarrball/gophertest/mocks"

func TestIsNilPassesIfNil(t *testing.T) {
	sut := SUT{testContext: t, actual: nil}

	sut.IsNil()
}

func TestIsNilFailsIfNotNil(t *testing.T) {
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected value was not nil.")

	sut := SUT{testContext: mock, actual: "non nil value"}

	sut.IsNil()
}
