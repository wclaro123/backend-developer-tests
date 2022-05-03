package people

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/wclaro123/stackpath/backend-developer-tests/rest-service/pkg/logging"
)

type Endpoints struct {
	GetPeople     endpoint.Endpoint
	GetPersonByID endpoint.Endpoint
	FindPerson    endpoint.Endpoint
}

func MakeEndpoints(s Service, l log.Logger) Endpoints {
	return Endpoints{
		GetPeople:     makeGetPeople(s, l),
		GetPersonByID: makeGetPersonByID(s, l),
		FindPerson:    makeFindPerson(s, l),
	}
}

func makeGetPeople(s Service, l log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		people, err := s.GetAll(ctx)
		if err != nil {
			_ = l.Log("error", logging.HandleError(err))
			return nil, err
		}

		return GetPeopleResponse{
			Data:  people,
			Total: len(people),
		}, err
	}
}

func makeGetPersonByID(s Service, l log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPersonByIDRequest)

		person, err := s.GetByID(ctx, req.ID)
		if err != nil {
			_ = l.Log("error", logging.HandleError(err))
			return nil, err
		}

		return person, nil
	}
}

func makeFindPerson(s Service, l log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPeopleRequest)

		people, err := s.Find(ctx, PersonFilter{
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			PhoneNumber: req.Phone,
		})

		if err != nil {
			_ = l.Log("error", logging.HandleError(err))
			return nil, err
		}

		return GetPeopleResponse{
			Data:  people,
			Total: len(people),
		}, err
	}
}
