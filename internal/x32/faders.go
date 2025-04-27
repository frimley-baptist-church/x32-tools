package x32

import (
	"fmt"
	"time"
)

func (d *Device) ZeroFaders() error {
	addrs := faderAddresses()

	fmt.Println("Zeroing channels...")

	for _, addr := range addrs {
		if err := d.SendFader(addr, 0.0); err != nil {
			fmt.Println("Warning: failed to zero", addr, ":", err)
		}
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("Fading Main LR down...")
	return d.FadeMainLRDown(2 * time.Second)
}

func (d *Device) FadeMainLRDown(duration time.Duration) error {
	addr := mainFaderAddress()
	startVal := float32(1.0)
	steps := 50
	delay := duration / time.Duration(steps)

	for i := 0; i <= steps; i++ {
		f := startVal * (1.0 - float32(i)/float32(steps))
		if f < 0 {
			f = 0
		}
		if err := d.SendFader(addr, f); err != nil {
			return err
		}
		time.Sleep(delay)
	}

	return nil
}
