package expects

func (source ExpectationInt) ToNotBe(target int) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt8) ToNotBe(target int8) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt16) ToNotBe(target int16) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt32) ToNotBe(target int32) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt64) ToNotBe(target int64) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationString) ToNotBe(target string) {
	if source.value == target {
		source.testContext.Errorf("'%s' should not equal '%s'",
			source.value,
			target)
	}
}

func (source ExpectationUint) ToNotBe(target uint) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint8) ToNotBe(target uint8) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint16) ToNotBe(target uint16) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint32) ToNotBe(target uint32) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint64) ToNotBe(target uint64) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}
