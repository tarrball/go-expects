package expect

import (
	"testing"

	"github.com/tarrball/gophertest/mocks"
)

func TestToBeLessThanFloatLessPasses(t *testing.T) {
	actual, expected := 1.1, 2.1

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanFloatLessFails(t *testing.T) {
	actual, expected := 2.1, float64(1.2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be less than '%f'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanFloatEqualFails(t *testing.T) {
	actual, expected := 1.1, 1.1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be less than '%f'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanFloat32LessPasses(t *testing.T) {
	actual, expected := float32(1.1), float32(2.1)

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanFloat32LessFails(t *testing.T) {
	actual, expected := float32(2.1), float32(1.2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be less than '%f'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanFloat32EqualFails(t *testing.T) {
	actual, expected := float32(1.1), float32(1.1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be less than '%f'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanIntLessPasses(t *testing.T) {
	actual, expected := 1, 2

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanIntLessFails(t *testing.T) {
	actual, expected := 2, 1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanIntEqualFails(t *testing.T) {
	actual, expected := 1, 1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt8LessPasses(t *testing.T) {
	actual, expected := int8(1), int8(2)

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt8LessFails(t *testing.T) {
	actual, expected := int8(2), int8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt8EqualFails(t *testing.T) {
	actual, expected := int8(1), int8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt16LessPasses(t *testing.T) {
	actual, expected := int16(1), int16(2)

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt16LessFails(t *testing.T) {
	actual, expected := int16(2), int16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt16EqualFails(t *testing.T) {
	actual, expected := int16(1), int16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt32LessPasses(t *testing.T) {
	actual, expected := int32(1), int32(2)

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt32LessFails(t *testing.T) {
	actual, expected := int32(2), int32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt32EqualFails(t *testing.T) {
	actual, expected := int32(1), int32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt64LessPasses(t *testing.T) {
	actual, expected := int64(1), int64(2)

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt64LessFails(t *testing.T) {
	actual, expected := int64(2), int64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanInt64EqualFails(t *testing.T) {
	actual, expected := int64(1), int64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanStringLessPasses(t *testing.T) {
	actual, expected := "abc", "def"

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanStringLessFails(t *testing.T) {
	actual, expected := "def", "abc"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to be less than '%s'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanStringEqualFails(t *testing.T) {
	actual, expected := "abc", "abc"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to be less than '%s'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUintLessPasses(t *testing.T) {
	actual, expected := uint(1), uint(2)

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUintLessFails(t *testing.T) {
	actual, expected := uint(2), uint(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUintEqualFails(t *testing.T) {
	actual, expected := uint(1), uint(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint8LessPasses(t *testing.T) {
	actual, expected := uint8(1), uint8(2)

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint8LessFails(t *testing.T) {
	actual, expected := uint8(2), uint8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint8EqualFails(t *testing.T) {
	actual, expected := uint8(1), uint8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint16LessPasses(t *testing.T) {
	actual, expected := uint16(1), uint16(2)

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint16LessFails(t *testing.T) {
	actual, expected := uint16(2), uint16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint16EqualFails(t *testing.T) {
	actual, expected := uint16(1), uint16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint32LessPasses(t *testing.T) {
	actual, expected := uint32(1), uint32(2)

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint32LessFails(t *testing.T) {
	actual, expected := uint32(2), uint32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint32EqualFails(t *testing.T) {
	actual, expected := uint32(1), uint32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint64LessPasses(t *testing.T) {
	actual, expected := uint64(1), uint64(2)

	sut := This(t, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint64LessFails(t *testing.T) {
	actual, expected := uint64(2), uint64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}

func TestToBeLessThanUint64EqualFails(t *testing.T) {
	actual, expected := uint64(1), uint64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'", actual, expected)

	sut := This(mock, actual)

	sut.ToBeLessThan(expected)
}
