package person

import (
	"clean-architecture/domain/person"
	"context"

	"github.com/stretchr/testify/mock"
)

func NewModelMock() *PersonModelMock {
	return &PersonModelMock{}
}

type PersonModelMock struct {
	mock.Mock
}

func (m *PersonModelMock) GetPersonByID(ctx context.Context, personID int) (*person.Person, error) {
	args := m.Called(ctx, personID)
	return args.Get(0).(*person.Person), args.Error(1)
}
