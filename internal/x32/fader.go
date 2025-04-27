package x32

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
	"unsafe"
)

// SendFader sets a fader value
func (d *Device) SendFader(addr string, value float32) error {
	var buf bytes.Buffer

	buf.WriteString(addr)
	padding := 4 - (len(addr) % 4)
	if padding != 4 {
		buf.Write(make([]byte, padding))
	}

	buf.WriteString(",f")
	buf.Write([]byte{0, 0})

	binary.Write(&buf, binary.BigEndian, value)

	_, err := d.Conn.Write(buf.Bytes())
	return err
}

// RequestFader asks for a fader value
func (d *Device) RequestFader(addr string) error {
	var buf bytes.Buffer

	buf.WriteString(addr)
	buf.WriteByte(0) // null terminator

	padding := (4 - (len(addr)+1)%4) % 4
	if padding != 0 {
		buf.Write(make([]byte, padding))
	}

	_, err := d.Conn.Write(buf.Bytes())
	return err
}

// ReceiveFaderReply reads a fader reply
func (d *Device) ReceiveFaderReply() (string, float32, error) {
	_ = d.Conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
	buf := make([]byte, 1024)
	n, _, err := d.Conn.ReadFromUDP(buf)
	if err != nil {
		return "", 0, err
	}

	payload := buf[:n]
	addrEnd := bytes.IndexByte(payload, 0)
	if addrEnd == -1 {
		return "", 0, fmt.Errorf("invalid OSC address")
	}
	addr := string(payload[:addrEnd])

	typeStart := (addrEnd + 4) &^ 3
	if typeStart+2 >= len(payload) || payload[typeStart] != ',' || payload[typeStart+1] != 'f' {
		return "", 0, fmt.Errorf("invalid type tag")
	}

	dataStart := typeStart + 4
	if dataStart+4 > len(payload) {
		return "", 0, fmt.Errorf("missing float32 data")
	}
	valBits := binary.BigEndian.Uint32(payload[dataStart : dataStart+4])
	val := float32FromBits(valBits)

	return addr, val, nil
}

func float32FromBits(bits uint32) float32 {
	return *(*float32)(unsafe.Pointer(&bits))
}
