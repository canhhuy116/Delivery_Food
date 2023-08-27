package skio

import (
	"Delivery_Food/component"
	"Delivery_Food/component/tokenprovider/jwt"
	"Delivery_Food/modules/user/userstorage"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"sync"
)

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, key string, data interface{}) error
	EmitToUser(userId int, key string, data interface{}) error
	Run(ctx component.AppContext, engine gin.Engine) error
}

type realtimeEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewRealtimeEngine() *realtimeEngine {
	return &realtimeEngine{
		storage: make(map[int][]AppSocket),
		locker:  new(sync.RWMutex),
	}
}

func (engine *realtimeEngine) saveAppSocket(userId int, socket AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()

	if _, ok := engine.storage[userId]; !ok {
		engine.storage[userId] = make([]AppSocket, 0)
	}

	engine.storage[userId] = append(engine.storage[userId], socket)
}

func (engine *realtimeEngine) getAppSockets(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()

	if sockets, ok := engine.storage[userId]; ok {
		return sockets
	}

	return nil
}

func (engine *realtimeEngine) removeAppSocket(userId int, socket AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()

	if sockets, ok := engine.storage[userId]; ok {
		for i, s := range sockets {
			if s.ID() == socket.ID() {
				engine.storage[userId] = append(sockets[:i], sockets[i+1:]...)
				return
			}
		}
	}
}

func (engine *realtimeEngine) UserSockets(userId int) []AppSocket {
	return engine.getAppSockets(userId)
}

func (engine *realtimeEngine) EmitToRoom(room string, key string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, key, data)
	return nil
}

func (engine *realtimeEngine) EmitToUser(userId int, key string, data interface{}) error {
	sockets := engine.getAppSockets(userId)
	for _, socket := range sockets {
		socket.Emit(key, data)
	}
	return nil
}

func (engine *realtimeEngine) Run(appCtx component.AppContext,
	r *gin.Engine) error {
	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	if err != nil {
		return err
	}

	engine.server = server

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), "IP:", s.RemoteAddr().String())
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("error:", e)
	})

	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		db := appCtx.GetMainDbConnection()
		store := userstorage.NewSQLStore(db)

		tokenProvider := jwt.NewJwtProvider(appCtx.SecretKey())

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			s.Emit("unauthorized", err.Error())
			s.Close()
			return
		}

		user, err := store.FindUser(context.Background(),
			map[string]interface{}{"id": payload.
				UserId})

		if err != nil {
			s.Emit("unauthorized", err.Error())
			s.Close()
			return
		}

		if user.Status == 0 {
			s.Emit("unauthorized", "user is not active")
			s.Close()
			return
		}

		user.Mask(false)

		appSck := NewAppSocket(s, user)
		engine.saveAppSocket(user.ID, appSck)

		s.Emit("authenticated", user)
	})

	go server.Serve()

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	return nil
}
