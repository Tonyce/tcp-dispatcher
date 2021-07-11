package main

import (
	"fmt"
	"net"
	"net/http"
	"tcp-dispatcher/channel"

	"github.com/gorilla/websocket"
)

const (
	TCP_PORT  = ":8765"
	HTTP_PORT = ":7654"
	QUIT_SIGN = "quit!"
)

func main() {

	c := channel.NewChannel("test")
	cm := channel.GetChannelManager()
	cm.AddChannel("test", c)

	go tcpServer()
	go wsServer()
	select {}
}

func tcpServer() {
	listener, _ := net.Listen("tcp", TCP_PORT)
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleChannelTcpConn(conn)
	}
}

func handleChannelTcpConn(conn net.Conn) {
	defer conn.Close()

	tcpClient := channel.NewTcpConn(conn)
	cm := channel.GetChannelManager()
	c := cm.GetChannel("test")
	c.AddNewClient(tcpClient)
	fmt.Println("--==-=-=")
	for {
		content, err := tcpClient.Read()
		if err != nil {
			return
		}
		fmt.Println("recied", content)
		c.DispatchMsg(content, tcpClient)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		process(w, r)
	})
	http.ListenAndServe(HTTP_PORT, nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	if conn, err := upgrader.Upgrade(w, r, nil); err == nil {
		// loginpkgdecode(connection.NewWsConn(conn))
		handleChannelWsConn(conn)
	}
}

func handleChannelWsConn(conn *websocket.Conn) {
	defer conn.Close()

	wsClient := channel.NewWsConn(conn)
	cm := channel.GetChannelManager()
	c := cm.GetChannel("test")
	c.AddNewClient(wsClient)
	for {
		content, err := wsClient.Read()
		if err != nil {
			return
		}
		fmt.Println("recied", content)
		c.DispatchMsg(content, wsClient)
	}
}

/**
func handleTcpConn(conn net.Conn) {
	defer conn.Close()
	var err error

	if err = authPacket(conn); err != nil {
		return
	}

	log.Println("----------")
	for {
		content, err := readLineFromConn(conn)
		if err != nil {
			log.Printf("---- %v\n", err)
			continue
		}
		if content == QUIT_SIGN {
			log.Println("Listener: Quit!")
			break
		}
		log.Printf("received content: %s\n", content)
	}
}


func readLineFromConn(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	var buffer bytes.Buffer
	for {
		ba, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
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

func authPacket(conn net.Conn) error {
	if err := conn.SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil {
		return err
	}
	content, err := readLineFromConn(conn)
	if err != nil {
		if e, ok := err.(net.Error); ok && e.Timeout() {
			fmt.Println("=----This was a net.Error with a Timeout")
			return err
		}
	}
	if content != "ok" {
		return errors.New("not auth")
	}
	conn.SetDeadline(time.Time{})
	return nil
}
*/
