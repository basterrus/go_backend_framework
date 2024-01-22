package roles

import (
	"github.com/basterrus/go_backend_framework/internal"
	"github.com/basterrus/go_backend_framework/internal/rest"
	"github.com/basterrus/go_backend_framework/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler struct {
	Logger      logging.Logger
	UserService Service
}

func (uh *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, internal.RolesURL, rest.ErrorMiddleware(uh.CreateRoleHandler))
	router.HandlerFunc(http.MethodGet, internal.RoleURL, rest.ErrorMiddleware(uh.GetRoleHandler))
	router.HandlerFunc(http.MethodGet, internal.RolesURL, rest.ErrorMiddleware(uh.GetRoleListHandler))
	router.HandlerFunc(http.MethodPatch, internal.RoleURL, rest.ErrorMiddleware(uh.UpdateRoleHandler))
	router.HandlerFunc(http.MethodDelete, internal.RoleURL, rest.ErrorMiddleware(uh.DeleteRoleHandler))
}

func (uh *Handler) CreateRoleHandler(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (uh *Handler) GetRoleHandler(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (uh *Handler) GetRoleListHandler(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (uh *Handler) UpdateRoleHandler(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func (uh *Handler) DeleteRoleHandler(w http.ResponseWriter, r *http.Request) error {

	return nil
}

//func (uh *Handler) CreateRole(w http.ResponseWriter, r *http.Request) error {
//
//	var cu CreateUserDTO
//
//	uh.Logger.Debug("call create user handler")
//	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)
//
//	uh.Logger.Debug("decode create user dto")
//	defer r.Body.Close()
//
//	if err := json.NewDecoder(r.Body).Decode(&cu); err != nil {
//		return apperror.BadRequestError("invalid JSON scheme. check swagger API")
//	}
//	userUUID, err := uh.UserService.Create(r.Context(), cu)
//	if err != nil {
//		return err
//	}
//	w.Header().Set("Location", fmt.Sprintf("%s/%s", internal.UsersURL, userUUID))
//	w.WriteHeader(http.StatusCreated)
//
//	return nil
//
//}
//
//func (uh *Handler) GetUser(w http.ResponseWriter, r *http.Request) error {
//
//	uh.Logger.Debug("call get user handler")
//	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)
//
//	uh.Logger.Info("get uuid user from query params")
//	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
//	userUuid := params.ByName("uuid")
//
//	user, err := uh.UserService.GetByUUID(r.Context(), userUuid)
//	if err != nil {
//		return err
//	}
//	userBytes, err := json.Marshal(user)
//	if err != nil {
//		return err
//	}
//	w.WriteHeader(http.StatusOK)
//	w.Write(userBytes)
//	return nil
//
//}
//
//func (uh *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
//
//	return nil
//}
//
//func (uh *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) error {
//
//	uh.Logger.Debug("call delete handler")
//	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)
//
//	uh.Logger.Info("get uuid user from query params")
//	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
//
//	userUuid := params.ByName("uuid")
//
//	err := uh.UserService.Delete(r.Context(), userUuid)
//	if err != nil {
//		return err
//	}
//	w.WriteHeader(http.StatusNoContent)
//	return nil
//}
