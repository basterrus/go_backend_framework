package authorization

import (
	"encoding/json"
	"github.com/basterrus/go_backend_framework/internal"
	"github.com/basterrus/go_backend_framework/internal/apperror"
	"github.com/basterrus/go_backend_framework/internal/domain/user"
	"github.com/basterrus/go_backend_framework/internal/rest"
	"github.com/basterrus/go_backend_framework/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler struct {
	Logger      logging.Logger
	AuthService Service
	UserService user.Service
}

func (uh *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, internal.SignIn, rest.ErrorMiddleware(uh.SignIn))
	router.HandlerFunc(http.MethodPost, internal.SignUp, rest.ErrorMiddleware(uh.SignUp))
}

func (uh *Handler) SignIn(w http.ResponseWriter, r *http.Request) error {

	var si SignIn

	uh.Logger.Debug("[Auth Handler] call sign_in handler")
	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)

	uh.Logger.Debug("[Auth Handler] decode request data")
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&si); err != nil {
		return apperror.BadRequestError("invalid JSON scheme. check swagger API")
	}
	if si.Email == "" || si.Password == "" {
		return apperror.BadRequestError("invalid JSON scheme. field email or password is not provide")
	}
	exists := uh.AuthService.GetUserByEmail(r.Context(), si.Email)

	_ = exists

	return nil
}

func (uh *Handler) SignUp(w http.ResponseWriter, r *http.Request) error {
	var cu user.CreateUserDTO

	uh.Logger.Debug("[Auth Handler] call sign_in handler")
	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)

	uh.Logger.Debug("[Auth Handler] decode request data")
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&cu); err != nil {
		return apperror.BadRequestError("invalid JSON scheme. check swagger API")
	}
	token, err := uh.AuthService.GenerateToken(cu.Email, cu.Password)
	if err != nil {
		//todo add error handler
	}

	data, err := json.Marshal(token)
	if err != nil {
		//todo add error handler
	}
	w.Write(data)
	w.WriteHeader(http.StatusCreated)
	return nil
}
