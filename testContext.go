package goexpectations

type testContext interface {
	Errorf(format string, args ...interface{})
}
