package test

import (
	"design/factory"
	"testing"
)

// 抽象工厂测试
func TestFactory(t *testing.T) {
	appFactory := new(factory.AppleFactory)

	apple := appFactory.CreatFruit()

	apple.Show()

	bananaFactory := new(factory.BananaFactory)
	banana := bananaFactory.CreatFruit()
	banana.Show()
}
