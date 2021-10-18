package service

import (
	"fmt"
)

type ServiceService struct{}

func NewService() *ServiceService {
	return &ServiceService{}
}

func (s *ServiceService) List() []Service {
	return allEntities
}

func (s *ServiceService) Get(idx int) (*Service, error) {
	if idx < 0 || idx >= len(allEntities) {
		return nil, fmt.Errorf("Range check error idx: %d ", idx)
	}
	return &allEntities[idx], nil
}

func (s *ServiceService) Delete(idx int) error {
	if idx < 0 || idx >= len(allEntities) {
		return fmt.Errorf("Range check error idx: %d ", idx)
	}
	allEntities = append(allEntities[0:idx], allEntities[idx+1:]...)
	return nil
}

func (s *ServiceService) New() (*Service, error) {
	allEntities = append(allEntities[:], Service{})
	idx := len(allEntities) - 1

	return &allEntities[idx], nil
}
