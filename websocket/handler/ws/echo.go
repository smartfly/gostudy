package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

/*
 * @author taohu
 * @date 2020/3/18
 * @time 上午10:01
 * @desc please describe function
**/

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func EchoMessage(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		// 读取客户端的消息
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		// 把消息打印到标准输出
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// 把消息写回客户端，完成回音
		if err = conn.WriteMessage(msgType, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
