package people

import (
	"context"

	"github.com/satori/go.uuid"

	"github.com/wclaro123/stackpath/backend-developer-tests/rest-service/pkg/shared"
)

type Service interface {
	GetAll(ctx context.Context) ([]Person, error)
	GetByID(ctx context.Context, id uuid.UUID) (Person, error)
	Find(ctx context.Context, filter PersonFilter) ([]Person, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) GetAll(ctx context.Context) ([]Person, error) {
	dtoArr, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, shared.HandleDbError(err, shared.ErrGetAll.Error(), shared.GetAllOp)
	}

	return fromDtoArr(dtoArr), nil
}

func (s service) GetByID(ctx context.Context, id uuid.UUID) (Person, error) {
	person, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return Person{}, shared.HandleDbError(err, shared.ErrPersonNotFound.Error(), shared.GetByIDOp)
	}

	return fromDto(person), nil
}

func (s service) Find(ctx context.Context, filter PersonFilter) ([]Person, error) {
	dtoArr, err := s.repo.Find(ctx, filter.FirstName, filter.LastName, filter.PhoneNumber)
	if err != nil {
		return nil, shared.HandleDbError(err, shared.ErrFindPerson.Error(), shared.FindOp)
	}

	return fromDtoArr(dtoArr), nil
}
