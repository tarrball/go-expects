package test

// IsNil expects the system under test value to be nil.
func (sut SUT) IsNil() {
	if sut.actual == nil {
		sut.testContext.Errorf("Expected value was not nil.")
	}
}
