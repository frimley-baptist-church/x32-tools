package main

import (
	"fmt"
	"os"

	"github.com/frimley-baptist-church/x32-tools/internal/x32client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: x32-zero-main <x32-ip-address>")
		os.Exit(1)
	}

	ipaddr := os.Args[1]

	device, err := x32client.ConnectAndPrepare(ipaddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer device.Close()

	fmt.Println("Zeroing Main LR fader...")

	err = device.SendFader("/main/st/mix/fader", 0.0)
	if err != nil {
		fmt.Println("Failed to send fader zero:", err)
		os.Exit(1)
	}

	fmt.Println("Done.")
}
