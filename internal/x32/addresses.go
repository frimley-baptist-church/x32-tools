package x32

import (
	"fmt"
)

func faderAddresses() []string {
	var addresses []string

	for ch := 1; ch <= 32; ch++ {
		addresses = append(addresses, fmt.Sprintf("/ch/%02d/mix/fader", ch))
	}

	for ch := 1; ch <= 8; ch++ {
		addresses = append(addresses, fmt.Sprintf("/auxin/%02d/mix/fader", ch))
	}

	for fx := 1; fx <= 4; fx++ {
		addresses = append(addresses, fmt.Sprintf("/fxrtn/%02d/mix/fader", fx))
	}

	return addresses
}

func mainFaderAddress() string {
	return "/main/st/mix/fader"
}
