package router

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var w = websocket.Upgrader{
	// 允许所有请求跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


