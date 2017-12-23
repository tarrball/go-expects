package test

import (
	"testing"

	"github.com/tarrball/gophertest/mocks"
)

func TestIsGreaterThanFloatGreaterPasses(t *testing.T) {
	actual, expected := 2.1, 1.2

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanFloatLessFails(t *testing.T) {
	actual, expected := 1.2, 2.1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be greater than '%f'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanFloatEqualFails(t *testing.T) {
	actual, expected := 1.1, 1.1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be greater than '%f'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanFloat32GreaterPasses(t *testing.T) {
	actual, expected := float32(2.1), float32(1.2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanFloat32LessFails(t *testing.T) {
	actual, expected := float32(1.2), float32(2.1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be greater than '%f'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanFloat32EqualFails(t *testing.T) {
	actual, expected := float32(1.1), float32(1.1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be greater than '%f'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanIntGreaterPasses(t *testing.T) {
	actual, expected := 2, 1

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanIntLessFails(t *testing.T) {
	actual, expected := 1, 2
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanIntEqualFails(t *testing.T) {
	actual, expected := 1, 1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt8GreaterPasses(t *testing.T) {
	actual, expected := int8(2), int8(1)

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt8LessFails(t *testing.T) {
	actual, expected := int8(1), int8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt8EqualFails(t *testing.T) {
	actual, expected := int8(1), int8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt16GreaterPasses(t *testing.T) {
	actual, expected := int16(2), int16(1)

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt16LessFails(t *testing.T) {
	actual, expected := int16(1), int16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt16EqualFails(t *testing.T) {
	actual, expected := int16(1), int16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt32GreaterPasses(t *testing.T) {
	actual, expected := int32(2), int32(1)

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt32LessFails(t *testing.T) {
	actual, expected := int32(1), int32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt32EqualFails(t *testing.T) {
	actual, expected := int32(1), int32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt64GreaterPasses(t *testing.T) {
	actual, expected := int64(2), int64(1)

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt64LessFails(t *testing.T) {
	actual, expected := int64(1), int64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanInt64EqualFails(t *testing.T) {
	actual, expected := int64(1), int64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanStringGreaterPasses(t *testing.T) {
	actual, expected := "def", "abc"

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanStringLessFails(t *testing.T) {
	actual, expected := "abc", "def"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to be greater than '%s'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanStringEqualFails(t *testing.T) {
	actual, expected := "abc", "abc"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to be greater than '%s'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUintGreaterPasses(t *testing.T) {
	actual, expected := uint(2), uint(1)

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUintLessFails(t *testing.T) {
	actual, expected := uint(1), uint(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUintEqualFails(t *testing.T) {
	actual, expected := uint(1), uint(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint8GreaterPasses(t *testing.T) {
	actual, expected := uint8(2), uint8(1)

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint8LessFails(t *testing.T) {
	actual, expected := uint8(1), uint8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint8EqualFails(t *testing.T) {
	actual, expected := uint8(1), uint8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint16GreaterPasses(t *testing.T) {
	actual, expected := uint16(2), uint16(1)

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint16LessFails(t *testing.T) {
	actual, expected := uint16(1), uint16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint16EqualFails(t *testing.T) {
	actual, expected := uint16(1), uint16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint32GreaterPasses(t *testing.T) {
	actual, expected := uint32(2), uint32(1)

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint32LessFails(t *testing.T) {
	actual, expected := uint32(1), uint32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint32EqualFails(t *testing.T) {
	actual, expected := uint32(1), uint32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint64GreaterPasses(t *testing.T) {
	actual, expected := uint64(2), uint64(1)

	sut := SUT{testContext: t, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint64LessFails(t *testing.T) {
	actual, expected := uint64(1), uint64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}

func TestIsGreaterThanUint64EqualFails(t *testing.T) {
	actual, expected := uint64(1), uint64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be greater than '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsGreaterThan(expected)
}
