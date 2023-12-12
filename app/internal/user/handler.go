package user

import (
	"encoding/json"
	"fmt"
	"github.com/basterrus/go_backend_framework/internal/rest"
	"github.com/basterrus/go_backend_framework/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersURL = "/api/users"
	userURL  = "/api/users/:uuid"
)

type Handler struct {
	Logger      logging.Logger
	UserService Service
}

func NewUserHandler() rest.Handler {
	return &Handler{}
}

func (uh *Handler) Register(router *httprouter.Router) {
	//commonHandler := alice.New()
	//router.GET(usersURL, commonHandler.Append(rest.AcceptHandler(rest.MimeTypeApplicationJSON)).ThenFunc(uh.GetUser))
	router.HandlerFunc(http.MethodGet, usersURL, rest.ErrorMiddleware(uh.GetUserByEmailAndPassword))
	router.HandlerFunc(http.MethodPost, usersURL, uh.CreateUser)
	router.HandlerFunc(http.MethodGet, userURL, uh.GetUser)
	router.HandlerFunc(http.MethodPatch, userURL, uh.PartiallyUpdateUser)
	router.HandlerFunc(http.MethodDelete, userURL, uh.DeleteUser)
}

func (uh *Handler) GetUserByEmailAndPassword(w http.ResponseWriter, r *http.Request) error {
	uh.Logger.Info("GET USER")
	w.Header().Set("Content-Type", "application/json")

	uh.Logger.Debug("get uuid from context")
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	userUUID := params.ByName("uuid")

	user, err := uh.UserService.GetOne(r.Context(), userUUID)
	if err != nil {
		return err
	}

	uh.Logger.Debug("marshal user")
	userBytes, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshall user. error: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userBytes)
	return nil
}

func (uh *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (uh *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (uh *Handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (uh *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
