package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/websocket"
)

const retStr = "<!doctype html>\r\n<html>\r\n  <head>\r\n    <meta charset=\"utf-8\">\r\n    <title></title>\r\n  </head>\r\n  <body>\r\n    <p>it works!</p>\r\n  </body>\r\n</html>\r\n"

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, retStr)
	}
	// h2 := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello world!\n")
	// }
	h3 := func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)
		defer conn.Close()

		if err != nil {
			log.Panicf("upgrader.Upgrade: %v", err)
			return
		}

		for {
			messageType, p, err := conn.ReadMessage()
			fmt.Println(string(p))

			if err != nil {
				log.Printf("conn.ReadMessage: %v", err)
				return
			}
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Printf("conn.WriteMessage: %v", err)
				return
			}
		}
	}
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/endpoint", h1)
	http.HandleFunc("/ws", h3)

	port := "8080"
	log.Fatal(http.ListenAndServe(":"+port, nil))
	fmt.Println("server is running")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
