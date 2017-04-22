package expectations

type expectation struct {
	testContext
}

type expectationBool struct {
	value bool
	expectation
}

type expectationInt struct {
	value int
	expectation
}

type expectationInt8 struct {
	value int8
	expectation
}

type expectationInt16 struct {
	value int16
	expectation
}

type expectationInt32 struct {
	value int32
	expectation
}

type expectationInt64 struct {
	value int64
	expectation
}

type expectationUint struct {
	value uint
	expectation
}

type expectationUint8 struct {
	value uint8
	expectation
}

type expectationUint16 struct {
	value uint16
	expectation
}


type expectationUint32 struct {
	value uint32
	expectation
}

type expectationUint64 struct {
	value uint64
	expectation
}

type expectationString struct {
	value string
	expectation
}
