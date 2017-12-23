package test

import (
	"reflect"
)

// IsTrue expects the system under test value to be true.
func (sut SUT) IsTrue() {
	if sut.actual == nil {
		sut.testContext.Errorf("Expected 'nil' to be true.")
	}

	sutType := reflect.TypeOf(sut.actual)

	if sutType.Kind() != reflect.Bool {
		sut.testContext.Errorf("Expected '%s' to be true.", sutType.Kind())
	} else if !sut.actual.(bool) {
		sut.testContext.Errorf("Expected 'false' to be 'true'.")
	}
}
