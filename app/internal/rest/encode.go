package rest

import (
	"encoding/json"
	"net/http"
)

func encodeJSONResponse(w http.ResponseWriter, r *http.Request, resp interface{}) {
	jEnc := json.NewEncoder(w)
	if _, ok := r.URL.Query()["pretty"]; ok {
		jEnc.SetIndent("", "  ")
	}
	w.Header().Set("Content-Type", MimeTypeApplicationJSON)
	jEnc.Encode(resp)
}
