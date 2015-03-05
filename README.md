# mint [![Build Status](https://travis-ci.org/otiai10/mint.svg?branch=master)](https://travis-ci.org/otiai10/mint) [![GoDoc](https://godoc.org/github.com/otiai10/mint?status.png)](https://godoc.org/github.com/otiai10/mint) [![Build Status](https://drone.io/github.com/otiai10/mint/status.png)](https://drone.io/github.com/otiai10/mint/latest)

The very minimum assertion for Go.

```go
package your_test

import "your"
import "testing"
import . "github.com/otiai10/mint"

func TestFoo(t *testing.T) {

    foo := your.Foo()
    Expect(t, foo).ToBe(1234)
    Expect(t, foo).TypeOf("int")
    Expect(t, foo).Not().ToBe(nil)

    // If assertion failed, exit 1 with message.
    Expect(t, foo).ToBe("foobarbuz")

    // You can run assertions without os.Exit
    res := Expect(t, foo).Dry().ToBe("bar")
    // res.OK() == false

    // You can ommit repeated `t`.
    m := mint.Blend(t)
    m.Expect(foo).ToBe(1234)
}
```

# features

- Simple syntax
- Loosely coupled
- Plain implementaion
- Under [WTFPL](http://en.wikipedia.org/wiki/WTFPL)

# tests
```
go test ./...
```

# use cases
Projects bellow use `mint`

- [github.com/otiai10/gosseract](https://github.com/otiai10/gosseract/blob/develop/all_test.go)
