package user

import (
	"encoding/json"
	"fmt"
	"github.com/basterrus/go_backend_framework/internal"
	"github.com/basterrus/go_backend_framework/internal/apperror"
	"github.com/basterrus/go_backend_framework/internal/rest"
	"github.com/basterrus/go_backend_framework/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler struct {
	Logger      logging.Logger
	UserService Service
}

func NewUserHandler() rest.Handler {
	return &Handler{}
}

func (uh *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, internal.UsersURL, rest.ErrorMiddleware(uh.CreateUser))
	router.HandlerFunc(http.MethodGet, internal.UserURL, rest.ErrorMiddleware(uh.GetUser))
	router.HandlerFunc(http.MethodPatch, internal.UserURL, rest.ErrorMiddleware(uh.UpdateUser))
	router.HandlerFunc(http.MethodDelete, internal.UserURL, rest.ErrorMiddleware(uh.DeleteUserHandler))
}

//func (uh *Handler) GetUserByEmailAndPassword(w http.ResponseWriter, r *http.Request) error {
//	uh.Logger.Info("GET USER")
//	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)
//
//	uh.Logger.Debug("get uuid from context")
//	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
//	userUUID := params.ByName("uuid")
//
//	user, err := uh.UserService.GetOne(r.Context(), userUUID)
//	if err != nil {
//		return err
//	}
//
//	uh.Logger.Debug("marshal user")
//	userBytes, err := json.Marshal(user)
//	if err != nil {
//		return fmt.Errorf("failed to marshall user. error: %w", err)
//	}
//
//	w.WriteHeader(http.StatusOK)
//	w.Write(userBytes)
//	return nil
//}

func (uh *Handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	uh.Logger.Debug("call create user handler")
	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)

	uh.Logger.Debug("decode create user dto")
	var cu CreateUserDTO
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&cu); err != nil {
		return apperror.BadRequestError("invalid JSON scheme. check swagger API")
	}
	userUUID, err := uh.UserService.Create(r.Context(), cu)
	if err != nil {
		return err
	}
	w.Header().Set("Location", fmt.Sprintf("%s/%s", internal.UsersURL, userUUID))
	w.WriteHeader(http.StatusCreated)

	return nil

}

func (uh *Handler) GetUser(w http.ResponseWriter, r *http.Request) error {
	uh.Logger.Debug("call get user handler")
	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)
	uh.Logger.Info("get uuid user from query params")
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)

	userUuid := params.ByName("uuid")

	user, err := uh.UserService.GetByUUID(r.Context(), userUuid)
	if err != nil {
		return err
	}
	uh.Logger.Debug(user)
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(userBytes)
	return nil

}

func (uh *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (uh *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) error {

	uh.Logger.Debug("call delete handler")
	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)

	uh.Logger.Info("get uuid user from query params")
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)

	userUuid := params.ByName("uuid")

	err := uh.UserService.Delete(r.Context(), userUuid)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}
