package x32client

import (
    "fmt"

    "github.com/frimley-baptist-church/x32-tools/internal/x32"
)

// ConnectAndPrepare connects to the desk and sends /xremote
func ConnectAndPrepare(ip string) (*x32.Device, error) {
    device, err := x32.NewDevice(ip)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to X32 at %s: %w", ip, err)
    }

    if err := device.SendXRemote(); err != nil {
        device.Close()
        return nil, fmt.Errorf("failed to start /xremote session: %w", err)
    }

    return device, nil
}
