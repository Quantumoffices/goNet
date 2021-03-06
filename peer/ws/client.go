package ws

import (
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	. "github.com/zjllib/goNet"
	"net/http"
	"time"
)

type client struct {
	PeerIdentify
	session *session
}

func init() {
	identify := PeerIdentify{}
	identify.SetType(PeertypeClient)
	c := &client{
		PeerIdentify: identify,
	}
	RegisterPeer(c)
}

func (c *client) Start() {
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 5 * time.Second,
	}
	logs.Info(c.Addr())
	conn, _, err := dialer.Dial(c.Addr(), nil)
	if err != nil {
		panic(err)
	}
	c.session = newSession(conn)
	go c.session.recvLoop()
	go c.session.sendLoop()
}

func (c *client) Stop() {
	c.session.socket.SetReadDeadline(time.Now())
}
