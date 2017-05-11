package expect

import "reflect"

// ToBeLessThan expects the system under test value to be less than the expected value
func (actual SUT) ToBeLessThan(expected interface{}) {
	if actual.value == nil && expected == nil {
		actual.testContext.Errorf("Both values were nil")
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
		return
	}

	didPass := false

	switch actualType.Kind() {
	case reflect.Int:
		didPass = actual.value.(int) < expectedValue.Interface().(int)
	case reflect.Int8:
		didPass = actual.value.(int8) < expectedValue.Interface().(int8)
	case reflect.Int16:
		didPass = actual.value.(int16) < expectedValue.Interface().(int16)
	case reflect.Int32:
		didPass = actual.value.(int32) < expectedValue.Interface().(int32)
	case reflect.Int64:
		didPass = actual.value.(int64) < expectedValue.Interface().(int64)
	case reflect.Uint:
		didPass = actual.value.(uint) < expectedValue.Interface().(uint)
	case reflect.Uint8:
		didPass = actual.value.(uint8) < expectedValue.Interface().(uint8)
	case reflect.Uint16:
		didPass = actual.value.(uint16) < expectedValue.Interface().(uint16)
	case reflect.Uint32:
		didPass = actual.value.(uint32) < expectedValue.Interface().(uint32)
	case reflect.Uint64:
		didPass = actual.value.(uint64) < expectedValue.Interface().(uint64)
	// case reflect.Bool: // TODO this should fail
	// 	didPass = actual.value.(bool) > expectedValue.Interface().(bool)
	case reflect.String:
		didPass = actual.value.(string) < expectedValue.Interface().(string)
	case reflect.Float32:
		didPass = actual.value.(float32) < expectedValue.Interface().(float32)
	case reflect.Float64:
		didPass = actual.value.(float64) < expectedValue.Interface().(float64)
	}

	if !didPass {
		toBeLessThanFail(actual.testContext,
			actualType,
			actual.value,
			expectedValue.Convert(actualType).Interface())
	}
}

func toBeLessThanFail(t testContext, objType reflect.Type, actual interface{}, expected interface{}) {
	switch objType.Kind() {
	case reflect.Int:
		t.Errorf("Expected '%d' to be less than '%d'", actual, expected)
	case reflect.Int8:
		t.Errorf("Expected '%d' to be less than '%d'", actual, expected)
	case reflect.Int16:
		t.Errorf("Expected '%d' to be less than '%d'", actual, expected)
	case reflect.Int32:
		t.Errorf("Expected '%d' to be less than '%d'", actual, expected)
	case reflect.Int64:
		t.Errorf("Expected '%d' to be less than '%d'", actual, expected)
	case reflect.Uint:
		t.Errorf("Expected '%d' to be less than '%d'", actual, expected)
	case reflect.Uint8:
		t.Errorf("Expected '%d' to be less than '%d'", actual, expected)
	case reflect.Uint16:
		t.Errorf("Expected '%d' to be less than '%d'", actual, expected)
	case reflect.Uint32:
		t.Errorf("Expected '%d' to be less than '%d'", actual, expected)
	case reflect.Uint64:
		t.Errorf("Expected '%d' to be less than '%d'", actual, expected)
	// case reflect.Bool: TODO should fail
	// 	t.Errorf("'%s' should was not less than '%s'", actual, expected)
	case reflect.String:
		t.Errorf("Expected '%s' to be less than '%s'", actual, expected)
	case reflect.Float32:
		t.Errorf("Expected '%f' to be less than '%f'", actual, expected)
	case reflect.Float64:
		t.Errorf("Expected '%f' to be less than '%f'", actual, expected)
	}
}
