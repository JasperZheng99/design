package test

import (
	"design/facade"
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	homePlayer := new(facade.HomePlayerFacade)

	homePlayer.DoKTV()

	fmt.Println("------------")

	homePlayer.DoGame()
}
