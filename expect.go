package goexpectations

import "testing"

func expectBool(t *testing.T, value bool) expectationBool {
	return expectationBool{value, expectation{t}}
}

func expectInt(t *testing.T, value int) expectationInt {
	return expectationInt{value, expectation{t}}
}

func expectInt8(t *testing.T, value int8) expectationInt8 {
	return expectationInt8{value, expectation{t}}
}

func expectInt16(t *testing.T, value int16) expectationInt16 {
	return expectationInt16{value, expectation{t}}
}

func expectInt32(t *testing.T, value int32) expectationInt32 {
	return expectationInt32{value, expectation{t}}
}

func expectInt64(t *testing.T, value int64) expectationInt64 {
	return expectationInt64{value, expectation{t}}
}

func expectUint(t *testing.T, value uint) expectationUint {
	return expectationUint{value, expectation{t}}
}

func expectUint8(t *testing.T, value uint8) expectationUint8 {
	return expectationUint8{value, expectation{t}}
}

func expectUint16(t *testing.T, value uint16) expectationUint16 {
	return expectationUint16{value, expectation{t}}
}

func expectUint32(t *testing.T, value uint32) expectationUint32 {
	return expectationUint32{value, expectation{t}}
}

func expectUint64(t *testing.T, value uint64) expectationUint64 {
	return expectationUint64{value, expectation{t}}
}

func expectString(t *testing.T, value string) expectationString {
	return expectationString{value, expectation{t}}
}
