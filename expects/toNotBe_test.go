package expects

import (
	"testing"

	"github.com/tarrball/go-expects/mocks"
)

func TestToNotBeFloatMismatchPasses(t *testing.T) {
	actual, expected := 1.1, 2.2

	expectation := ExpectationFloat{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeFloatMatchFails(t *testing.T) {
	actual, expected := 1.1, 1.1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%f' should not equal '%f'", actual, expected)

	expectation := ExpectationFloat{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeFloat32MismatchPasses(t *testing.T) {
	actual, expected := float32(1.1), float32(2.2)

	expectation := ExpectationFloat32{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeFloat32MatchFails(t *testing.T) {
	actual, expected := float32(1.1), float32(1.1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%f' should not equal '%f'", actual, expected)

	expectation := ExpectationFloat32{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeIntMismatchPasses(t *testing.T) {
	actual, expected := 1, 2

	expectation := ExpectationInt{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeIntMatchFails(t *testing.T) {
	actual, expected := 1, 1
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := ExpectationInt{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeInt8MismatchPasses(t *testing.T) {
	actual, expected := int8(1), int8(2)

	expectation := ExpectationInt8{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeInt8MatchFails(t *testing.T) {
	actual, expected := int8(1), int8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := ExpectationInt8{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeInt16MismatchPasses(t *testing.T) {
	actual, expected := int16(1), int16(2)

	expectation := ExpectationInt16{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeInt16MatchFails(t *testing.T) {
	actual, expected := int16(1), int16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := ExpectationInt16{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeInt32MismatchPasses(t *testing.T) {
	actual, expected := int32(1), int32(2)

	expectation := ExpectationInt32{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeInt32MatchFails(t *testing.T) {
	actual, expected := int32(1), int32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := ExpectationInt32{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeInt64MismatchPasses(t *testing.T) {
	actual, expected := int64(1), int64(2)

	expectation := ExpectationInt64{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeInt64MatchFails(t *testing.T) {
	actual, expected := int64(1), int64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := ExpectationInt64{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeStringMismatchPasses(t *testing.T) {
	actual, expected := "abc", "123"

	expectation := ExpectationString{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeStringMatchFails(t *testing.T) {
	actual, expected := "test", "test"
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%s' should not equal '%s'", actual, expected)

	expectation := ExpectationString{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeUintMismatchPasses(t *testing.T) {
	actual, expected := uint(1), uint(2)

	expectation := ExpectationUint{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeUintMatchFails(t *testing.T) {
	actual, expected := uint(1), uint(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := ExpectationUint{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeUint8MismatchPasses(t *testing.T) {
	actual, expected := uint8(1), uint8(2)

	expectation := ExpectationUint8{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeUint8MatchFails(t *testing.T) {
	actual, expected := uint8(1), uint8(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := ExpectationUint8{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeUint16MismatchPasses(t *testing.T) {
	actual, expected := uint16(1), uint16(2)

	expectation := ExpectationUint16{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeUint16MatchFails(t *testing.T) {
	actual, expected := uint16(1), uint16(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := ExpectationUint16{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeUint32MismatchPasses(t *testing.T) {
	actual, expected := uint32(1), uint32(2)

	expectation := ExpectationUint32{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeUint32MatchFails(t *testing.T) {
	actual, expected := uint32(1), uint32(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := ExpectationUint32{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}

func TestToNotBeUint64MismatchPasses(t *testing.T) {
	actual, expected := uint64(1), uint64(2)

	expectation := ExpectationUint64{actual, Expectation{t}}

	expectation.ToNotBe(expected)
}

func TestToNotBeUint64MatchFails(t *testing.T) {
	actual, expected := uint64(1), uint64(1)
	mock := mocks.GetMock(t)
	mock.EXPECT().Errorf("'%d' should not equal '%d'", actual, expected)

	expectation := ExpectationUint64{actual, Expectation{mock}}

	expectation.ToNotBe(expected)
}
