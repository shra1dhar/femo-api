package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ding dong")
	})
}