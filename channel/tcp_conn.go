package channel

import (
	"bufio"
	"bytes"
	"io"
	"net"
)

type TcpConn struct {
	valid bool
	conn  net.Conn
}

func NewTcpConn(conn net.Conn) *TcpConn {
	return &TcpConn{
		valid: true,
		conn:  conn,
	}
}

func (this *TcpConn) Read() (string, error) {
	reader := bufio.NewReader(this.conn)
	var buffer bytes.Buffer
	for {
		ba, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				this.valid = false
				return "", err
			}
			return "", err
		}
		buffer.Write(ba)
		if !isPrefix {
			break
		}
	}
	return buffer.String(), nil
}

func (this *TcpConn) Write(s string) error {
	_, err := this.conn.Write(append([]byte(s), DELIMITER))
	return err
}

func (this *TcpConn) Close() {
	if this.conn != nil {
		this.valid = false
		this.conn.Close()
	}
}
func (this *TcpConn) CheckValid() bool {
	return this.valid
}
