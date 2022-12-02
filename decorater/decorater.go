package decorater

import "fmt"

// eg: gin框架的中间间便是通过装饰器实现，一个执行完便执行下一个

// 与代理模式相比，装饰器倾向于添加增强功能，而代理模式倾向于控制对象的访问
// 一个类可能会经过多个装饰器，但一般不会经过多层代理
// ----------- 抽象层 -----------
type Food interface {
	Show()
}

type Decorater struct {
	food Food
}

func (decorater *Decorater) Show() {

}

// ----------- 实现层 -----------
type Soup struct {
}

func (soup *Soup) Show() {
	fmt.Println("做一碗汤")
}

type SaltDecorater struct {
	Decorater
}

func (decorater *SaltDecorater) Show() {
	decorater.food.Show()
	fmt.Println("加点盐")
}

func NewSaltDecorater(food Food) Food {
	return &SaltDecorater{Decorater{food: food}}
}

type WaterDecorater struct {
	Decorater
}

func (decorater *WaterDecorater) Show() {
	decorater.food.Show()
	fmt.Println("加点水")
}

func NewWaterDecorater(food Food) Food {
	return &WaterDecorater{Decorater{food: food}}
}
