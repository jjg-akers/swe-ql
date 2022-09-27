package http

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	server          *http.Server
	router          *mux.Router
	addr string
	port            string
	shutdownTimeout time.Duration
}
