# go-expects
Simple and intuitive test assertions for testing Go. You've tried the rest, now try this one.

<br />

# Version 0.1.0
I've created this little assertion package for two reasons:
  * take a personal tour of Go!
  * create an assertion syntax that I like.

<br />

# Installation
`go get github.com/tarrball/go-expects`

<br />

# Documentation
It's supposed to be intuitive right? So how important could documentation _really_ be? Even so, I'll give a little walkthrough. 

Unfortunately, due to the nature of Go testing as I understand it today, it's necessary to pass the test context into each first call:

    func TestExpectIntToBe3(t *testing.T) {
	    expects.Int(t, 3)...
    }
    
After the `expectation` has the test context, you can use the assertions:

    func TestExpectIntToBe3(t *testing.T) {
	    expects.Int(t, 3).toBe(3)
    }
    
Currently, the expectation types `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `String` may use:
* `toBe()`
* `toNotBe()`
* `toBeGreaterThan()`
* `toBeLessThan()`

and `Bool` may use:
* `toBeTrue()`
* `toBeFalse()`
    
And that's it! Pretty sweet, right?

<br />

# Coming Soon
arrays, slices, maps, (expect.toContain, expect.toNotContain, expect.toBeNil, expect.toNotBeNil, etc)

<br />

# Contact
Thoughts? Opinions? Requests? Suggestions?
<andrew@tarrball.com>
