package ws

import (
	"net/http"
)

/*
 * @author taohu
 * @date 2020/3/18
 * @time 上午10:07
 * @desc please describe function
**/

func DisplayEcho(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websocket/views/websockets.html")
}
