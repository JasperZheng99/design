package state

import (
	"fmt"
)

type State interface {
	Work()
	GetHumen() Humen
	SetHumen(Humen)
}

type AbstractState struct {
	humen Humen
}

func (state *AbstractState) Work() {

}

func (state *AbstractState) GetHumen() Humen {
	return state.humen
}

func (state *AbstractState) SetHumen(humen Humen) {
	state.humen = humen
}

type Humen interface {
	Work()
	SetState(State)
	GetName() string
	SetName(string)
}

type AbstractHumen struct {
	state State
	name  string
}

func (humen *AbstractHumen) Work() {
	humen.state.SetHumen(humen)
	humen.state.Work()
}

func (humen *AbstractHumen) SetState(s State) {
	humen.state = s
}
func (humen *AbstractHumen) GetName() string {
	return humen.name
}

func (humen *AbstractHumen) SetName(name string) {
	humen.name = name
}

// 儿童时期
type ChildState struct {
	AbstractState
}

func (state *ChildState) Work() {
	fmt.Println(state.humen.GetName(), "小孩子就是玩")
	state.humen.SetState(new(AdultState))
}

// 成年人
type AdultState struct {
	AbstractState
}

func (state *AdultState) Work() {
	fmt.Println(state.humen.GetName(), "成年人工作赚钱")
	state.humen.SetState(new(OldState))
}

// 老年人
type OldState struct {
	AbstractState
}

func (state *OldState) Work() {
	fmt.Println(state.humen.GetName(), "老年人养老")
}
