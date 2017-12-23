package test

import (
	"reflect"
)

// ToBeFalse expects the system under test value to be false.
func (sut SUT) ToBeFalse() {
	if sut.actual == nil {
		sut.testContext.Errorf("Expected 'nil' to be false.")
	}

	sutType := reflect.TypeOf(sut.actual)

	if sutType.Kind() != reflect.Bool {
		sut.testContext.Errorf("Expected '%s' to be false.", sutType.Kind())
	} else if sut.actual.(bool) {
		sut.testContext.Errorf("Expected 'true' to be 'false'.")
	}
}
