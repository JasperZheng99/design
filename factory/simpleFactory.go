package factory

import "fmt"

type Ball interface {
	Play()
}

type Baseketball struct {
}

func (ball *Baseketball) Play() {
	fmt.Println("play baseketball!!!")
}

type Football struct {
}

func (ball *Football) Play() {
	fmt.Println("play football!!!")
}

type SimpleFactory struct {
}

func (factory *SimpleFactory) CreatFactory(ballType string) Ball {

	var ball Ball

	if ballType == "baseketball" {
		ball = &Baseketball{}
	} else if ballType == "football" {
		ball = &Football{}
	}
	return ball
}
