package goexpectations

func (source expectationInt) toBeLessThan(target int) {
	if source.value >= target {
		source.testContext.Errorf("Expected '%d' to be less than '%d'",
			source.value,
			target)
	}
}
