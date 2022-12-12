package flyweight

import "fmt"

// 享元模式

type Dress interface {
	getColor() string
}

// 具体享元对象 恐怖份子衣服
type TerroistDress struct {
	color string
}

func (t *TerroistDress) getColor() string {
	return t.color
}

func NewTerroistDress() *TerroistDress {
	return &TerroistDress{color: "red"}
}

//反恐精英
type CounterTerroistDress struct {
	color string
}

func (t *CounterTerroistDress) getColor() string {

	return t.color
}

func NewCounterTerroistDress() *CounterTerroistDress {
	return &CounterTerroistDress{color: "green"}
}

//---------------享元工厂------------
const (
	TerroristDressType       = "tDress"
	CounterTerroistDressType = "ctDress"
)

var (
	dressFactoruSingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = NewTerroistDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroistDressType {
		d.dressMap[dressType] = NewCounterTerroistDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("找不到相应类型")
}

func GetDressFactorySingleton() *DressFactory {
	return dressFactoruSingleInstance
}

//--------------背景----------------

type Player struct {
	dress      Dress
	playerType string //T 或者 CT
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *Player {
	dress, _ := GetDressFactorySingleton().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
