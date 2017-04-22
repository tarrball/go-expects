package expectations

import "testing"

func TestExpectBoolReturnsExpectation(t *testing.T) {
	value := true
	result := expectBool(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectIntReturnsExpectation(t *testing.T) {
	value := 1
	result := expectInt(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectInt8ReturnsExpectation(t *testing.T) {
	value := int8(1)
	result := expectInt8(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectInt16ReturnsExpectation(t *testing.T) {
	value := int16(1)
	result := expectInt16(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectInt32ReturnsExpectation(t *testing.T) {
	value := int32(1)
	result := expectInt32(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectInt64ReturnsExpectation(t *testing.T) {
	value := int64(1)
	result := expectInt64(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectUintReturnsExpectation(t *testing.T) {
	value := uint(1)
	result := expectUint(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}
func TestExpectUint8ReturnsExpectation(t *testing.T) {
	value := uint8(1)
	result := expectUint8(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectUint16ReturnsExpectation(t *testing.T) {
	value := uint16(1)
	result := expectUint16(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectUint32ReturnsExpectation(t *testing.T) {
	value := uint32(1)
	result := expectUint32(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}
func TestExpectUint64ReturnsExpectation(t *testing.T) {
	value := uint64(1)
	result := expectUint64(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectStringReturnsExpectation(t *testing.T) {
	value := "test"
	result := expectString(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}
