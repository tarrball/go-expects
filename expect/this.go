package expect

import "reflect"

// SUT System Under Test contains value and testContext
type SUT struct {
	testContext
	value interface{}
}

// This takes a testContext and an object to be asserted on
func This(t testContext, actual interface{}) SUT {
	return SUT{testContext: t, value: actual}
}

// ConversionFail fails the test and outputs a message with the types
func ConversionFail(t testContext, actualType reflect.Type, expectedType reflect.Type) {
	t.Errorf("Could not convert actual type '%s' to expected type '%s'",
		actualType,
		expectedType)
}

type TTest struct {
	actual int
	testContext
}

// Test does stuff
func Test(t testContext) TTest {
	return TTest{0, t}
}

func (context TTest) If(actual int) TTest {
	context.actual = actual
	return context
}

func (context TTest) Is(expected int) {
	if context.actual != expected {
		context.Errorf("You're a failure.")
	}
}

type Factory struct {
	service Service
}

type Service interface {
	DoSomething(int)
}

type RealService struct{}

func (r RealService) DoSomething(i int) {
	println(string(i))
}

func (f Factory) AddProxy(s Service) {
	f.service = s
}

func (f Factory) ResolveProxy() Service {
	return f.service
}
