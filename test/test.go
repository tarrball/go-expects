package test

import "reflect"

// SUT System Under Test contains value and testContext
type SUT struct {
	testContext
	actual interface{}
}

// Test takes a testContext and an object to be asserted on
func Test(t testContext) SUT {
	return SUT{testContext: t, actual: nil}
}

// TODO Create new SUT or modify the passed in one?
func (sut SUT) If(actual interface{}) SUT {
	sut.actual = actual
	return sut
}

// ConversionFail fails the test and outputs a message with the types
func ConversionFail(t testContext, actualType reflect.Type, expectedType reflect.Type) {
	t.Errorf("Could not convert actual type '%s' to expected type '%s'",
		actualType,
		expectedType)
}
