package test

// IsNotNil expects the sytem under test value to be non-nil.
func (sut SUT) IsNotNil() {
	if sut.actual == nil {
		sut.testContext.Errorf("Expected value was nil.")
	}
}
