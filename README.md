# go-expect
Simple and intuitive test assertions for testing Go. You've tried the rest, now try this one.

<br />

# Version 0.3.0
I've created this little assertion package for two reasons:
  * take a personal tour of Go!
  * create an assertion syntax that I like.

<br />

# Step 1: Install
`go get github.com/tarrball/go-expect`

<br />

# Step 2: Import
`import github.com/tarrball/go-expect/expect`

<br />

# Step 3: Use it (optional: _enjoy_)!
It's supposed to be intuitive right? So how important could documentation _really_ be? Even so, I'll give a little walkthrough. 

Unfortunately, due to the nature of Go testing as I understand it today, it's necessary to pass the test context into each first call:

    func TestExpectIntToBe3(t *testing.T) {
	    expect.This(t, 3)...
    }
    
After the `expectation` has the test context, you can use the assertions:

    func TestExpectIntToBe3(t *testing.T) {
	    expect.This(t, 3).ToBe(3)
    }
    
And that's it! Pretty sweet, right?
Currently the supported assertions are: 

All types:
* `ToBe()`
* `ToNotBe()`

Strings and numeric types:
* `ToBeGreaterThan()`
* `ToBeLessThan()`

Boolean:
* `ToBeTrue()`
* `ToBeFalse()`

<br />

# Coming Soon
arrays, slices, maps, (ToContain, ToNotContain, ToBeNil, ToNotBeNil, etc)

<br />

# Contact
Thoughts? Opinions? Requests? Suggestions?
<andrew@tarrball.com>
