package expects

func (source ExpectationFloat) ToBeLessThan(target float64) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%f' to be less than '%f'",
			source.value,
			target)
	}
}

func (source ExpectationFloat32) ToBeLessThan(target float32) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%f' to be less than '%f'",
			source.value,
			target)
	}
}

func (source ExpectationInt) ToBeLessThan(target int) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt8) ToBeLessThan(target int8) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt16) ToBeLessThan(target int16) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt32) ToBeLessThan(target int32) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt64) ToBeLessThan(target int64) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationString) ToBeLessThan(target string) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%s' to be less than '%s'",
			source.value,
			target)
	}
}

func (source ExpectationUint) ToBeLessThan(target uint) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint8) ToBeLessThan(target uint8) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint16) ToBeLessThan(target uint16) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint32) ToBeLessThan(target uint32) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint64) ToBeLessThan(target uint64) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}
