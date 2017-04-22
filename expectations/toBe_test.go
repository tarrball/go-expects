package expects

import (
	"testing"

	"github.com/tarrball/go-expectations/mocks"
)

func TestToBeIntMatchPasses(t *testing.T) {
	actual, expected := 1, 1

	expectation := ExpectationInt{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeIntMismatchFails(t *testing.T) {
	actual, expected := 1, 2
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := ExpectationInt{actual, Expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeInt8MatchPasses(t *testing.T) {
	actual, expected := int8(1), int8(1)

	expectation := ExpectationInt8{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeInt8MismatchFails(t *testing.T) {
	actual, expected := int8(1), int8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := ExpectationInt8{actual, Expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeInt16MatchPasses(t *testing.T) {
	actual, expected := int16(1), int16(1)

	expectation := ExpectationInt16{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeInt16MismatchFails(t *testing.T) {
	actual, expected := int16(1), int16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := ExpectationInt16{actual, Expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeInt32MatchPasses(t *testing.T) {
	actual, expected := int32(1), int32(1)

	expectation := ExpectationInt32{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeInt32MismatchFails(t *testing.T) {
	actual, expected := int32(1), int32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := ExpectationInt32{actual, Expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeInt64MatchPasses(t *testing.T) {
	actual, expected := int64(1), int64(1)

	expectation := ExpectationInt64{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeInt64MismatchFails(t *testing.T) {
	actual, expected := int64(1), int64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := ExpectationInt64{actual, Expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeStringMatchPasses(t *testing.T) {
	actual, expected := "test", "test"

	expectation := ExpectationString{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeStringMismatchFails(t *testing.T) {
	actual, expected := "abc", "123"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%s' to equal '%s'", actual, expected)

	expectation := ExpectationString{actual, Expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeUintMatchPasses(t *testing.T) {
	actual, expected := uint(1), uint(1)

	expectation := ExpectationUint{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeUintMismatchFails(t *testing.T) {
	actual, expected := uint(1), uint(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := ExpectationUint{actual, Expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeUint8MatchPasses(t *testing.T) {
	actual, expected := uint8(1), uint8(1)

	expectation := ExpectationUint8{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeUint8MismatchFails(t *testing.T) {
	actual, expected := uint8(1), uint8(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := ExpectationUint8{actual, Expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeUint16MatchPasses(t *testing.T) {
	actual, expected := uint16(1), uint16(1)

	expectation := ExpectationUint16{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeUint16MismatchFails(t *testing.T) {
	actual, expected := uint16(1), uint16(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := ExpectationUint16{actual, Expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeUint32MatchPasses(t *testing.T) {
	actual, expected := uint32(1), uint32(1)

	expectation := ExpectationUint32{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeUint32MismatchFails(t *testing.T) {
	actual, expected := uint32(1), uint32(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := ExpectationUint32{actual, Expectation{mock}}

	expectation.toBe(expected)
}

func TestToBeUint64MatchPasses(t *testing.T) {
	actual, expected := uint64(1), uint64(1)

	expectation := ExpectationUint64{actual, Expectation{t}}

	expectation.toBe(expected)
}

func TestToBeUint64MismatchFails(t *testing.T) {
	actual, expected := uint64(1), uint64(2)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("Expected '%d' to equal '%d'", actual, expected)

	expectation := ExpectationUint64{actual, Expectation{mock}}

	expectation.toBe(expected)
}
