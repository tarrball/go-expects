package expect

import (
	"testing"

	"github.com/tarrball/gophertest/mocks"
)

func TestToBeTrueFailsWithFalse(t *testing.T) {
	actual := false
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected 'false' to be 'true'")

	sut := This(mock, actual)

	sut.ToBeTrue()
}

func TestToBeTruePassesWithTrue(t *testing.T) {
	actual := true

	sut := This(t, actual)

	sut.ToBeTrue()
}
