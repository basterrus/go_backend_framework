package rest

import (
	"encoding/json"
	"github.com/basterrus/go_backend_framework/internal"
	"net/http"
)

func EncodeJSONResponse(w http.ResponseWriter, r *http.Request, resp interface{}) {
	jEnc := json.NewEncoder(w)
	if _, ok := r.URL.Query()["pretty"]; ok {
		jEnc.SetIndent("", "  ")
	}
	w.Header().Set(internal.MimeTypeContentType, internal.MimeTypeApplicationJSON)
	jEnc.Encode(resp)
}
