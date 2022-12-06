package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	homePlayer := new(HomePlayerFacade)

	homePlayer.DoKTV()

	fmt.Println("------------")

	homePlayer.DoGame()
}
