package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"text/template"

	config "github.com/KoLLlaka/bongocat/internal/config/json"
	"github.com/gorilla/websocket"
)

type Position struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type Message struct {
	TypeOfMsg string `json:"type_of_msg"`
	Position  `json:"position,omitempty"`
	KeyBoard  string `json:"key_board,omitempty"`
}

var tmpl *template.Template
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}

func init() {
	tmpl = template.Must(template.ParseGlob("static/templates/*.html"))
}

type Server struct {
	mux          *http.ServeMux
	cfg          config.Config
	clients      map[string]*websocket.Conn
	posChan      chan Position
	keyBoardChan chan string
}

func New(mux *http.ServeMux, cfg config.Config, posChan chan Position, keyBoardChan chan string) *Server {
	return &Server{
		mux:          mux,
		cfg:          cfg,
		clients:      make(map[string]*websocket.Conn),
		posChan:      posChan,
		keyBoardChan: keyBoardChan,
	}
}

func (server *Server) Start(path string) {
	server.mux.HandleFunc(path+"", server.bongocat)
	server.mux.HandleFunc(path+"/ws", server.bongocatws)
}

func (server *Server) bongocat(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "bongocat.html", server.cfg.Keyboard)
}

func (server *Server) bongocatws(w http.ResponseWriter, r *http.Request) {
	var posBefore Position
	conn, _ := upgrader.Upgrade(w, r, nil)
	server.clients["bongocat"] = conn
	defer delete(server.clients, "bongocat")
	defer conn.Close()

	go func() {
		for {
			mt, message, err := conn.ReadMessage()

			if err != nil || mt == websocket.CloseMessage {
				break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
			}
			go server.handleMessage(message)

			//time.Sleep(time.Second * 10)
		}
	}()

	for {
		msg := Message{}
		select {
		case pos := <-server.posChan:
			if math.Abs(pos.X-posBefore.X) > 1 || math.Abs(pos.Y-posBefore.Y) > 1 {
				posBefore = pos
				afterPos := server.pointer(pos)
				msg = Message{
					TypeOfMsg: "mouse",
					Position:  afterPos,
				}
			}
		case key := <-server.keyBoardChan:
			msg = Message{
				TypeOfMsg: "keyboard",
				KeyBoard:  key,
			}
		}

		if msg == (Message{}) {
			continue
		}

		server.writeMessage(msg)
	}
}

func (server *Server) writeMessage(message Message) {

	conn := server.clients["bongocat"]

	var network bytes.Buffer
	enc := json.NewEncoder(&network)
	err := enc.Encode(message)
	if err != nil {
		log.Println(err)
		return
	}

	conn.WriteMessage(websocket.TextMessage, network.Bytes())
}

func (server *Server) handleMessage(message []byte) {
	fmt.Printf("[chat] %s", message)
}

func (server *Server) pointer(pos Position) Position {
	var afterPos Position

	if pos.X > server.cfg.Decstop.X {
		pos.X -= server.cfg.Decstop.X
	}

	afterPos.X = 100/server.cfg.Decstop.X*pos.X - 50/server.cfg.Decstop.Y*pos.Y + 80
	afterPos.Y = 30/server.cfg.Decstop.X*pos.X + 30/server.cfg.Decstop.Y*pos.Y + 180

	return afterPos
}
