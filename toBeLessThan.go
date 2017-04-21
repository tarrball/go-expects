package goexpectations

func (source expectationInt) toBeLessThan(target int) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source expectationInt8) toBeLessThan(target int8) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source expectationInt16) toBeLessThan(target int16) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source expectationInt32) toBeLessThan(target int32) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source expectationInt64) toBeLessThan(target int64) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source expectationString) toBeLessThan(target string) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%s' to be less than '%s'",
			source.value,
			target)
	}
}

func (source expectationUint) toBeLessThan(target uint) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source expectationUint8) toBeLessThan(target uint8) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source expectationUint16) toBeLessThan(target uint16) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source expectationUint32) toBeLessThan(target uint32) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source expectationUint64) toBeLessThan(target uint64) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}
