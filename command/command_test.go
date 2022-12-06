package command

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	chef := new(Chef)
	finshCommand := new(CommandCookFish)
	finshCommand.SetChef(chef)
	fmt.Println("顾客点了一盘鱼")

	soupCommand := new(CommandCookSoup)
	soupCommand.SetChef(chef)
	fmt.Println("顾客点了一碗汤")
	fmt.Println("-------------------")

	waiter := new(Waiter)
	waiter.AppendCommand(finshCommand)
	waiter.AppendCommand(soupCommand)
	waiter.Notify()

}
