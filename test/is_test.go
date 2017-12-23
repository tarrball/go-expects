package test

import (
	"testing"

	"github.com/tarrball/gophertest/mocks"
)

func TestIsFloatMatchPasses(t *testing.T) {
	actual, expected := 3.14, 3.14

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsFloatMismatchFails(t *testing.T) {
	actual, expected := 1.1, 2.2
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be '%f'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsFloat32MatchPasses(t *testing.T) {
	actual, expected := float32(3.14), float32(3.14)

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsFloat32MismatchFails(t *testing.T) {
	actual, expected := float32(1.1), float32(2.2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%f' to be '%f'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsIntMatchPasses(t *testing.T) {
	actual, expected := 1, 1

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsIntMismatchFails(t *testing.T) {
	actual, expected := 1, 2
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsInt8MatchPasses(t *testing.T) {
	actual, expected := int8(1), int8(1)

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsInt8MismatchFails(t *testing.T) {
	actual, expected := int8(1), int8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsInt16MatchPasses(t *testing.T) {
	actual, expected := int16(1), int16(1)

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsInt16MismatchFails(t *testing.T) {
	actual, expected := int16(1), int16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsInt32MatchPasses(t *testing.T) {
	actual, expected := int32(1), int32(1)

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsInt32MismatchFails(t *testing.T) {
	actual, expected := int32(1), int32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsInt64MatchPasses(t *testing.T) {
	actual, expected := int64(1), int64(1)

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsInt64MismatchFails(t *testing.T) {
	actual, expected := int64(1), int64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsStringMatchPasses(t *testing.T) {
	actual, expected := "test", "test"

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsStringMismatchFails(t *testing.T) {
	actual, expected := "abc", "123"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to be '%s'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsUintMatchPasses(t *testing.T) {
	actual, expected := uint(1), uint(1)

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsUintMismatchFails(t *testing.T) {
	actual, expected := uint(1), uint(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsUint8MatchPasses(t *testing.T) {
	actual, expected := uint8(1), uint8(1)

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsUint8MismatchFails(t *testing.T) {
	actual, expected := uint8(1), uint8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsUint16MatchPasses(t *testing.T) {
	actual, expected := uint16(1), uint16(1)

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsUint16MismatchFails(t *testing.T) {
	actual, expected := uint16(1), uint16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsUint32MatchPasses(t *testing.T) {
	actual, expected := uint32(1), uint32(1)

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsUint32MismatchFails(t *testing.T) {
	actual, expected := uint32(1), uint32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}

func TestIsUint64MatchPasses(t *testing.T) {
	actual, expected := uint64(1), uint64(1)

	sut := Test(t)

	sut.If(actual).Is(expected)
}

func TestIsUint64MismatchFails(t *testing.T) {
	actual, expected := uint64(1), uint64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to be '%d'", actual, expected)

	sut := Test(mock)

	sut.If(actual).Is(expected)
}
