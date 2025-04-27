package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/frimley-baptist-church/x32-tools/internal/x32client"
)

func main() {
	filePath := flag.String("file", "", "snapshot file to read")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: x32-restore-main --file snapshot.json <x32-ip-address>")
		os.Exit(1)
	}

	if *filePath == "" {
		fmt.Println("Error: --file argument is required")
		os.Exit(1)
	}

	ipaddr := flag.Arg(0)

	device, err := x32client.ConnectAndPrepare(ipaddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer device.Close()

	// Read the snapshot file
	data, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		os.Exit(1)
	}

	var snapshot map[string]float32
	err = json.Unmarshal(data, &snapshot)
	if err != nil {
		fmt.Println("Failed to parse snapshot:", err)
		os.Exit(1)
	}

	mainFaderAddr := "/main/st/mix/fader"

	value, ok := snapshot[mainFaderAddr]
	if !ok {
		fmt.Printf("Snapshot does not contain %s\n", mainFaderAddr)
		os.Exit(1)
	}

	fmt.Printf("Restoring Main LR fader to %.3f\n", value)

	err = device.SendFader(mainFaderAddr, value)
	if err != nil {
		fmt.Println("Failed to send fader value:", err)
		os.Exit(1)
	}

	fmt.Println("Done.")
}
