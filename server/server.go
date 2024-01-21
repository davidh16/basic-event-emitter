package server

import (
	"basic-event-emitter/config"
	"github.com/gorilla/mux"
	gosocketio "github.com/graarh/golang-socketio"
	"log"
	"net/http"
)

type Server struct {
	router *mux.Router
	cfg    *config.Config
	client *gosocketio.Client
}

func NewServer(cfg *config.Config, socketClient *gosocketio.Client) *Server {

	router := mux.NewRouter()

	srv := &Server{
		cfg:    cfg,
		router: router,
		client: socketClient,
	}

	router.HandleFunc("/emit-event", srv.EmitEvent).Methods("POST")

	return srv

}

func (s *Server) Serve() {
	log.Panic(http.ListenAndServe(s.cfg.Port, s.router))
}

func (s *Server) returnResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	} else {
		w.Write([]byte("event emitted successfully"))
		return
	}
}
