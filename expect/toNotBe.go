package expect

import "reflect"

// ToNotBe expects the system under test value to not equal the expected value
func (actual SUT) ToNotBe(expected interface{}) {
	if actual.value == nil && expected == nil {
		actual.testContext.Errorf("Both values were nil")
		return
	}

	if actual.value == nil {
		return
	}

	if expected == nil {
		return
	}

	actualType := reflect.TypeOf(actual.value)
	expectedType := reflect.TypeOf(expected)
	expectedValue := reflect.ValueOf(expected)

	if !expectedType.ConvertibleTo(actualType) || !expectedValue.Type().ConvertibleTo(actualType) {
		ConversionFail(actual.testContext, actualType, expectedType)
	} else if reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual.value) {
		toNotBeFail(actual.testContext,
			actualType,
			actual.value,
			expectedValue.Convert(actualType).Interface())
	}
}

func toNotBeFail(t testContext, objType reflect.Type, actual interface{}, expected interface{}) {
	switch objType.Kind() {
	case reflect.Int:
		t.Errorf("'%d' should not be '%d'", actual, expected)
	case reflect.Int8:
		t.Errorf("'%d' should not be '%d'", actual, expected)
	case reflect.Int16:
		t.Errorf("'%d' should not be '%d'", actual, expected)
	case reflect.Int32:
		t.Errorf("'%d' should not be '%d'", actual, expected)
	case reflect.Int64:
		t.Errorf("'%d' should not be '%d'", actual, expected)
	case reflect.Uint:
		t.Errorf("'%d' should not be '%d'", actual, expected)
	case reflect.Uint8:
		t.Errorf("'%d' should not be '%d'", actual, expected)
	case reflect.Uint16:
		t.Errorf("'%d' should not be '%d'", actual, expected)
	case reflect.Uint32:
		t.Errorf("'%d' should not be '%d'", actual, expected)
	case reflect.Uint64:
		t.Errorf("'%d' should not be '%d'", actual, expected)
	case reflect.Bool:
		t.Errorf("'%s' should not be '%s'", actual, expected)
	case reflect.String:
		t.Errorf("'%s' should not be '%s'", actual, expected)
	case reflect.Float32:
		t.Errorf("'%f' should not be '%f'", actual, expected)
	case reflect.Float64:
		t.Errorf("'%f' should not be '%f'", actual, expected)
	}
}
