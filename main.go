package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/doneland/yquotes"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

var connections map[*websocket.Conn]bool

type TickerInfo struct {
	Ticker    string
	Bid       string
	Ask       string
	Symbol    string
	Open      string
	Last      string
	Date      string
	PrevClose string
}

func sendAll(msg []byte) {
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(connections, conn)
			conn.Close()
		}
	}
}
func sendOne(msg []byte, conn *websocket.Conn) {
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Taken from gorilla's website
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	log.Println("Succesfully upgraded connection")
	connections[conn] = true
	//stopLoop := make(chan string)
	var stopLoop chan string

	for {
		// Blocks until a message is read
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Connection closed!")
			delete(connections, conn)
			conn.Close()
			close(stopLoop)
			return
		}
		log.Println("stopLoop ", stopLoop)
		if stopLoop != nil {
			stopLoop <- ""
		}
		stopLoop = make(chan string)
		log.Println("Stock to follow: ", string(msg))

		sendOne([]byte(string(msg)+" is the stock to follow"), conn)
		go func() {

			for {
				select {
				default:
					timer1 := time.NewTimer(time.Second * 1)
					<-timer1.C
					stock, err := yquotes.NewStock(string(msg), false)
					if err != nil {
						// handle error
					}
					symbol := stock.Symbol // AAPL
					name := stock.Name     // Apple Inc.
					// Price information
					price := stock.Price // Price struct
					bid := price.Bid
					ask := price.Ask
					open := price.Open
					prevClose := price.PreviousClose
					last := price.Last
					date := price.Date
					tickerInfo := &TickerInfo{Ticker: name, Bid: strconv.FormatFloat(bid, 'f', 6, 64) + " (BID)", Ask: strconv.FormatFloat(ask, 'f', 6, 64) + " (ASK)", Symbol: symbol, Open: strconv.FormatFloat(open, 'f', 6, 64) + " (OPEN)", Last: strconv.FormatFloat(last, 'f', 6, 64) + " (LAST)", PrevClose: strconv.FormatFloat(prevClose, 'f', 6, 64) + " (Previous Close)", Date: date.Format(time.RFC3339)}
					tInfo, _ := json.Marshal(tickerInfo)
					sendOne([]byte(tInfo), conn)
					// log.Println("Symbol ", symbol, " -Name ", name)
					// sendOne([]byte("Symbol "+symbol+" -Name "+name), conn)
					// log.Println("Price:Bid ", bid)
					// sendOne([]byte("Price:Bid "+strconv.FormatFloat(bid, 'f', 6, 64)), conn)
					// log.Println("Price:Ask ", ask)
					// sendOne([]byte("Price:Ask "+strconv.FormatFloat(ask, 'f', 6, 64)), conn)

				case <-stopLoop:
					return
				}
			}
		}()
	}
}

func main() {

	// command line flags
	port := flag.Int("port", 8080, "port to serve on")
	dir := flag.String("directory", "client/", "directory of web files")
	flag.Parse()

	connections = make(map[*websocket.Conn]bool)
	// handle all requests by serving a file of the same name
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)
	http.Handle("/", fileHandler)
	http.HandleFunc("/ws", wsHandler)

	log.Printf("Running on port %d\n", *port)
	addr := fmt.Sprintf(":%d", *port)
	// this call blocks -- the progam runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
