package x32

import (
	"bytes"
)

// SendXRemote sends the /xremote OSC message to start remote session
func (d *Device) SendXRemote() error {
	var buf bytes.Buffer

	addr := "/xremote"
	buf.WriteString(addr)
	buf.WriteByte(0) // null-terminate

	// pad to 4-byte boundary after null
	padding := (4 - (len(addr)+1)%4) % 4
	if padding != 0 {
		buf.Write(make([]byte, padding))
	}

	_, err := d.Conn.Write(buf.Bytes())
	return err
}
