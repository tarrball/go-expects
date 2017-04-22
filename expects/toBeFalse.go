package expects

func (source ExpectationBool) ToBeFalse() {
	if source.value != false {
		source.testContext.Errorf("Expected 'true' to be 'false'")
	}
}
