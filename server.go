package main

import (
	"github.com/gorilla/websocket"
	"learnWebSocket/impl"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true;
		},
	}
)

func wsHandle(w http.ResponseWriter, r *http.Request){
	//w.Write([]byte("hello world"))

	var(
		wsConn *websocket.Conn
		err error
		//msgType int
		data []byte
		conn *impl.Connection
	)
	// upgrader：websocket  http升级wbsocket  得到一个长连接conn
	if wsConn,err = upgrader.Upgrade(w,r,nil);err!=nil{
		return
	}


	// websocket.Conn
	//for{
	//	// Text Binary
	//	if _,data,err = conn.ReadMessage();err!=nil{
	//		goto ERR
	//	}
	//	if err = conn.WriteMessage(websocket.TextMessage,data);err!=nil{
	//		goto ERR
	//	}
	//}

	go func() {
		var (
			err error
		)
		for{
			if err = conn.WriteMessage([]byte("heartbeat"));err!=nil{
				return
			}
			time.Sleep(1*time.Second)
		}
	}()


	if conn,err = impl.InitConnection(wsConn);err!=nil{
		goto ERR
	}

	for{
		if data,err = conn.ReadMessage();err!=nil{
			goto ERR
		}
		if err = conn.WriteMessage(data);err!=nil{
			goto ERR
		}
	}

	ERR :
		conn.Close()
}


func main(){
	http.HandleFunc("/ws",wsHandle)


	http.ListenAndServe("0.0.0.0:8080",nil)
}
