package proxy

import "fmt"

// 代理模式
type Good struct {
	Name string //物品名称
}

type Shopper interface {
	Buy(good *Good)
}

type PesonalShopper struct {
}

func (shoper *PesonalShopper) Buy(good *Good) {
	fmt.Println("购买了" + good.Name)
}

type ProxyShopper struct {
	Shopper Shopper
}

func (shopper *ProxyShopper) Buy(good *Good) {
	fmt.Println("我帮你代购")
	shopper.Shopper.Buy(good)
	fmt.Println("买完了")
}
