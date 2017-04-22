package expects

type testContext interface {
	Errorf(format string, args ...interface{})
}
