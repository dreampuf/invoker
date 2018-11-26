package web

import (
	"bufio"
	"context"
	"github.com/dreampuf/invoker/service/log"
	"github.com/dreampuf/invoker/service/terminal"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func WebSocket(c *Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.WithError(err).Error("websocket trying upgrade failure")
		c.Error(err)
		return
	}

	/* handle websocket */
	chgen := make(chan []byte)
	ptmx, cancel := terminal.Connect()
	go strgen(chgen, ptmx)
	go websocketRead(conn, ptmx, cancel)
	go websocketWrite(conn, chgen, cancel)
}

func strgen(ch chan []byte, r io.ReadCloser) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		//log.Print(scanner.Text())
		ch <- scanner.Bytes()
	}
	r.Close()
	close(ch)
}

func websocketWrite(conn *websocket.Conn, send <-chan []byte, cancelFunc context.CancelFunc) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		conn.Close()
	}()
	for {
		select {
		case message, ok := <-send:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				conn.WriteMessage(websocket.CloseMessage, []byte{})
				log.Info("close ws write handle")
				cancelFunc()
				return
			}

			w, err := conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(send)
			for i := 0; i < n; i++ {
				//w.Write(newline)
				w.Write(<-send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func websocketRead(conn *websocket.Conn, w io.WriteCloser, cancelFunc context.CancelFunc) {
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.WithError(err).Error("reading from websocket failure")
			}
			log.WithError(err).Error("exit ws read handle")
			break
		}
		w.Write(message)
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		log.WithField("message", message).Info("news")
	}
	cancelFunc()
	w.Close()
}
