package expect

import (
	"reflect"
)

// ToBeTrue expects the system under test value to be true.
func (actual SUT) ToBeTrue() {
	if actual.value == nil {
		actual.testContext.Errorf("Expected 'nil' to be true")
	}

	actualType := reflect.TypeOf(actual.value)

	if actualType.Kind() != reflect.Bool {
		actual.testContext.Errorf("Expected '%s' to be true", actualType.Kind())
	} else if !actual.value.(bool) {
		actual.testContext.Errorf("Expected 'false' to be 'true'")
	}
}
