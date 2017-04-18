package goexpectations

func (source expectationString) toNotBe(target string) {
	if source.value == target {
		source.context.Errorf("Expected '%s' to not equal '%s'",
			source.value,
			target)
	}
}

func (source expectationInt) toNotBe(target int) {
	if source.value == target {
		source.context.Errorf("Expected '%d' to not equal '%d'",
			source.value,
			target)
	}
}
