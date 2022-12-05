package test

import (
	"design/decorater"
	"testing"
)

func TestDecorator(t *testing.T) {

	var soup decorater.Food = new(decorater.Soup)
	soup.Show()

	soup = decorater.NewSaltDecorator(soup)
	soup.Show()

	soup = decorater.NewWaterDecorator(soup)
	soup.Show()

}
