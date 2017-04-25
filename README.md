# go-expects
Simple and intuitive test assertions for testing Go. You've tried the rest, now try this one.

<br />

# Version 0.2.0
I've created this little assertion package for two reasons:
  * take a personal tour of Go!
  * create an assertion syntax that I like.

<br />

# Step 1: Install
`go get github.com/tarrball/go-expects`

<br />

# Step 2: Import
`import github.com/tarrball/go-expects/expects`

<br />

# Step 3: Use it (optional: _enjoy_)!
It's supposed to be intuitive right? So how important could documentation _really_ be? Even so, I'll give a little walkthrough. 

Unfortunately, due to the nature of Go testing as I understand it today, it's necessary to pass the test context into each first call:

    func TestExpectIntToBe3(t *testing.T) {
	    expects.Int(t, 3)...
    }
    
After the `expectation` has the test context, you can use the assertions:

    func TestExpectIntToBe3(t *testing.T) {
	    expects.Int(t, 3).ToBe(3)
    }
    
And that's it! Pretty sweet, right?
    
Currently, the expectation types `Float`, `Float32`, `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `String` may use:
* `ToBe()`
* `ToNotBe()`
* `ToBeGreaterThan()`
* `ToBeLessThan()`

and `Bool` may use:
* `ToBeTrue()`
* `ToBeFalse()`

<br />

# Coming Soon
arrays, slices, maps, (expect.toContain, expect.toNotContain, expect.toBeNil, expect.toNotBeNil, etc)

<br />

# Contact
Thoughts? Opinions? Requests? Suggestions?
<andrew@tarrball.com>
