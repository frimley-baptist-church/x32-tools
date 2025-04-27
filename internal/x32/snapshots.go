package x32

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Snapshot map[string]float32

// StoreSnapshot saves all fader values into a file
func (d *Device) StoreSnapshot(filePath string) error {
	addrs := append(faderAddresses(), mainFaderAddress())
	snapshot := make(Snapshot)

	fmt.Println("Requesting fader values...")

	for _, addr := range addrs {
		if err := d.RequestFader(addr); err != nil {
			fmt.Println("Warning: couldn't request", addr, ":", err)
			continue
		}

		replyAddr, val, err := d.ReceiveFaderReply()
		if err != nil {
			fmt.Println("Warning: no reply for", addr, ":", err)
			continue
		}

		replyAddr = strings.TrimRight(replyAddr, "\x00")
		if replyAddr != addr {
			fmt.Println("Warning: unexpected reply", replyAddr, "for", addr)
			continue
		}

		snapshot[addr] = val
		time.Sleep(10 * time.Millisecond)
	}

	data, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// RestoreSnapshot applies a snapshot
func (d *Device) RestoreSnapshot(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var snapshot Snapshot
	err = json.Unmarshal(data, &snapshot)
	if err != nil {
		return err
	}

	fmt.Println("Restoring faders...")

	for addr, val := range snapshot {
		if err := d.SendFader(addr, val); err != nil {
			fmt.Println("Warning: failed to send", addr, ":", err)
		}
		time.Sleep(10 * time.Millisecond)
	}

	return nil
}
