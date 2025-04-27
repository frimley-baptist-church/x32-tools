package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/frimley-baptist-church/x32-tools/internal/x32client"
)

func main() {
	storeFile := flag.String("store", "", "store snapshot to file")
	restoreFile := flag.String("restore", "", "restore snapshot from file")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: x32-snapshot [--store file.json|--restore file.json] <x32-ip-address>")
		os.Exit(1)
	}

	ipaddr := flag.Arg(0)

	device, err := x32client.ConnectAndPrepare(ipaddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer device.Close()

	if *storeFile != "" {
		if err := device.StoreSnapshot(*storeFile); err != nil {
			fmt.Println("Store failed:", err)
			os.Exit(1)
		}
	}

	if *restoreFile != "" {
		if err := device.RestoreSnapshot(*restoreFile); err != nil {
			fmt.Println("Restore failed:", err)
			os.Exit(1)
		}
	}
}
