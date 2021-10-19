package theService

import "github.com/real-mielofon/omp-bot/internal/model/raiting"

type ServiceService interface {
	Describe(serviceID uint64) (*raiting.TheService, error)
	List(cursor uint64, limit uint64) ([]raiting.TheService, error)
	Create(raiting.TheService) (uint64, error)
	Update(serviceID uint64, service raiting.TheService) error
	Remove(serviceID uint64) (bool, error)
}

type DummyServiceService struct{}

func (d DummyServiceService) Describe(serviceID uint64) (*raiting.TheService, error) {
	return nil, nil
}

func (d DummyServiceService) List(cursor uint64, limit uint64) ([]raiting.TheService, error) {
	return []raiting.TheService{}, nil
}

func (d DummyServiceService) Create(service raiting.TheService) (uint64, error) {
	return 0, nil
}

func (d DummyServiceService) Update(serviceID uint64, service raiting.TheService) error {
	return nil
}

func (d DummyServiceService) Remove(serviceID uint64) (bool, error) {
	return false, nil
}

func NewDummyServiceService() *DummyServiceService {
	return &DummyServiceService{}
}

// ...
