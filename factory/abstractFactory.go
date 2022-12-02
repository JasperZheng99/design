package factory

import "fmt"

type AbstractShoes interface {
	ShowShoesArea()
}

type AbstractHat interface {
	ShowHatArea()
}

type AbstractFactory interface {
	CreatShoes() AbstractShoes
	CreatHat() AbstractHat
}

type ChineseShoes struct {
}

func (shoes *ChineseShoes) ShowShoesArea() {
	fmt.Println("make shoes in China")
}

type ChineseHat struct {
}

func (shoes *ChineseHat) ShowHatArea() {
	fmt.Println("make hat in China")
}

type ChineseFactory struct {
}

func (factory *ChineseFactory) CreatShoes() AbstractShoes {
	return &ChineseShoes{}
}

func (factory *ChineseFactory) CreatHat() AbstractHat {
	return &ChineseHat{}
}

type JananeseShoes struct {
}

func (shoes *JananeseShoes) ShowShoesArea() {
	fmt.Println("make shoes in Japan")
}

type JananeseHat struct {
}

func (shoes *JananeseHat) ShowHatArea() {
	fmt.Println("make hat in Japan")
}

type JananeseFactory struct {
}

func (factory *JananeseFactory) CreatShoes() AbstractShoes {
	return &JananeseShoes{}
}

func (factory *JananeseFactory) CreatHat() AbstractHat {
	return &JananeseHat{}
}
