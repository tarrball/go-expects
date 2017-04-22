package expects

func (source ExpectationInt) toBe(target int) {
	if source.value != target {
		source.testContext.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt8) toBe(target int8) {
	if source.value != target {
		source.testContext.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt16) toBe(target int16) {
	if source.value != target {
		source.testContext.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt32) toBe(target int32) {
	if source.value != target {
		source.testContext.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt64) toBe(target int64) {
	if source.value != target {
		source.testContext.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationString) toBe(target string) {
	if source.value != target {
		source.testContext.Errorf("Expected '%s' to equal '%s'",
			source.value,
			target)
	}
}

func (source ExpectationUint) toBe(target uint) {
	if source.value != target {
		source.testContext.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint8) toBe(target uint8) {
	if source.value != target {
		source.testContext.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint16) toBe(target uint16) {
	if source.value != target {
		source.testContext.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint32) toBe(target uint32) {
	if source.value != target {
		source.testContext.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint64) toBe(target uint64) {
	if source.value != target {
		source.testContext.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}
