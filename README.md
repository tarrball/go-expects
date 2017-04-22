# go-expectations
Simple and intuitive test assertions for testing Go.

<br />

# Version 0.0.1
I've created this little assertion package for two reasons:
  * take a personal tour of Go!
  * create an assertion syntax that I like.

<br />

# Installation
`go get github.com/tarrball/go-expectations`

<br />

# Documentation
It's supposed to be intuitive right? So how important could documentation _really_ be? Even so, I'll give a little walkthrough. 

Unfortunately, due to the nature of Go testing as I understand it today, it's necessary to pass the test context into each first call:

    func TestExpectIntToBe3(t *testing.T) {
	    expectInt(t, 3)...
    }
    
After the `expectation` has the test context, you can use the assertions:

    func TestExpectIntToBe3(t *testing.T) {
	    expectInt(t, 3).toBe(3)
    }
    
Currently, the expectations `expectInt`, `expectInt8`, `expectInt16`, `expectInt32`, `expectInt64`, `expectUint`, `expectUint8`, `expectUint16`, `expectUint32`, `expectUint64`, `expectString` may use:
* `toBe()`
* `toNotBe()`
* `toBeGreaterThan()`
* `toBeLessThan()`

and `expectBool' may use:
* `toBeTrue()`
* `toBeFalse()`
    
And that's it! Pretty sweet, right?

<br />

# Coming Soon
arrays, slices, maps, (expect.toContain, expect.toNotContain, expect.toBeNil, expect.toNotBeNil, etc)
