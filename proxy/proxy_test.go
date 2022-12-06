package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {

	fmt.Println("====================\n不使用代购")

	good := &Good{Name: "iphone14"}

	personalShopper := new(PersonalShopper)
	personalShopper.Buy(good)

	fmt.Println("====================\n使用代购")
	proxyShopper := &ProxyShopper{Shopper: personalShopper}

	proxyShopper.Buy(good)
}
