package apater

import "fmt"

// ---------------------安卓充电线------------------------
type TypeC interface {
	UseTypeC()
}

// ---------------------苹果手机线------------------------
type Lighting interface {
	UseLighting()
}

// ---------------------安卓-苹果转接头------------------------
type Apater struct {
	typeC TypeC
}

func (apater *Apater) UseLighting() {
	fmt.Println("type-c转lighting")
	apater.typeC.UseTypeC()
}

func NewApater(typeC TypeC) *Apater {
	return &Apater{typeC: typeC}
}

// ---------------------苹果手机------------------------
type IPhone struct {
	ChargeType Lighting
}

func NewIPhone(c Lighting) *IPhone {
	return &IPhone{c}
}

func (phone *IPhone) Charge() {
	fmt.Println("IPhone 进行充电 ")
	phone.ChargeType.UseLighting()
}
