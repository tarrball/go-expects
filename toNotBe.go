package goexpectations

func (source expectationInt) toNotBe(target int) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source expectationInt8) toNotBe(target int8) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source expectationInt16) toNotBe(target int16) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source expectationInt32) toNotBe(target int32) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source expectationInt64) toNotBe(target int64) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source expectationString) toNotBe(target string) {
	if source.value == target {
		source.testContext.Errorf("'%s' should not equal '%s'",
			source.value,
			target)
	}
}

func (source expectationUint) toNotBe(target uint) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source expectationUint8) toNotBe(target uint8) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source expectationUint16) toNotBe(target uint16) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source expectationUint32) toNotBe(target uint32) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}

func (source expectationUint64) toNotBe(target uint64) {
	if source.value == target {
		source.testContext.Errorf("'%d' should not equal '%d'",
			source.value,
			target)
	}
}
