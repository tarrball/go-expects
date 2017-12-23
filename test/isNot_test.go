package test

import (
	"testing"

	"github.com/tarrball/gophertest/mocks"
)

func TestIsNotFloatMismatchPasses(t *testing.T) {
	actual, expected := 1.1, 2.2

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotFloatMatchFails(t *testing.T) {
	actual, expected := 1.1, 1.1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%f' should not be '%f'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotFloat32MismatchPasses(t *testing.T) {
	actual, expected := float32(1.1), float32(2.2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotFloat32MatchFails(t *testing.T) {
	actual, expected := float32(1.1), float32(1.1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%f' should not be '%f'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotIntMismatchPasses(t *testing.T) {
	actual, expected := 1, 2

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotIntMatchFails(t *testing.T) {
	actual, expected := 1, 1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotInt8MismatchPasses(t *testing.T) {
	actual, expected := int8(1), int8(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotInt8MatchFails(t *testing.T) {
	actual, expected := int8(1), int8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotInt16MismatchPasses(t *testing.T) {
	actual, expected := int16(1), int16(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotInt16MatchFails(t *testing.T) {
	actual, expected := int16(1), int16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotInt32MismatchPasses(t *testing.T) {
	actual, expected := int32(1), int32(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotInt32MatchFails(t *testing.T) {
	actual, expected := int32(1), int32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotInt64MismatchPasses(t *testing.T) {
	actual, expected := int64(1), int64(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotInt64MatchFails(t *testing.T) {
	actual, expected := int64(1), int64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotStringMismatchPasses(t *testing.T) {
	actual, expected := "abc", "123"

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotStringMatchFails(t *testing.T) {
	actual, expected := "test", "test"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%s' should not be '%s'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotUintMismatchPasses(t *testing.T) {
	actual, expected := uint(1), uint(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotUintMatchFails(t *testing.T) {
	actual, expected := uint(1), uint(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotUint8MismatchPasses(t *testing.T) {
	actual, expected := uint8(1), uint8(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotUint8MatchFails(t *testing.T) {
	actual, expected := uint8(1), uint8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotUint16MismatchPasses(t *testing.T) {
	actual, expected := uint16(1), uint16(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotUint16MatchFails(t *testing.T) {
	actual, expected := uint16(1), uint16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotUint32MismatchPasses(t *testing.T) {
	actual, expected := uint32(1), uint32(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotUint32MatchFails(t *testing.T) {
	actual, expected := uint32(1), uint32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotUint64MismatchPasses(t *testing.T) {
	actual, expected := uint64(1), uint64(2)

	sut := SUT{testContext: t, actual: actual}

	sut.IsNot(expected)
}

func TestIsNotUint64MatchFails(t *testing.T) {
	actual, expected := uint64(1), uint64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not be '%d'.", actual, expected)

	sut := SUT{testContext: mock, actual: actual}

	sut.IsNot(expected)
}
