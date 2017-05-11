package expect

import (
	"reflect"
)

// ToBe expects the system under test value to equal the expected value.
func (actual SUT) ToBe(expected interface{}) {
	if actual.value == nil && expected == nil {
		return
	}

	if actual.value == nil {
		actual.testContext.Errorf("Actual value was nil")
		return
	}

	if expected == nil {
		actual.testContext.Errorf("Expected value was nil")
		return
	}

	actualType := reflect.TypeOf(actual.value)
	expectedType := reflect.TypeOf(expected)
	expectedValue := reflect.ValueOf(expected)

	if !expectedType.ConvertibleTo(actualType) || !expectedValue.Type().ConvertibleTo(actualType) {
		ConversionFail(actual.testContext, actualType, expectedType)
	} else if !reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual.value) {
		toBeFail(actual.testContext,
			actualType,
			actual.value,
			expectedValue.Convert(actualType).Interface())
	}
}

func toBeFail(t testContext, objType reflect.Type, actual interface{}, expected interface{}) {
	switch objType.Kind() {
	case reflect.Int:
		t.Errorf("Expected '%d' to be '%d'", actual, expected)
	case reflect.Int8:
		t.Errorf("Expected '%d' to be '%d'", actual, expected)
	case reflect.Int16:
		t.Errorf("Expected '%d' to be '%d'", actual, expected)
	case reflect.Int32:
		t.Errorf("Expected '%d' to be '%d'", actual, expected)
	case reflect.Int64:
		t.Errorf("Expected '%d' to be '%d'", actual, expected)
	case reflect.Uint:
		t.Errorf("Expected '%d' to be '%d'", actual, expected)
	case reflect.Uint8:
		t.Errorf("Expected '%d' to be '%d'", actual, expected)
	case reflect.Uint16:
		t.Errorf("Expected '%d' to be '%d'", actual, expected)
	case reflect.Uint32:
		t.Errorf("Expected '%d' to be '%d'", actual, expected)
	case reflect.Uint64:
		t.Errorf("Expected '%d' to be '%d'", actual, expected)
	case reflect.Bool:
		t.Errorf("Expected '%s' to be '%s'", actual, expected)
	case reflect.String:
		t.Errorf("Expected '%s' to be '%s'", actual, expected)
	case reflect.Float32:
		t.Errorf("Expected '%f' to be '%f'", actual, expected)
	case reflect.Float64:
		t.Errorf("Expected '%f' to be '%f'", actual, expected)
	}
}
