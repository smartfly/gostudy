package main

import (
	"context"
	"github.com/gorilla/mux"
	"go/websocket/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

/*
 * @author taohu
 * @date 2020/3/18
 * @time 上午10:19
 * @desc please describe function
**/

func main() {

	muxRouter := mux.NewRouter()
	router.RegisterRoutes(muxRouter)

	server := &http.Server{
		Addr:    ":8080",
		Handler: muxRouter,
	}

	// 创建系统信号接收器
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	log.Println("Starting HTTP server...")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}
