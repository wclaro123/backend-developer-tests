package people

import (
	"context"
	"errors"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
)

const (
	getAllMethod = "GetAll"
)

type ServiceTestSuite struct {
	suite.Suite
	repo      *MockRepository
	underTest Service
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.repo = &MockRepository{}
	suite.underTest = NewService(suite.repo)
}

// this is just an example on how to add test coverage
func (suite *ServiceTestSuite) TestGetAllNoError() {
	person := PersonDTO{
		ID:          uuid.Must(uuid.FromString("81eb745b-3aae-400b-959f-748fcafafd81")),
		FirstName:   "John",
		LastName:    "Doe",
		PhoneNumber: "123456",
	}
	resultDTO := []PersonDTO{person}
	result := []Person{fromDto(person)}

	ctx := context.Background()
	suite.repo.Mock.On(getAllMethod, ctx).Return(resultDTO, nil)
	people, err := suite.underTest.GetAll(ctx)

	suite.NoError(err)
	suite.EqualValues(result, people)

}

func (suite *ServiceTestSuite) TestGetAllError() {
	ctx := context.Background()
	errNew := errors.New("error")
	suite.repo.Mock.On(getAllMethod, ctx).Return(nil, errNew)

	people, err := suite.underTest.GetAll(ctx)
	suite.Error(err)
	suite.Nil(people)
}
