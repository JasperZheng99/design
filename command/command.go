package command

import "fmt"

type Chef struct {
}

func (c *Chef) CookFish() {
	fmt.Println("厨师做了一盘鱼")
}

func (c *Chef) CookSoup() {
	fmt.Println("厨师做了一碗汤")
}

type Command interface {
	CommandChefToCook()
}

type CommandCookFish struct {
	chef *Chef
}

func (command *CommandCookFish) SetChef(chef *Chef) {
	command.chef = chef
}

func (command *CommandCookFish) CommandChefToCook() {
	command.chef.CookFish()
}

type CommandCookSoup struct {
	chef *Chef
}

func (command *CommandCookSoup) SetChef(chef *Chef) {
	command.chef = chef
}

func (command *CommandCookSoup) CommandChefToCook() {
	command.chef.CookSoup()
}

// Waiter 服务员，负责收集用户点餐
type Waiter struct {
	commands []Command
}

func (waiter *Waiter) Notify() {
	if waiter.commands == nil {
		fmt.Println("无订单")
		return
	}
	fmt.Println("服务员提交订单")
	for _, cmd := range waiter.commands {
		cmd.CommandChefToCook()
	}
}

func (waiter *Waiter) AppendCommand(cmd Command) {
	waiter.commands = append(waiter.commands, cmd)
}
