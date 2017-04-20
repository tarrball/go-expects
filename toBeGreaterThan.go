package goexpectations

func (source expectationInt) toBeGreaterThan(target int) {
	if source.value <= target {
		source.testContext.Errorf("Expected '%d' to be greater than '%d'",
			source.value,
			target)
	}
}
