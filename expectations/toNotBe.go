package expects

func (source ExpectationInt) toNotBe(target int) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt8) toNotBe(target int8) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt16) toNotBe(target int16) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt32) toNotBe(target int32) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationInt64) toNotBe(target int64) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationString) toNotBe(target string) {
	if source.value == target {
		source.testContext.Errorf("'%s' should not equal '%s'",
			source.value,
			target)
	}
}

func (source ExpectationUint) toNotBe(target uint) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint8) toNotBe(target uint8) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint16) toNotBe(target uint16) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint32) toNotBe(target uint32) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source ExpectationUint64) toNotBe(target uint64) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}
