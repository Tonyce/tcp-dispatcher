package channel

import (
	"log"
	"sync"
	"time"
)

type Channel struct {
	name    string
	mu      sync.Mutex
	clients []Client
}

func NewChannel(name string) *Channel {
	c := &Channel{
		name:    name,
		clients: make([]Client, 0),
	}
	go c.checkInvalid()
	return c
}

func (this *Channel) AddNewClient(client Client) {
	this.mu.Lock()
	this.clients = append(this.clients, client)
	this.mu.Unlock()
}

func (this *Channel) DispatchMsg(msg string, client Client) {
	for i := 0; i < len(this.clients); i++ {
		c := this.clients[i]
		if c != client {
			c.Write(msg)
		}
	}
}

func (this *Channel) checkInvalid() {
	for {
		for i := 0; i < len(this.clients); i++ {
			c := this.clients[i]
			if !c.CheckValid() {
				_, ok := c.(*TcpConn)
				if ok {
					log.Printf("tcpConn: %+v", c)
				}
				_, ok = c.(*WsConn)
				if ok {
					log.Printf("wsConn: %+v", c)
				}
				this.removeInvalid(i)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func (this *Channel) removeInvalid(index int) {
	this.mu.Lock()
	this.clients = append(this.clients[:index], this.clients[index+1:]...)
	this.mu.Unlock()
}
