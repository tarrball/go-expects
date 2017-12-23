package test

type testContext interface {
	Errorf(format string, args ...interface{})
}
