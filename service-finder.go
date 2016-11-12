package servicefinder

import (
	"reflect"
)

// ServiceFinder is an interface that defines
// the method implemented by a concret service finder.
type ServiceFinder interface {
	Register(some interface{})
	Get(some interface{}) bool
	MustGet(some interface{})
	Each(some interface{}, this func(concrete interface{}))
}

// ServiceRegisry references services to identify.
type ServiceRegistry struct {
	services []interface{}
	types    []reflect.Type
	values   []reflect.Value
}

// New is a ServiceFinder constructor.
func New() ServiceFinder {
	return &ServiceRegistry{}
}

// Register a new service to locate.
func (s *ServiceRegistry) Register(some interface{}) {
	s.services = append(s.services, some)
	s.types = append(s.types, reflect.TypeOf(some))
	s.values = append(s.values, reflect.ValueOf(some))
}

// Get a service of some type,
// returns true on success,
// false on failure.
func (s *ServiceRegistry) Get(some interface{}) bool {
	k := reflect.TypeOf(some).Elem()
	kind := k.Kind()
	if kind == reflect.Ptr {
		k = k.Elem()
		kind = k.Kind()
	}
	for i, t := range s.types {
		if kind == reflect.Interface && t.Implements(k) {
			reflect.Indirect(
				reflect.ValueOf(some),
			).Set(s.values[i])
			return true
		} else if kind == reflect.Struct && k.AssignableTo(t.Elem()) {
			reflect.ValueOf(some).Elem().Set(s.values[i])
			return true
		}
	}
	return false
}

// MustGet panics if a service of some type
// is not found.
func (s *ServiceRegistry) MustGet(some interface{}) {
	if s.Get(some) == false {
		k := reflect.TypeOf(some).Elem().Name()
		panic(k + " service not found")
	}
}

// Each calls this callback with every services of some type found.
func (s *ServiceRegistry) Each(some interface{}, this func(concrete interface{})) {
	k := reflect.TypeOf(some).Elem()
	kind := k.Kind()
	if kind == reflect.Ptr {
		k = k.Elem()
		kind = k.Kind()
	}
	for i, t := range s.types {
		if kind == reflect.Interface && t.Implements(k) {
			this(s.services[i])
		} else if kind == reflect.Struct && k.AssignableTo(t.Elem()) {
			this(s.services[i])
		}
	}
}
