package goexpectations

import (
	"testing"

	mocks "github.com/user/goexpectations/mocks"
)

func TestToNotBeIntMismatchPasses(t *testing.T) {
	actual, expected := 1, 2

	expectation := expectationInt{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeIntMatchFails(t *testing.T) {
	actual, expected := 1, 1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := expectationInt{actual, expectation{mock}}

	expectation.toNotBe(expected)
}

func TestToNotBeInt8MismatchPasses(t *testing.T) {
	actual, expected := int8(1), int8(2)

	expectation := expectationInt8{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeInt8MatchFails(t *testing.T) {
	actual, expected := int8(1), int8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := expectationInt8{actual, expectation{mock}}

	expectation.toNotBe(expected)
}

func TestToNotBeInt16MismatchPasses(t *testing.T) {
	actual, expected := int16(1), int16(2)

	expectation := expectationInt16{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeInt16MatchFails(t *testing.T) {
	actual, expected := int16(1), int16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := expectationInt16{actual, expectation{mock}}

	expectation.toNotBe(expected)
}

func TestToNotBeInt32MismatchPasses(t *testing.T) {
	actual, expected := int32(1), int32(2)

	expectation := expectationInt32{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeInt32MatchFails(t *testing.T) {
	actual, expected := int32(1), int32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := expectationInt32{actual, expectation{mock}}

	expectation.toNotBe(expected)
}

func TestToNotBeInt64MismatchPasses(t *testing.T) {
	actual, expected := int64(1), int64(2)

	expectation := expectationInt64{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeInt64MatchFails(t *testing.T) {
	actual, expected := int64(1), int64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := expectationInt64{expected, expectation{mock}}

	expectation.toNotBe(expected)
}

func TestToNotBeStringMismatchPasses(t *testing.T) {
	actual, expected := "abc", "123"

	expectation := expectationString{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeStringMatchFails(t *testing.T) {
	actual, expected := "test", "test"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%s' should not equal '%s'", actual, expected)

	expectation := expectationString{actual, expectation{mock}}

	expectation.toNotBe(expected)
}

func TestToNotBeUintMismatchPasses(t *testing.T) {
	actual, expected := uint(1), uint(2)

	expectation := expectationUint{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeUintMatchFails(t *testing.T) {
	actual, expected := uint(1), uint(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := expectationUint{actual, expectation{mock}}

	expectation.toNotBe(expected)
}

func TestToNotBeUint8MismatchPasses(t *testing.T) {
	actual, expected := uint8(1), uint8(2)

	expectation := expectationUint8{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeUint8MatchFails(t *testing.T) {
	actual, expected := uint8(1), uint8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := expectationUint8{actual, expectation{mock}}

	expectation.toNotBe(expected)
}

func TestToNotBeUint16MismatchPasses(t *testing.T) {
	actual, expected := uint16(1), uint16(2)

	expectation := expectationUint16{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeUint16MatchFails(t *testing.T) {
	actual, expected := uint16(1), uint16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := expectationUint16{actual, expectation{mock}}

	expectation.toNotBe(expected)
}

func TestToNotBeUint32MismatchPasses(t *testing.T) {
	actual, expected := uint32(1), uint32(2)

	expectation := expectationUint32{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeUint32MatchFails(t *testing.T) {
	actual, expected := uint32(1), uint32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := expectationUint32{expected, expectation{mock}}

	expectation.toNotBe(expected)
}

func TestToNotBeUint64MismatchPasses(t *testing.T) {
	actual, expected := uint64(1), uint64(2)

	expectation := expectationUint64{actual, expectation{t}}

	expectation.toNotBe(expected)
}

func TestToNotBeUint64MatchFails(t *testing.T) {
	actual, expected := uint64(1), uint64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := expectationUint64{actual, expectation{mock}}

	expectation.toNotBe(expected)
}
