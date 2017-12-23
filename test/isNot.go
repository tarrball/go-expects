package test

import "reflect"

// IsNot expects the system under test value to not equal the expected value
func (sut SUT) IsNot(expected interface{}) {
	if sut.actual == nil {
		sut.testContext.Errorf("Actual value was nil.")
		return
	}

	if expected == nil {
		sut.testContext.Errorf("Expected value was nil.")
		return
	}

	sutType := reflect.TypeOf(sut.actual)
	expectedType := reflect.TypeOf(expected)
	expectedValue := reflect.ValueOf(expected)

	if !expectedType.ConvertibleTo(sutType) || !expectedValue.Type().ConvertibleTo(sutType) {
		ConversionFail(sut.testContext, sutType, expectedType)
	} else if reflect.DeepEqual(expectedValue.Convert(sutType).Interface(), sut.actual) {
		toNotBeFail(sut.testContext,
			sutType,
			sut.actual,
			expectedValue.Convert(sutType).Interface())
	}
}

func toNotBeFail(t testContext, objType reflect.Type, sut interface{}, expected interface{}) {
	switch objType.Kind() {
	case reflect.Int:
		t.Errorf("'%d' should not be '%d'.", sut, expected)
	case reflect.Int8:
		t.Errorf("'%d' should not be '%d'.", sut, expected)
	case reflect.Int16:
		t.Errorf("'%d' should not be '%d'.", sut, expected)
	case reflect.Int32:
		t.Errorf("'%d' should not be '%d'.", sut, expected)
	case reflect.Int64:
		t.Errorf("'%d' should not be '%d'.", sut, expected)
	case reflect.Uint:
		t.Errorf("'%d' should not be '%d'.", sut, expected)
	case reflect.Uint8:
		t.Errorf("'%d' should not be '%d'.", sut, expected)
	case reflect.Uint16:
		t.Errorf("'%d' should not be '%d'.", sut, expected)
	case reflect.Uint32:
		t.Errorf("'%d' should not be '%d'.", sut, expected)
	case reflect.Uint64:
		t.Errorf("'%d' should not be '%d'.", sut, expected)
	case reflect.Bool:
		t.Errorf("'%s' should not be '%s'.", sut, expected)
	case reflect.String:
		t.Errorf("'%s' should not be '%s'.", sut, expected)
	case reflect.Float32:
		t.Errorf("'%f' should not be '%f'.", sut, expected)
	case reflect.Float64:
		t.Errorf("'%f' should not be '%f'.", sut, expected)
	}
}
