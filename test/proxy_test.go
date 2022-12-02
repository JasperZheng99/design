package test

import (
	"design/proxy"
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {

	fmt.Println("====================\n不使用代购")

	good := &proxy.Good{Name: "iphone14"}

	personalShopper := new(proxy.PesonalShopper)
	personalShopper.Buy(good)

	fmt.Println("====================\n使用代购")
	proxyShopper := &proxy.ProxyShopper{personalShopper}

	proxyShopper.Buy(good)
}