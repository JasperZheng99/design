package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {

	var soup Food = new(Soup)
	soup.Show()

	soup = NewSaltDecorator(soup)
	soup.Show()

	soup = NewWaterDecorator(soup)
	soup.Show()

}
