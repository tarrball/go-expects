package expects

import (
	"reflect"
)

// ToBeFalse expects the system under test value to be false.
func (actual SUT) ToBeFalse() {
	if actual.value == nil {
		actual.testContext.Errorf("Expected 'nil' to be false")
	}

	actualType := reflect.TypeOf(actual.value)

	if actualType.Kind() != reflect.Bool {
		actual.testContext.Errorf("Expected '%s' to be false", actualType.Kind())
	} else if actual.value.(bool) {
		actual.testContext.Errorf("Expected 'true' to be 'false'")
	}
}
