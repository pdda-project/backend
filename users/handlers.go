// users/handlers.go

package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pdda-project/backend/libs/responses"
	"github.com/pdda-project/backend/libs/utils"
	"gorm.io/gorm"
)

type UserHandler struct {
	s *UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	s := NewUserService(db)
	return &UserHandler{s: s}
}

func (h *UserHandler) RegisterRouter(r *http.ServeMux) {
	r.HandleFunc("POST /users", h.register)

}

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
