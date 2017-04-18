package goexpectations

func (source expectationInt) toBe(target int) {
	if source.value != target {
		source.context.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source expectationInt8) toBe(target int8) {
	if source.value != target {
		source.context.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}git 
}

func (source expectationInt16) toBe(target int16) {
	if source.value != target {
		source.context.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source expectationInt32) toBe(target int32) {
	if source.value != target {
		source.context.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source expectationInt64) toBe(target int64) {
	if source.value != target {
		source.context.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source expectationUint) toBe(target uint) {
	if source.value != target {
		source.context.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source expectationUint8) toBe(target uint8) {
	if source.value != target {
		source.context.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source expectationUint16) toBe(target uint16) {
	if source.value != target {
		source.context.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source expectationUint32) toBe(target uint32) {
	if source.value != target {
		source.context.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source expectationUint64) toBe(target uint64) {
	if source.value != target {
		source.context.Errorf("Expected '%d' to equal '%d'",
			source.value,
			target)
	}
}

func (source expectationString) toBe(target string) {
	if source.value != target {
		source.context.Errorf("Expected '%s' to equal '%s'",
			source.value,
			target)
	}
}
