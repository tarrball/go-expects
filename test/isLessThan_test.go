package test

import (
	"testing"

	"github.com/tarrball/gophertest/mocks"
)

func TestIsLessThanFloatLessPasses(t *testing.T) {
	actual, expected := 1.1, 2.1

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanFloatLessFails(t *testing.T) {
	actual, expected := 2.1, float64(1.2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be less than '%f'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanFloatEqualFails(t *testing.T) {
	actual, expected := 1.1, 1.1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be less than '%f'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanFloat32LessPasses(t *testing.T) {
	actual, expected := float32(1.1), float32(2.1)

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanFloat32LessFails(t *testing.T) {
	actual, expected := float32(2.1), float32(1.2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be less than '%f'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanFloat32EqualFails(t *testing.T) {
	actual, expected := float32(1.1), float32(1.1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be less than '%f'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanIntLessPasses(t *testing.T) {
	actual, expected := 1, 2

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanIntLessFails(t *testing.T) {
	actual, expected := 2, 1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanIntEqualFails(t *testing.T) {
	actual, expected := 1, 1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt8LessPasses(t *testing.T) {
	actual, expected := int8(1), int8(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt8LessFails(t *testing.T) {
	actual, expected := int8(2), int8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt8EqualFails(t *testing.T) {
	actual, expected := int8(1), int8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt16LessPasses(t *testing.T) {
	actual, expected := int16(1), int16(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt16LessFails(t *testing.T) {
	actual, expected := int16(2), int16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt16EqualFails(t *testing.T) {
	actual, expected := int16(1), int16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt32LessPasses(t *testing.T) {
	actual, expected := int32(1), int32(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt32LessFails(t *testing.T) {
	actual, expected := int32(2), int32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt32EqualFails(t *testing.T) {
	actual, expected := int32(1), int32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt64LessPasses(t *testing.T) {
	actual, expected := int64(1), int64(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt64LessFails(t *testing.T) {
	actual, expected := int64(2), int64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanInt64EqualFails(t *testing.T) {
	actual, expected := int64(1), int64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanStringLessPasses(t *testing.T) {
	actual, expected := "abc", "def"

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanStringLessFails(t *testing.T) {
	actual, expected := "def", "abc"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to be less than '%s'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanStringEqualFails(t *testing.T) {
	actual, expected := "abc", "abc"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to be less than '%s'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUintLessPasses(t *testing.T) {
	actual, expected := uint(1), uint(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUintLessFails(t *testing.T) {
	actual, expected := uint(2), uint(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUintEqualFails(t *testing.T) {
	actual, expected := uint(1), uint(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint8LessPasses(t *testing.T) {
	actual, expected := uint8(1), uint8(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint8LessFails(t *testing.T) {
	actual, expected := uint8(2), uint8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint8EqualFails(t *testing.T) {
	actual, expected := uint8(1), uint8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint16LessPasses(t *testing.T) {
	actual, expected := uint16(1), uint16(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint16LessFails(t *testing.T) {
	actual, expected := uint16(2), uint16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint16EqualFails(t *testing.T) {
	actual, expected := uint16(1), uint16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint32LessPasses(t *testing.T) {
	actual, expected := uint32(1), uint32(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint32LessFails(t *testing.T) {
	actual, expected := uint32(2), uint32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint32EqualFails(t *testing.T) {
	actual, expected := uint32(1), uint32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint64LessPasses(t *testing.T) {
	actual, expected := uint64(1), uint64(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint64LessFails(t *testing.T) {
	actual, expected := uint64(2), uint64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}

func TestIsLessThanUint64EqualFails(t *testing.T) {
	actual, expected := uint64(1), uint64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be less than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsLessThan(expected)
}
