package test

import (
	"design/command"
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	chef := new(command.Chef)
	finshCommand := new(command.CommandCookFish)
	finshCommand.SetChef(chef)
	fmt.Println("顾客点了一盘鱼")

	soupCommand := new(command.CommandCookSoup)
	soupCommand.SetChef(chef)
	fmt.Println("顾客点了一碗汤")
	fmt.Println("-------------------")

	waiter := new(command.Waiter)
	waiter.AppendCommand(finshCommand)
	waiter.AppendCommand(soupCommand)
	waiter.Notify()

}
