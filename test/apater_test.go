package test

import (
	"design/apater"
	"fmt"
	"testing"
)

// 继承TypeC接口的安卓充电线
type AndoridTypeC struct {
}

func (chargeType *AndoridTypeC) UseTypeC() {

	fmt.Println("使用安卓充电线")
}

func TestApater(t *testing.T) {

	iphone14 := apater.NewIPhone(apater.NewApater(new(AndoridTypeC)))
	iphone14.Charge()
}
