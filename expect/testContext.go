package expect

type testContext interface {
	Errorf(format string, args ...interface{})
}
