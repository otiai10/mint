# mint

The very minimum assertion for Go.

```go
package your_test

import "your"
import "testing"

import . "github.com/otiai10/mint"

func TestYour_SomeFunc(t *testing.T) {

    Expect(t, 1).ToBe(1)
    Expect(t, "foo").TypeOf("string")
    Expect(t, "exists").Not().ToBe(nil)

    Expect(t, your.SomeFunc()).ToBe("My Func!!")
    // If assertion failed, exit 1 with message.

    // You can run assertions without os.Exit
    res := Expect(t, "foo").Dry().ToBe("bar").Result
    // res.OK == false

    // Or if you don't want to write `t` repeatedly,
    m := mint.Blend(t)
    // you can use blended mint
    m.Expect(1).ToBe(1)
}
```

[![Build Status](https://travis-ci.org/otiai10/mint.svg?branch=master)](https://travis-ci.org/otiai10/mint)

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
