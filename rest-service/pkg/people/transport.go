package people

import (
	"context"
	"encoding/json"
	"net/http"

	kitTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	"github.com/wclaro123/stackpath/backend-developer-tests/rest-service/pkg/shared"
)

func AddRouteToHandler(endpoints Endpoints, r *mux.Router, options []kitTransport.ServerOption) {
	r.Methods(http.MethodGet).Path("/people").Handler(
		kitTransport.NewServer(
			endpoints.GetPeople,
			decodeEmptyRequest,
			shared.EncodeResponse,
			options...,
		),
	)

	r.Methods(http.MethodGet).Path("/people/{personID}").Handler(
		kitTransport.NewServer(
			endpoints.GetPersonByID,
			decodeGetPersonByIDRequest,
			shared.EncodeResponse,
			options...,
		),
	)

	r.Methods(http.MethodPost).Path("/people").Handler(
		kitTransport.NewServer(
			endpoints.FindPerson,
			decodeGetPeopleRequest,
			shared.EncodeResponse,
			options...,
		),
	)
}

func decodeEmptyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func decodeGetPeopleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var getPeopleRequest GetPeopleRequest
	if err := json.NewDecoder(r.Body).Decode(&getPeopleRequest); err != nil {
		return nil, shared.NewMSError(shared.ErrJSONInvalid.Error(), shared.BadRequest, shared.DecodeGetPeopleRequest, shared.TransportLevel, err)
	}

	if !((getPeopleRequest.FirstName != "" && getPeopleRequest.LastName != "") || getPeopleRequest.Phone != "") {
		return nil, shared.NewMSError(shared.ErrInvalidRequest.Error(), shared.BadRequest, shared.DecodeGetPeopleRequest, shared.TransportLevel, shared.ErrInvalidRequest)
	}

	return getPeopleRequest, nil
}

func decodeGetPersonByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	str := vars["personID"]
	id, err := uuid.FromString(str)
	if err != nil {
		return nil, shared.NewMSError(shared.ErrInvalidUUID.Error(), shared.BadRequest, shared.DecodeGetPersonByIDRequest, shared.TransportLevel, err)
	}

	req := GetPersonByIDRequest{
		ID: id,
	}
	return req, nil
}
