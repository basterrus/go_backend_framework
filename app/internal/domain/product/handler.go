package product

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
	Logger          logging.Logger
	CategoryService Service
}

func (uh *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, internal.CategoriesURL, rest.ErrorMiddleware(uh.CreateCategory))
	//router.HandlerFunc(http.MethodGet, internal.UserURL, rest.ErrorMiddleware(uh.GetUser))
	//router.HandlerFunc(http.MethodPatch, internal.UserURL, rest.ErrorMiddleware(uh.UpdateUser))
	router.HandlerFunc(http.MethodDelete, internal.CategoryURL, rest.ErrorMiddleware(uh.DeleteCategoryHandler))
}

func (uh *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) error {
	var cc CreateCategoryDTO

	uh.Logger.Debug("call create investing handler")
	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)

	uh.Logger.Debug("decode create user dto")
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&cc); err != nil {
		apperror.BadRequestError("invalid JSON scheme. check swagger API")
	}

	return nil
}

func (uh *Handler) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) error {
	uh.Logger.Debug("call create investing handler")
	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)

	uh.Logger.Info("get uuid user from query params")
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	categoryId := params.ByName("uuid")
	err := uh.CategoryService.Delete(r.Context(), categoryId)
	if err != nil {
		apperror.BadRequestError("failed")
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintf("investing %s was deleted", categoryId)))

	return nil
}
