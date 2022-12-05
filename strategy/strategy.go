package strategy

import "fmt"

// 销售策略
type SellStrategy interface {
	//根据原价得到售卖价
	GetPrice(price float64) float64
}

// 销售策略A，按八折出售
type StrategyA struct {
}

func (s *StrategyA) GetPrice(price float64) float64 {
	return price * 0.8
}

// 销售策略B，满200-100
type StrategyB struct {
}

func (s *StrategyB) GetPrice(price float64) float64 {
	if price >= 200 {
		return price - 100
	}
	return price
}

// 环境
type Goods struct {
	Price    float64
	strategy SellStrategy
}

func (goods *Goods) SetStrategy(s SellStrategy) {
	goods.strategy = s
}

func (goods *Goods) SellGoods() float64 {
	fmt.Println("物品原价为:", goods.Price, "元")
	return goods.strategy.GetPrice(float64(goods.Price))
}
