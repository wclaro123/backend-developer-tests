package shared

import (
	"context"
	"encoding/json"
	"net/http"
)

func AddJsonHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	AddJsonHeader(w)
	if msError := GetMSError(err); msError != nil {
		w.WriteHeader(CodesMap[msError.Code()])
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	AddJsonHeader(w)
	return json.NewEncoder(w).Encode(response)
}
