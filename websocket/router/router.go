package router

import (
	"github.com/gorilla/mux"
	"go/websocket/handler/ws"
)

/*
 * @author taohu
 * @date 2020/3/18
 * @time 上午10:18
 * @desc please describe function
**/

func RegisterRoutes(r *mux.Router) {
	wsRouter := r.PathPrefix("/ws").Subrouter()
	wsRouter.HandleFunc("/echo", ws.EchoMessage)
	wsRouter.HandleFunc("/echo_display", ws.DisplayEcho)
}
