package expects

type Expectation struct {
	testContext
}

type ExpectationBool struct {
	value bool
	Expectation
}

type ExpectationFloat struct {
	value float64
	Expectation
}

type ExpectationFloat32 struct {
	value float32
	Expectation
}

type ExpectationInt struct {
	value int
	Expectation
}

type ExpectationInt8 struct {
	value int8
	Expectation
}

type ExpectationInt16 struct {
	value int16
	Expectation
}

type ExpectationInt32 struct {
	value int32
	Expectation
}

type ExpectationInt64 struct {
	value int64
	Expectation
}

type ExpectationUint struct {
	value uint
	Expectation
}

type ExpectationUint8 struct {
	value uint8
	Expectation
}

type ExpectationUint16 struct {
	value uint16
	Expectation
}

type ExpectationUint32 struct {
	value uint32
	Expectation
}

type ExpectationUint64 struct {
	value uint64
	Expectation
}

type ExpectationString struct {
	value string
	Expectation
}
