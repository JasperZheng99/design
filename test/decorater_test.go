package test

import (
	"design/decorater"
	"testing"
)

func TestDecorater(t *testing.T) {

	var soup decorater.Food = new(decorater.Soup)
	soup.Show()

	soup = decorater.NewSaltDecorater(soup)
	soup.Show()

	soup = decorater.NewWaterDecorater(soup)
	soup.Show()

}
