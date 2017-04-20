package goexpectations

func expectBool(t testContext, value bool) expectationBool {
	return expectationBool{value, expectation{t}}
}

func expectInt(t testContext, value int) expectationInt {
	return expectationInt{value, expectation{t}}
}

func expectInt8(t testContext, value int8) expectationInt8 {
	return expectationInt8{value, expectation{t}}
}

func expectInt16(t testContext, value int16) expectationInt16 {
	return expectationInt16{value, expectation{t}}
}

func expectInt32(t testContext, value int32) expectationInt32 {
	return expectationInt32{value, expectation{t}}
}

func expectInt64(t testContext, value int64) expectationInt64 {
	return expectationInt64{value, expectation{t}}
}

func expectString(t testContext, value string) expectationString {
	return expectationString{value, expectation{t}}
}

func expectUint(t testContext, value uint) expectationUint {
	return expectationUint{value, expectation{t}}
}

func expectUint8(t testContext, value uint8) expectationUint8 {
	return expectationUint8{value, expectation{t}}
}

func expectUint16(t testContext, value uint16) expectationUint16 {
	return expectationUint16{value, expectation{t}}
}

func expectUint32(t testContext, value uint32) expectationUint32 {
	return expectationUint32{value, expectation{t}}
}

func expectUint64(t testContext, value uint64) expectationUint64 {
	return expectationUint64{value, expectation{t}}
}
