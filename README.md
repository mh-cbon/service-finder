# Service finder

[![GoDoc](https://godoc.org/github.com/mh-cbon/service-finder?status.svg)](https://godoc.org/github.com/mh-cbon/service-finder)

Register a concrete service, then locate it by an interface or its concrete type.

# Install

```sh
go get github.com/mh-cbon/service-finder
glide install github.com/mh-cbon/service-finder
```

# Example

```go
package main

import (
  "fmt"
  "github.com/mh-cbon/service-finder"
)

type Doer interface {
  Do()
}
type ConcreteDo struct {}
func (c *ConcreteDo) Do() {
  fmt.Println("Did something")
}

// Example_main demonstrates usage of servicefinder package.
func main() {

    finder := servicefinder.New()

    finder.Register(&ConcreteDo{})

    var concrete *ConcreteDo
    finder.MustGet(&concrete)
    concrete.Do()

    var intface Doer
    finder.MustGet(&intface)
    intface.Do()

}
```
