package expect

import (
	"testing"

	"github.com/tarrball/gophertest/mocks"
)

func testThat(test func(*testing.T)) interface{} {
	return test
}

// Test(t).If(true).Is(false).AndFail().WithReason("i hate you")

func TestSomething(t *testing.T) {
	Test(t).If(10).Is(10)
}

func TestToBeFloatMatchPasses(t *testing.T) {
	actual, expected := 3.14, 3.14

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeFloatMismatchFails(t *testing.T) {
	actual, expected := 1.1, 2.2
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be '%f'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeFloat32MatchPasses(t *testing.T) {
	actual, expected := float32(3.14), float32(3.14)

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeFloat32MismatchFails(t *testing.T) {
	actual, expected := float32(1.1), float32(2.2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be '%f'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeIntMatchPasses(t *testing.T) {
	actual, expected := 1, 1

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeIntMismatchFails(t *testing.T) {
	actual, expected := 1, 2
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeInt8MatchPasses(t *testing.T) {
	actual, expected := int8(1), int8(1)

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeInt8MismatchFails(t *testing.T) {
	actual, expected := int8(1), int8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeInt16MatchPasses(t *testing.T) {
	actual, expected := int16(1), int16(1)

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeInt16MismatchFails(t *testing.T) {
	actual, expected := int16(1), int16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeInt32MatchPasses(t *testing.T) {
	actual, expected := int32(1), int32(1)

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeInt32MismatchFails(t *testing.T) {
	actual, expected := int32(1), int32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeInt64MatchPasses(t *testing.T) {
	actual, expected := int64(1), int64(1)

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeInt64MismatchFails(t *testing.T) {
	actual, expected := int64(1), int64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeStringMatchPasses(t *testing.T) {
	actual, expected := "test", "test"

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeStringMismatchFails(t *testing.T) {
	actual, expected := "abc", "123"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to be '%s'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeUintMatchPasses(t *testing.T) {
	actual, expected := uint(1), uint(1)

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeUintMismatchFails(t *testing.T) {
	actual, expected := uint(1), uint(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeUint8MatchPasses(t *testing.T) {
	actual, expected := uint8(1), uint8(1)

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeUint8MismatchFails(t *testing.T) {
	actual, expected := uint8(1), uint8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeUint16MatchPasses(t *testing.T) {
	actual, expected := uint16(1), uint16(1)

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeUint16MismatchFails(t *testing.T) {
	actual, expected := uint16(1), uint16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeUint32MatchPasses(t *testing.T) {
	actual, expected := uint32(1), uint32(1)

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeUint32MismatchFails(t *testing.T) {
	actual, expected := uint32(1), uint32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}

func TestToBeUint64MatchPasses(t *testing.T) {
	actual, expected := uint64(1), uint64(1)

	sut := This(t, actual)

	sut.ToBe(expected)
}

func TestToBeUint64MismatchFails(t *testing.T) {
	actual, expected := uint64(1), uint64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBe(expected)
}
