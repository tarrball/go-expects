package expectations

type testContext interface {
	Errorf(format string, args ...interface{})
}
