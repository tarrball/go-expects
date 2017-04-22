package expects

func (source ExpectationInt) toBeLessThan(target int) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt8) toBeLessThan(target int8) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt16) toBeLessThan(target int16) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt32) toBeLessThan(target int32) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt64) toBeLessThan(target int64) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationString) toBeLessThan(target string) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%s' to be less than '%s'",
			source.value,
			target)
	}
}

func (source ExpectationUint) toBeLessThan(target uint) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint8) toBeLessThan(target uint8) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint16) toBeLessThan(target uint16) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint32) toBeLessThan(target uint32) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint64) toBeLessThan(target uint64) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}
