package test

import (
	"design/observer"
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	zhangsan := observer.Student{
		Name:     "张三",
		DoThings: "打王者",
	}
	lihua := observer.Student{
		Name:     "李华",
		DoThings: "刮痧",
	}
	daitou := observer.Student{
		Name:     "呆头",
		DoThings: "看漫画",
	}

	// 班长，通知者
	monitor := observer.Monitor{}
	monitor.AddListener(&zhangsan)
	monitor.AddListener(&lihua)
	monitor.AddListener(&daitou)

	fmt.Println("老师来之前")
	zhangsan.Doing()
	lihua.Doing()
	daitou.Doing()
	fmt.Println("-----------\n老师来了")
	monitor.Notify()

}
