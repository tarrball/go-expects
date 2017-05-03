package expects

// SUT System Under Test contains value and testContext
type SUT struct {
	testContext
	value interface{}
}

// const tBool string = "bool"
// const tFloat32 string = "float32"
// const tFloat64 string = "float64"
// const tInt string = "int"
// const tInt8 string = "int8"
// const tInt16 string = "int16"
// const tInt32 string = "int32"
// const tInt64 string = "int64"
// const tString string = "string"
// const tUint string = "uint"
// const tUint8 string = "uint8"
// const tUint16 string = "uint16"
// const tUint32 string = "uint32"
// const tUint64 string = "uint64"

// This takes a testContext and an object to be asserted on
func This(t testContext, actual interface{}) SUT {
	return SUT{testContext: t, value: actual}
}

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
