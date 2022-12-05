package test

import (
	"design/strategy"
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	goods := strategy.Goods{
		Price: 200,
	}

	fmt.Println("上午采用策略A")
	goods.SetStrategy(new(strategy.StrategyA))
	fmt.Println("打八折出售,价格为：", goods.SellGoods(), "元")

	fmt.Println("下午采用策略B")
	goods.SetStrategy(new(strategy.StrategyB))
	fmt.Println("满200-100，价格为", goods.SellGoods(), "元")
}
