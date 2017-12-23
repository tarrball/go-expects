package test

import (
	"reflect"
)

// Is expects the system under test value to equal the expected value.
func (sut SUT) Is(expected interface{}) {
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
	} else if !reflect.DeepEqual(expectedValue.Convert(sutType).Interface(), sut.actual) {
		isFail(sut.testContext,
			sutType,
			sut.actual,
			expectedValue.Convert(sutType).Interface())
	}
}

func isFail(t testContext, objType reflect.Type, sut interface{}, expected interface{}) {
	switch objType.Kind() {
	case reflect.Int:
		t.Errorf("Expected '%d' to be '%d'.", sut, expected)
	case reflect.Int8:
		t.Errorf("Expected '%d' to be '%d'.", sut, expected)
	case reflect.Int16:
		t.Errorf("Expected '%d' to be '%d'.", sut, expected)
	case reflect.Int32:
		t.Errorf("Expected '%d' to be '%d'.", sut, expected)
	case reflect.Int64:
		t.Errorf("Expected '%d' to be '%d'.", sut, expected)
	case reflect.Uint:
		t.Errorf("Expected '%d' to be '%d'.", sut, expected)
	case reflect.Uint8:
		t.Errorf("Expected '%d' to be '%d'.", sut, expected)
	case reflect.Uint16:
		t.Errorf("Expected '%d' to be '%d'.", sut, expected)
	case reflect.Uint32:
		t.Errorf("Expected '%d' to be '%d'.", sut, expected)
	case reflect.Uint64:
		t.Errorf("Expected '%d' to be '%d'.", sut, expected)
	case reflect.Bool:
		t.Errorf("Expected '%s' to be '%s'.", sut, expected)
	case reflect.String:
		t.Errorf("Expected '%s' to be '%s'.", sut, expected)
	case reflect.Float32:
		t.Errorf("Expected '%f' to be '%f'.", sut, expected)
	case reflect.Float64:
		t.Errorf("Expected '%f' to be '%f'.", sut, expected)
	}
}
