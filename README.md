# mint

The very minimum assertion for Go.

[![Build Status](https://travis-ci.org/otiai10/mint.svg?branch=master)](https://travis-ci.org/otiai10/mint)

# usage
```go
package your_test

import "your"
import "testing"
import "github.com/otiai10/mint"

func TestYour_SomeFunc(t *testing.T) {
    mint.Expect(t, 1).ToBe(1)
}
```
