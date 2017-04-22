package expects

func (source ExpectationInt) ToBeGreaterThan(target int) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt8) ToBeGreaterThan(target int8) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt16) ToBeGreaterThan(target int16) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt32) ToBeGreaterThan(target int32) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt64) ToBeGreaterThan(target int64) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationString) ToBeGreaterThan(target string) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%s' to be greater than '%s'",
			source.value,
			target)
	}
}

func (source ExpectationUint) ToBeGreaterThan(target uint) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint8) ToBeGreaterThan(target uint8) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint16) ToBeGreaterThan(target uint16) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint32) ToBeGreaterThan(target uint32) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint64) ToBeGreaterThan(target uint64) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}
