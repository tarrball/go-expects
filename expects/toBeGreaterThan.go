package expects

func (source ExpectationInt) toBeGreaterThan(target int) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt8) toBeGreaterThan(target int8) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt16) toBeGreaterThan(target int16) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt32) toBeGreaterThan(target int32) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt64) toBeGreaterThan(target int64) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationString) toBeGreaterThan(target string) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%s' to be greater than '%s'",
			source.value,
			target)
	}
}

func (source ExpectationUint) toBeGreaterThan(target uint) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint8) toBeGreaterThan(target uint8) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint16) toBeGreaterThan(target uint16) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint32) toBeGreaterThan(target uint32) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint64) toBeGreaterThan(target uint64) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}
