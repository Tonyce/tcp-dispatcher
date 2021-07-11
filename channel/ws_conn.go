package channel

import (
	"log"

	"github.com/gorilla/websocket"
)

type WsConn struct {
	valid bool
	conn  *websocket.Conn
}

func NewWsConn(conn *websocket.Conn) *WsConn {
	return &WsConn{
		valid: true,
		conn:  conn,
	}
}

func (this *WsConn) Read() (string, error) {
	_, data, err := this.conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("error: %v", err)
		}
		this.Close()
		return "", err
	}
	return string(data), nil
}

func (this *WsConn) Write(s string) error {
	if err := this.conn.WriteMessage(websocket.BinaryMessage, []byte(s)); err != nil {
		this.Close()
		return err
	}
	return nil
}

func (this *WsConn) Close() {
	if this.conn != nil {
		this.valid = false
		this.conn.Close()
	}
}

func (this *WsConn) CheckValid() bool {
	return this.valid
}
