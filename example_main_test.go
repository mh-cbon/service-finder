package servicefinder_test

import (
	"fmt"
	"github.com/mh-cbon/service-finder"
)

type Doer interface {
	Do()
}
type ConcreteDo struct{}

func (c *ConcreteDo) Do() {
	fmt.Printf("Did something")
}

// Example_main demonstrates usage of servicefinder package.
func Example_main() {

	finder := servicefinder.New()

	finder.Register(&ConcreteDo{})

	var concrete *ConcreteDo
	finder.MustGet(&concrete)
	concrete.Do()

	var intface Doer
	finder.MustGet(&intface)
	intface.Do()

}
