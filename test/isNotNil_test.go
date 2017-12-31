package test

import "testing"
import "github.com/tarrball/gophertest/mocks"

func TestIsNotNilFailsIfNil(t *testing.T) {
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected value was nil.")

	sut := SUT{testContext: mock, actual: nil}

	sut.IsNotNil()
}

func TestIsNotNilPassesIfNotNil(t *testing.T) {
	sut := SUT{testContext: t, actual: "non nil value"}

	sut.IsNotNil()
}
