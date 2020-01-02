# gophertest
Simple and intuitive test assertions for testing Go. You've tried the rest, now try this one.

<br />

# Version 0.4.1
I've created this little assertion package for two reasons:
  * take a personal tour of Go!
  * create an assertion syntax that I like.

<br />

# Step 1: Install
`go get github.com/tarrball/gophertest`

<br />

# Step 2: Import
`import . "github.com/tarrball/gophertest/test"`

<br />

# Step 3: Use it (optional: _enjoy_)!
 
    package pkg

    import "testing"
    import . "github.com/tarrball/gophertest/test"

    func TestFoo(t *testing.T) {
        actual := 5
        expected := 5
        
        Test(t).If(actual).Is(expected)
    }

<br />

# Contact
Thoughts? Opinions? Requests? Suggestions?
<andrew@tarrball.com>
