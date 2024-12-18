// users/handlers.go

package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/pdda-project/backend/libs/responses"
	"github.com/pdda-project/backend/libs/utils"
	"gorm.io/gorm"
)

// Type
type UserHandler struct {
	s *UserService
}

// Constructor
func NewUserHandler(db *gorm.DB) *UserHandler {
	s := NewUserService(db)
	return &UserHandler{s: s}
}

// Add handlers to server router
func (h *UserHandler) RegisterRouter(r *http.ServeMux) {
	r.HandleFunc("POST /users", h.register)
	r.HandleFunc("GET /users", h.list)
	r.HandleFunc("PATCH /users/me", h.update)
}

// Handle register user
func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s from %s", r.Method, r.URL, r.RemoteAddr)
	// Set header
	w.Header().Add("Content-Type", "application/json")

	// Parse body
	var body ReqRegister
	if err := utils.FromRequest(r, &body); err != nil {
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err.Response)
		return
	}

	// Run service
	res, err := h.s.register(&body)

	// Check error on service
	if err != nil {
		log.Printf("HTTPError: %s", err.Error())
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err.Response)
		return
	}

	// Success
	if res != nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(responses.Data{Status: "success", Data: res})
		return
	} else {
		w.WriteHeader(http.StatusNoContent)
		return
	}

}

// Handle list users
func (h *UserHandler) list(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s from %s", r.Method, r.URL, r.RemoteAddr)

	query := r.URL.Query()

	take := 10
	start := 0

	startQuery := query.Get("start")
	intStartQuery, err := strconv.Atoi(startQuery)
	if err == nil {
		start = intStartQuery
		log.Println(start)
	}

	takeQuery := query.Get("take")
	intTakeQuery, err := strconv.Atoi(takeQuery)
	if err == nil {
		take = intTakeQuery
		log.Println(take)
	}

	list, httpError := h.s.list(start, take)
	if httpError != nil {
		w.WriteHeader(httpError.Code)
		json.NewEncoder(w).Encode(httpError.Response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses.Data{Status: "success", Data: list})
}

// Handle update users
func (h *UserHandler) update(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
