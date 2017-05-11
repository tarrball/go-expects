package expect

import (
	"testing"

	"github.com/tarrball/go-expect/mocks"
)

func TestToBeFalseFailsWithTrue(t *testing.T) {
	actual := true
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected 'true' to be 'false'")

	sut := This(mock, actual)

	sut.ToBeFalse()
}

func TestToBeFalsePassesWithFalse(t *testing.T) {
	actual := false

	sut := This(t, actual)

	sut.ToBeFalse()
}
