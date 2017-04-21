package goexpectations

func (source expectationBool) toBeFalse() {
	if source.value != false {
		source.testContext.Errorf("Expected 'true' to be 'false'")
	}
}
