package expectations

import (
	"testing"

	"github.com/tarrball/goexpectations/mocks"
)

func TestToBeIntMatchPasses(t *testing.T) {
	actual, expected := 1, 1

	expectation := expectationInt{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeIntMismatchFails(t *testing.T) {
	actual, expected := 1, 2
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := expectationInt{actual, expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeInt8MatchPasses(t *testing.T) {
	actual, expected := int8(1), int8(1)

	expectation := expectationInt8{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeInt8MismatchFails(t *testing.T) {
	actual, expected := int8(1), int8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := expectationInt8{actual, expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeInt16MatchPasses(t *testing.T) {
	actual, expected := int16(1), int16(1)

	expectation := expectationInt16{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeInt16MismatchFails(t *testing.T) {
	actual, expected := int16(1), int16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := expectationInt16{actual, expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeInt32MatchPasses(t *testing.T) {
	actual, expected := int32(1), int32(1)

	expectation := expectationInt32{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeInt32MismatchFails(t *testing.T) {
	actual, expected := int32(1), int32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := expectationInt32{actual, expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeInt64MatchPasses(t *testing.T) {
	actual, expected := int64(1), int64(1)

	expectation := expectationInt64{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeInt64MismatchFails(t *testing.T) {
	actual, expected := int64(1), int64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := expectationInt64{actual, expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeStringMatchPasses(t *testing.T) {
	actual, expected := "test", "test"

	expectation := expectationString{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeStringMismatchFails(t *testing.T) {
	actual, expected := "abc", "123"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to equal '%s'", actual, expected)

	expectation := expectationString{actual, expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeUintMatchPasses(t *testing.T) {
	actual, expected := uint(1), uint(1)

	expectation := expectationUint{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeUintMismatchFails(t *testing.T) {
	actual, expected := uint(1), uint(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := expectationUint{actual, expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeUint8MatchPasses(t *testing.T) {
	actual, expected := uint8(1), uint8(1)

	expectation := expectationUint8{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeUint8MismatchFails(t *testing.T) {
	actual, expected := uint8(1), uint8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := expectationUint8{actual, expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeUint16MatchPasses(t *testing.T) {
	actual, expected := uint16(1), uint16(1)

	expectation := expectationUint16{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeUint16MismatchFails(t *testing.T) {
	actual, expected := uint16(1), uint16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := expectationUint16{actual, expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeUint32MatchPasses(t *testing.T) {
	actual, expected := uint32(1), uint32(1)

	expectation := expectationUint32{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeUint32MismatchFails(t *testing.T) {
	actual, expected := uint32(1), uint32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := expectationUint32{actual, expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeUint64MatchPasses(t *testing.T) {
	actual, expected := uint64(1), uint64(1)

	expectation := expectationUint64{actual, expectation{t}}

	expectation.toBe(expected)
}

func TestToBeUint64MismatchFails(t *testing.T) {
	actual, expected := uint64(1), uint64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := expectationUint64{actual, expectation{mock}}

	expectation.toBe(expected)
}
