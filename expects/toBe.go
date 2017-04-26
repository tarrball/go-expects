package expects

import (
	"reflect"
)

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
		actual.testContext.Errorf("Could not convert actual type '%d' to expected type '%d'",
			actualType,
			expectedType)
	} else if !reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual.value) {
		actual.testContext.Errorf("Expected '%d' to equal '%d'",
			actual.value,
			expectedValue.Convert(actualType).Interface())
	}
}

// func (source ExpectationFloat) ToBe(target float64) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%f' to equal '%f'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationFloat32) ToBe(target float32) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%f' to equal '%f'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationInt) ToBe(target int) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%d' to equal '%d'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationInt8) ToBe(target int8) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%d' to equal '%d'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationInt16) ToBe(target int16) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%d' to equal '%d'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationInt32) ToBe(target int32) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%d' to equal '%d'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationInt64) ToBe(target int64) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%d' to equal '%d'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationString) ToBe(target string) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%s' to equal '%s'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationUint) ToBe(target uint) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%d' to equal '%d'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationUint8) ToBe(target uint8) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%d' to equal '%d'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationUint16) ToBe(target uint16) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%d' to equal '%d'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationUint32) ToBe(target uint32) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%d' to equal '%d'",
// 			source.value,
// 			target)
// 	}
// }

// func (source ExpectationUint64) ToBe(target uint64) {
// 	if source.value != target {
// 		source.testContext.Errorf("Expected '%d' to equal '%d'",
// 			source.value,
// 			target)
// 	}
// }
