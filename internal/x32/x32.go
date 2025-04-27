package x32

import (
	"net"
)

const DefaultPort = "10023"

type Device struct {
	Addr *net.UDPAddr
	Conn *net.UDPConn
}

// NewDevice creates a new X32 connection
func NewDevice(ip string) (*Device, error) {
	remoteAddr, err := net.ResolveUDPAddr("udp", ip+":"+DefaultPort)
	if err != nil {
		return nil, err
	}

	localAddr, err := net.ResolveUDPAddr("udp", ":0")
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		return nil, err
	}

	return &Device{
		Addr: remoteAddr,
		Conn: conn,
	}, nil
}

// Close closes the UDP connection
func (d *Device) Close() error {
	return d.Conn.Close()
}
