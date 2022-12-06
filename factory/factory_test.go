package factory

import (
	"testing"
)

/**

● 简单工厂模式：一个工厂负责创建所有产品
  ○ 违反“开闭原则”，添加新产品需要修改工厂逻辑，工厂越来越复杂
● 工厂方法模式：一个工厂创建一个产品
  ○ 系统的可扩展性也就变得非常好，无需修改接口和原类
  ○ 增加系统中类的个数，复杂度和理解度增加（一个具体产品就需要对应一个具体工厂）
● 抽象方法模式：一个工厂创建一系列（一个产品族）的产品
  ○ 增加新的产品族很方便，无须修改已有系统，符合“开闭原则”
  ○ 增加新的产品等级结构麻烦，需要对原有系统进行较大的修改，违背了“开闭原则”
  ○ 相当于在工厂方法模式的基础下进行了折中
    ■ 对于产品族来说遵循了开闭原则
    ■ 对于产品等级结构来说没有遵循开闭原则
    ■ 如果产品结构等级稳定，那么就相当于完全遵循开闭原则

**/

// 简单工厂模式
func TestSimpleFactory(t *testing.T) {
	fac := SimpleFactory{}
	football := fac.CreatFactory("football")
	football.Play()

	basketball := fac.CreatFactory("basketball")
	basketball.Play()
}

// 工厂方法测试
func TestFactory(t *testing.T) {
	var appFactory Factory = new(AppleFactory)

	apple := appFactory.CreatFruit()

	apple.Show()

	var bananaFactory Factory = new(BananaFactory)
	banana := bananaFactory.CreatFruit()
	banana.Show()
}

// 抽象工厂测试
func TestAbstractFactory(t *testing.T) {
	var chineseFac AbstractFactory = new(ChineseFactory)
	shoes := chineseFac.CreatShoes()
	shoes.ShowShoesArea()

	var japaneseFac AbstractFactory = new(JananeseFactory)
	hat := japaneseFac.CreatHat()
	hat.ShowHatArea()
}
