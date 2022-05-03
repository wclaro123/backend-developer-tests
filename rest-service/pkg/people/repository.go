package people

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"github.com/wclaro123/stackpath/backend-developer-tests/rest-service/pkg/database"
	"github.com/wclaro123/stackpath/backend-developer-tests/rest-service/pkg/shared"
)

type Repository interface {
	GetAll(ctx context.Context) ([]PersonDTO, error)
	GetByID(ctx context.Context, id uuid.UUID) (PersonDTO, error)
	Find(ctx context.Context, firstName, lastName, phone string) ([]PersonDTO, error)
}

type repository struct {
	db database.Database
}

func NewRepository(db database.Database) Repository {
	return &repository{db: db}
}

func (r repository) GetAll(_ context.Context) ([]PersonDTO, error) {
	ppl := r.db.AllPeople()
	result := make([]PersonDTO, len(ppl))
	for i, person := range ppl {
		result[i] = fromDatabase(person)
	}

	return result, nil
}

func (r repository) GetByID(_ context.Context, id uuid.UUID) (PersonDTO, error) {
	person, err := r.db.FindPersonByID(id)
	if err != nil {
		return PersonDTO{}, shared.NewDBError(shared.ErrDatabase.Error(), err)
	}

	return fromDatabase(person), nil
}

func (r repository) Find(_ context.Context, firstName, lastName, phone string) ([]PersonDTO, error) {
	var searchArr []*database.Person

	if firstName != "" || lastName != "" {
		searchArr = r.db.FindPeopleByName(firstName, lastName)
	}

	if phone != "" {
		searchArr = append(searchArr, r.db.FindPeopleByPhoneNumber(phone)...)
	}

	var result []PersonDTO
	resultMap := make(map[uuid.UUID]bool)
	for _, person := range searchArr {
		if resultMap[person.ID] {
			continue
		}

		resultMap[person.ID] = true
		result = append(result, fromDatabase(person))
	}

	return result, nil
}
