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
