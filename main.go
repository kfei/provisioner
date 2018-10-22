package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
)

var addr = flag.String("addr", "localhost:8081", "service address in host:port format")

var available int64

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wsConnPool = sync.Map{}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade:", err)
		return
	}
	defer c.Close()

	wsConnPool.Store(c, c)
	defer func() {
		wsConnPool.Delete(c)
	}()

	for {
		var request = &Message{}
		if err = c.ReadJSON(&request); err != nil {
			log.Println("Read:", err)
			break
		}
		if request.Command == "request-coupon" {
			atomic.AddInt64(&available, -request.RequestCoupon.Count)
		}

		log.Printf("Recv: %+v", request)

		err = c.WriteJSON(&Message{
			Command: "update-stock",
			AvailableCoupon: &messageAvailableCoupon{
				Count: available,
			},
		})
		if err != nil {
			log.Println("Write:", err)
			break
		}
	}
}

func broadcast() {
	wsConnPool.Range(func(key, value interface{}) bool {
		conn, ok := value.(*websocket.Conn)
		if ok {
			conn.WriteJSON(&Message{
				Command: "update-stock",
				AvailableCoupon: &messageAvailableCoupon{
					Count: available,
				},
			})
		}
		return true
	})
}

func runTicker(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			atomic.AddInt64(&available, 10)
			broadcast()
		}
	}
}

func main() {
	flag.Parse()

	// Serve static frontend files directly
	var webRoot = filepath.Join("web", "dist")
	http.Handle("/", http.FileServer(http.Dir(webRoot)))

	http.HandleFunc("/ws", ws)

	var eg, ctx = errgroup.WithContext(context.Background())
	eg.Go(func() error {
		return runTicker(ctx)
	})
	eg.Go(func() error {
		return http.ListenAndServe(*addr, nil)
	})

	log.Fatal(eg.Wait())
}
