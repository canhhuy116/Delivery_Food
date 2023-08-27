package skio

import (
	"Delivery_Food/common"
	"net"
	"net/http"
	"net/url"
)

type Conn interface {
	ID() string
	Close() error
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	Context() interface{}
	SetContext(v interface{})
	Namespace() string
	Emit(event string, v ...interface{})
	Join(room string)
	Leave(room string)
	LeaveAll()
	Rooms() []string
}

type AppSocket interface {
	Conn
	common.Requester
}

type appSocket struct {
	Conn
	common.Requester
}

func NewAppSocket(conn Conn, req common.Requester) *appSocket {
	return &appSocket{conn, req}
}
