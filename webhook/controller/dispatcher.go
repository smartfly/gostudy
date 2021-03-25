package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
 * @author taohu
 * @date 2021/2/24
 * @time 下午3:27
 * @desc please describe function
**/

type Dispatcher struct {
	client       *http.Client
	destinations map[string]string
	mu           *sync.Mutex
}

func NewDispatcher(client *http.Client, destinations map[string]string, mu *sync.Mutex) *Dispatcher {
	return &Dispatcher{client: client, destinations: destinations, mu: mu}
}

func (d *Dispatcher) Add(name, destination string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.destinations[name] = destination
}

func (d *Dispatcher) sendHttpPost(user, destination string) {
	req, err := http.NewRequest("POST", destination,
		bytes.NewBufferString(fmt.Sprintf("Hello %s, current time is %s", user, time.Now().String())))
	if err != nil {
		// probably don't allow creating invalid destinations
		return
	}
	resp, err := d.client.Do(req)
	if err != nil {
		// should probably check response status code and retry if it's timeout or 500
		return
	}
	fmt.Printf("Webhook to '%s' dispatched, response code: %d \n", destination, resp.StatusCode)
}

func (d *Dispatcher) dispatch() {
	d.mu.Lock()
	defer d.mu.Unlock()
	for user, destination := range d.destinations {
		go func(user, destination string) {
			d.sendHttpPost(user, destination)
		}(user, destination)
	}
}

func (d *Dispatcher) Start() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			d.dispatch()
		}
	}
}
