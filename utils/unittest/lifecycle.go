package unittest

import (
	"github.com/stretchr/testify/mock"
)

// ReadyDoneify sets up a generated mock to respond to Ready and Done
// lifecycle methods. Any mock type generated by mockery can be used.
func ReadyDoneify(toMock interface{}) {

	mockable, ok := toMock.(interface {
		On(string, ...interface{}) *mock.Call
	})
	if !ok {
		panic("attempted to mock invalid type")
	}

	rwch := make(chan struct{})
	var ch <-chan struct{} = rwch
	close(rwch)

	mockable.On("Ready").Return(ch).Maybe()
	mockable.On("Done").Return(ch).Maybe()
}
