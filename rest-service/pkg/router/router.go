package router

import (
	"net/http"

	kitTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/wclaro123/stackpath/backend-developer-tests/rest-service/pkg/people"
	"github.com/wclaro123/stackpath/backend-developer-tests/rest-service/pkg/shared"
)

// Add your endpoints here
type Endpoints struct {
	People people.Endpoints
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewHandler(endpoints Endpoints, mwf ...mux.MiddlewareFunc) http.Handler {
	r := mux.NewRouter()
	r.Use(mwf...)
	r.StrictSlash(true)
	healthRoute(r)

	options := []kitTransport.ServerOption{
		kitTransport.ServerErrorEncoder(shared.EncodeError),
	}

	people.AddRouteToHandler(endpoints.People, r, options)

	return r
}

func healthRoute(r *mux.Router) {
	r.Path("/health").
		Methods("GET").
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			shared.AddJsonHeader(w)
			_, _ = w.Write([]byte(`"OK"`))
		})
}
