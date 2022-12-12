package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	o := new(Originator)
	o.SetSate("良好")
	fmt.Printf("当前状态：%v\n", o.GetState())

	c := new(Caretaker)
	c.SetMemento(o.CreatMemento())

	// 更改状态
	fmt.Println("更改当前状态")
	o.SetSate("很差")
	fmt.Println("当前状态为：", o.GetState())

	//恢复状态
	o.SetSate(c.GetMemento().GetState())
	fmt.Println("恢复后的状态为：", o.GetState())

}
