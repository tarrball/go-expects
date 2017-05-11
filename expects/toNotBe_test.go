package expects

import (
	"testing"

	"github.com/tarrball/go-expects/mocks"
)

func TestToNotBeFloatMismatchPasses(t *testing.T) {
	actual, expected := 1.1, 2.2

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeFloatMatchFails(t *testing.T) {
	actual, expected := 1.1, 1.1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%f' should not be '%f'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeFloat32MismatchPasses(t *testing.T) {
	actual, expected := float32(1.1), float32(2.2)

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeFloat32MatchFails(t *testing.T) {
	actual, expected := float32(1.1), float32(1.1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%f' should not be '%f'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeIntMismatchPasses(t *testing.T) {
	actual, expected := 1, 2

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeIntMatchFails(t *testing.T) {
	actual, expected := 1, 1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeInt8MismatchPasses(t *testing.T) {
	actual, expected := int8(1), int8(2)

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeInt8MatchFails(t *testing.T) {
	actual, expected := int8(1), int8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeInt16MismatchPasses(t *testing.T) {
	actual, expected := int16(1), int16(2)

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeInt16MatchFails(t *testing.T) {
	actual, expected := int16(1), int16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeInt32MismatchPasses(t *testing.T) {
	actual, expected := int32(1), int32(2)

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeInt32MatchFails(t *testing.T) {
	actual, expected := int32(1), int32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeInt64MismatchPasses(t *testing.T) {
	actual, expected := int64(1), int64(2)

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeInt64MatchFails(t *testing.T) {
	actual, expected := int64(1), int64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeStringMismatchPasses(t *testing.T) {
	actual, expected := "abc", "123"

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeStringMatchFails(t *testing.T) {
	actual, expected := "test", "test"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%s' should not be '%s'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeUintMismatchPasses(t *testing.T) {
	actual, expected := uint(1), uint(2)

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeUintMatchFails(t *testing.T) {
	actual, expected := uint(1), uint(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeUint8MismatchPasses(t *testing.T) {
	actual, expected := uint8(1), uint8(2)

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeUint8MatchFails(t *testing.T) {
	actual, expected := uint8(1), uint8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeUint16MismatchPasses(t *testing.T) {
	actual, expected := uint16(1), uint16(2)

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeUint16MatchFails(t *testing.T) {
	actual, expected := uint16(1), uint16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeUint32MismatchPasses(t *testing.T) {
	actual, expected := uint32(1), uint32(2)

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeUint32MatchFails(t *testing.T) {
	actual, expected := uint32(1), uint32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeUint64MismatchPasses(t *testing.T) {
	actual, expected := uint64(1), uint64(2)

	sut := This(t, actual)

	sut.ToNotBe(expected)
}

func TestToNotBeUint64MatchFails(t *testing.T) {
	actual, expected := uint64(1), uint64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToNotBe(expected)
}
