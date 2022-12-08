package state

import (
	"testing"
)

type testHumen struct {
	AbstractHumen
}

func TestState(t *testing.T) {
	zhang3 := new(testHumen)
	zhang3.SetName("张三")
	zhang3.SetState(new(ChildState))
	zhang3.Work()
	zhang3.Work()
	zhang3.Work()

	li4 := new(testHumen)
	li4.SetName("李四")
	li4.SetState(new(AdultState))
	li4.Work()
	li4.Work()
}
