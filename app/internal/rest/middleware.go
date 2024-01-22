package rest

import (
	"errors"
	"net/http"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func ErrorMiddleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := h(w, r)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				}
				var err *AppError
				errors.As(err, &err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(err.Marshal())
				return
			}
			w.WriteHeader(http.StatusTeapot)
			w.Write(systemError(err.Error()).Marshal())
		}
	}
}

//func AcceptHandler(cType string) func(http.Handler) http.Handler {
//	m := func(next http.Handler) http.Handler {
//		fn := func(w http.ResponseWriter, r *http.Request) {
//			if r.Header.Get("Accept") != cType {
//				writeError(w, r, newNotAcceptableError(cType))
//				return
//			}
//
//			next.ServeHTTP(w, r)
//		}
//		return http.HandlerFunc(fn)
//	}
//
//	return m
//}
