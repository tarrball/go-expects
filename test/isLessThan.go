package test

import "reflect"

// IsLessThan expects the system under test value to be less than the expected value
func (sut SUT) IsLessThan(expected interface{}) {
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
		return
	}

	didPass := false

	switch sutType.Kind() {
	case reflect.Int:
		didPass = sut.actual.(int) < expectedValue.Interface().(int)
	case reflect.Int8:
		didPass = sut.actual.(int8) < expectedValue.Interface().(int8)
	case reflect.Int16:
		didPass = sut.actual.(int16) < expectedValue.Interface().(int16)
	case reflect.Int32:
		didPass = sut.actual.(int32) < expectedValue.Interface().(int32)
	case reflect.Int64:
		didPass = sut.actual.(int64) < expectedValue.Interface().(int64)
	case reflect.Uint:
		didPass = sut.actual.(uint) < expectedValue.Interface().(uint)
	case reflect.Uint8:
		didPass = sut.actual.(uint8) < expectedValue.Interface().(uint8)
	case reflect.Uint16:
		didPass = sut.actual.(uint16) < expectedValue.Interface().(uint16)
	case reflect.Uint32:
		didPass = sut.actual.(uint32) < expectedValue.Interface().(uint32)
	case reflect.Uint64:
		didPass = sut.actual.(uint64) < expectedValue.Interface().(uint64)
	// case reflect.Bool: // TODO this should fail
	// 	didPass = sut.actual.(bool) > expectedValue.Interface().(bool)
	case reflect.String:
		didPass = sut.actual.(string) < expectedValue.Interface().(string)
	case reflect.Float32:
		didPass = sut.actual.(float32) < expectedValue.Interface().(float32)
	case reflect.Float64:
		didPass = sut.actual.(float64) < expectedValue.Interface().(float64)
	}

	if !didPass {
		toBeLessThanFail(sut.testContext,
			sutType,
			sut.actual,
			expectedValue.Convert(sutType).Interface())
	}
}

func toBeLessThanFail(t testContext, objType reflect.Type, sut interface{}, expected interface{}) {
	switch objType.Kind() {
	case reflect.Int:
		t.Errorf("Expected '%d' to be less than '%d'.", sut, expected)
	case reflect.Int8:
		t.Errorf("Expected '%d' to be less than '%d'.", sut, expected)
	case reflect.Int16:
		t.Errorf("Expected '%d' to be less than '%d'.", sut, expected)
	case reflect.Int32:
		t.Errorf("Expected '%d' to be less than '%d'.", sut, expected)
	case reflect.Int64:
		t.Errorf("Expected '%d' to be less than '%d'.", sut, expected)
	case reflect.Uint:
		t.Errorf("Expected '%d' to be less than '%d'.", sut, expected)
	case reflect.Uint8:
		t.Errorf("Expected '%d' to be less than '%d'.", sut, expected)
	case reflect.Uint16:
		t.Errorf("Expected '%d' to be less than '%d'.", sut, expected)
	case reflect.Uint32:
		t.Errorf("Expected '%d' to be less than '%d'.", sut, expected)
	case reflect.Uint64:
		t.Errorf("Expected '%d' to be less than '%d'.", sut, expected)
	// case reflect.Bool: TODO should fail
	// 	t.Errorf("'%s' should was not less than '%s'.", sut, expected)
	case reflect.String:
		t.Errorf("Expected '%s' to be less than '%s'.", sut, expected)
	case reflect.Float32:
		t.Errorf("Expected '%f' to be less than '%f'.", sut, expected)
	case reflect.Float64:
		t.Errorf("Expected '%f' to be less than '%f'.", sut, expected)
	}
}
