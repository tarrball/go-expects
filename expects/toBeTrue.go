package expects

func (source ExpectationBool) toBeTrue() {
	if source.value != true {
		source.testContext.Errorf("Expected 'false' to be 'true'")
	}
}
