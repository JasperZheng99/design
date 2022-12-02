package factory

/**
工厂方法模式
*/
import "fmt"

type Fruit interface {
	Show()
}

type Factory interface {
	CreatFruit() Fruit
}

type Apple struct {
}

func (a *Apple) Show() {
	fmt.Println("this is apple!")
}

type Banana struct {
}

func (b *Banana) Show() {
	fmt.Println("this is banana!")
}

type AppleFactory struct {
}

func (receiver *AppleFactory) CreatFruit() Fruit {
	return &Apple{}
}

type BananaFactory struct {
}

func (receiver *BananaFactory) CreatFruit() Fruit {
	return &Banana{}
}
