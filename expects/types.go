package expects

func Bool(t testContext, value bool) ExpectationBool {
	return ExpectationBool{value, Expectation{t}}
}

func Float(t testContext, value float64) ExpectationFloat {
	return ExpectationFloat{value, Expectation{t}}
}

func Float32(t testContext, value float32) ExpectationFloat32 {
	return ExpectationFloat32{value, Expectation{t}}
}

func Int(t testContext, value int) ExpectationInt {
	return ExpectationInt{value, Expectation{t}}
}

func Int8(t testContext, value int8) ExpectationInt8 {
	return ExpectationInt8{value, Expectation{t}}
}

func Int16(t testContext, value int16) ExpectationInt16 {
	return ExpectationInt16{value, Expectation{t}}
}

func Int32(t testContext, value int32) ExpectationInt32 {
	return ExpectationInt32{value, Expectation{t}}
}

func Int64(t testContext, value int64) ExpectationInt64 {
	return ExpectationInt64{value, Expectation{t}}
}

func String(t testContext, value string) ExpectationString {
	return ExpectationString{value, Expectation{t}}
}

func Uint(t testContext, value uint) ExpectationUint {
	return ExpectationUint{value, Expectation{t}}
}

func Uint8(t testContext, value uint8) ExpectationUint8 {
	return ExpectationUint8{value, Expectation{t}}
}

func Uint16(t testContext, value uint16) ExpectationUint16 {
	return ExpectationUint16{value, Expectation{t}}
}

func Uint32(t testContext, value uint32) ExpectationUint32 {
	return ExpectationUint32{value, Expectation{t}}
}

func Uint64(t testContext, value uint64) ExpectationUint64 {
	return ExpectationUint64{value, Expectation{t}}
}
