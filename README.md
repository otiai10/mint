# mint

The very minimum assertion for Go.

[![Build Status](https://travis-ci.org/otiai10/mint.svg?branch=master)](https://travis-ci.org/otiai10/mint)

# features

- Simple syntax
- Loosely coupled
- Plain implementaion
- Under [WTFPL](http://en.wikipedia.org/wiki/WTFPL)

# usage
```go
package your_test

import "your"
import "testing"
// Import globally
import . "github.com/otiai10/mint"
// Or import by package name such as
// import "github.com/otiai10/mint"

func TestYour_SomeFunc(t *testing.T) {

    Expect(t, 1).ToBe(1)

    Expect(t, "foo").TypeOf("string")

    Expect(t, "exists").Not().ToBe(nil)

    Expect(t, your.SomeFunc()).ToBe("My Func!!")
    // If assertion failed, exit 1 with message.

    // You can run assertions without os.Exit
    res := Expect(t, "foo").Dry().ToBe("bar").Result
    // res.OK == false
}
```

# contributions
Behavior Driven Development itself
```sh
git clone github.com/otia10/mint
cd mint
go test ./...
```
