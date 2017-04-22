package expects

func (source ExpectationBool) ToBeTrue() {
	if source.value != true {
		source.testContext.Errorf("Expected 'false' to be 'true'")
	}
}
