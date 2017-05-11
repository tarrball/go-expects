package expect

import "reflect"

// SUT System Under Test contains value and testContext
type SUT struct {
	testContext
	value interface{}
}

// This takes a testContext and an object to be asserted on
func This(t testContext, actual interface{}) SUT {
	return SUT{testContext: t, value: actual}
}

// ConversionFail fails the test and outputs a message with the types
func ConversionFail(t testContext, actualType reflect.Type, expectedType reflect.Type) {
	t.Errorf("Could not convert actual type '%s' to expected type '%s'",
		actualType,
		expectedType)
}
