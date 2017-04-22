package expects

import "testing"

func TestExpectBoolReturnsExpectation(t *testing.T) {
	value := true
	result := Bool(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectIntReturnsExpectation(t *testing.T) {
	value := 1
	result := Int(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectInt8ReturnsExpectation(t *testing.T) {
	value := int8(1)
	result := Int8(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectInt16ReturnsExpectation(t *testing.T) {
	value := int16(1)
	result := Int16(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectInt32ReturnsExpectation(t *testing.T) {
	value := int32(1)
	result := Int32(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectInt64ReturnsExpectation(t *testing.T) {
	value := int64(1)
	result := Int64(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectUintReturnsExpectation(t *testing.T) {
	value := uint(1)
	result := Uint(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}
func TestExpectUint8ReturnsExpectation(t *testing.T) {
	value := uint8(1)
	result := Uint8(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectUint16ReturnsExpectation(t *testing.T) {
	value := uint16(1)
	result := Uint16(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectUint32ReturnsExpectation(t *testing.T) {
	value := uint32(1)
	result := Uint32(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}
func TestExpectUint64ReturnsExpectation(t *testing.T) {
	value := uint64(1)
	result := Uint64(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}

func TestExpectStringReturnsExpectation(t *testing.T) {
	value := "test"
	result := String(t, value)

	if result.testContext != t || result.value != value {
		t.Fail()
	}
}
