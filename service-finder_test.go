package servicefinder

import (
	"testing"
)

type Doer interface {
	Do()
}
type ConcreteDo struct{}

func (c *ConcreteDo) Do() {}

func Test_FindByInterface(t *testing.T) {

	finder := New()

	finder.Register(&ConcreteDo{})
	var similar Doer
	finder.MustGet(&similar)
}

func Test_FindByConcreteType(t *testing.T) {

	finder := New()

	finder.Register(&ConcreteDo{})

	similar := &ConcreteDo{}
	finder.MustGet(&similar)
}
func Test_FindByConcreteType2(t *testing.T) {

	finder := New()

	finder.Register(&ConcreteDo{})

	var similar *ConcreteDo
	finder.MustGet(&similar)
}
