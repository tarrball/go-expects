package expectations

func (source expectationInt) toBeGreaterThan(target int) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source expectationInt8) toBeGreaterThan(target int8) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source expectationInt16) toBeGreaterThan(target int16) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source expectationInt32) toBeGreaterThan(target int32) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source expectationInt64) toBeGreaterThan(target int64) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source expectationString) toBeGreaterThan(target string) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%s' to be greater than '%s'",
			source.value,
			target)
	}
}

func (source expectationUint) toBeGreaterThan(target uint) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source expectationUint8) toBeGreaterThan(target uint8) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source expectationUint16) toBeGreaterThan(target uint16) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source expectationUint32) toBeGreaterThan(target uint32) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}

func (source expectationUint64) toBeGreaterThan(target uint64) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}
