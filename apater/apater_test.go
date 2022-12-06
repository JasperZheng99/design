package apater

import (
	"fmt"
	"testing"
)

// 继承TypeC接口的安卓充电线
type AndroidTypeC struct {
}

func (chargeType *AndroidTypeC) UseTypeC() {

	fmt.Println("使用安卓充电线")
}

func TestApater(t *testing.T) {

	iphone14 := NewIPhone(NewApater(new(AndroidTypeC)))
	iphone14.Charge()
}
