curr
=====

> This package was moved from https://github.com/otiai10/mint/curr

Current file and dir privider for Golang.

Just a sugar for [runtime](https://golang.org/pkg/runtime/).

```go
import "curr"

// __FILE__
f := curr.File()

// __DIR__
d := curr.Dir()

// __LINE__
l := curr.Line()

// __FUNCTION__
fn := curr.Func()

// basename(__FILE__)
b := curr.Basename()
```


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fotiai10%2Fcurr.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fotiai10%2Fcurr?ref=badge_large)