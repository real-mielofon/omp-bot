package theService

import "github.com/real-mielofon/omp-bot/internal/model/raiting"

type ServiceService interface {
	Describe(serviceID uint64) (*raiting.Service, error)
	List(cursor uint64, limit uint64) ([]raiting.Service, error)
	Create(raiting.Service) (uint64, error)
	Update(serviceID uint64, service raiting.Service) error
	Remove(serviceID uint64) (bool, error)
}

type DummyServiceService struct{}

func (d DummyServiceService) Describe(serviceID uint64) (*raiting.Service, error) {
	return nil, nil
}

func (d DummyServiceService) List(cursor uint64, limit uint64) ([]raiting.Service, error) {
	return []raiting.Service{}, nil
}

func (d DummyServiceService) Create(service raiting.Service) (uint64, error) {
	return 0, nil
}

func (d DummyServiceService) Update(serviceID uint64, service raiting.Service) error {
	return nil
}

func (d DummyServiceService) Remove(serviceID uint64) (bool, error) {
	return false, nil
}

func NewDummyServiceService() *DummyServiceService {
	return &DummyServiceService{}
}

// ...
