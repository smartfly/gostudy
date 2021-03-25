package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"go/webhook/controller"
)

/*
 * @author taohu
 * @date 2021/2/24
 * @time 下午3:23
 * @desc webhook分发应用
**/

// port - default port to start application on
const port = ":8090"

type WebhookRequest struct {
	Name        string
	Destination string
}

func main() {
	destinations := make(map[string]string)
	mu := &sync.Mutex{}
	dispatcher := controller.NewDispatcher(&http.Client{}, destinations, mu)
	// preparing HTTP server
	srv := &http.Server{Addr: port, Handler: http.DefaultServeMux}
	// webhook registration handler
	http.HandleFunc("/webhooks", func(resp http.ResponseWriter, req *http.Request) {
		dec := json.NewDecoder(req.Body)
		var wr WebhookRequest
		err := dec.Decode(&wr)
		if err != nil {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}
		dispatcher.Add(wr.Name, wr.Destination)
	})
	// start dispatching webhooks
	go dispatcher.Start()
	fmt.Printf("Create webhooks on http://localhost%s/webhooks \n", port)
	// starting server
	err := srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
