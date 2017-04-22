package expectations

func (source expectationBool) toBeTrue() {
	if source.value != true {
		source.testContext.Errorf("Expected 'false' to be 'true'")
	}
}
