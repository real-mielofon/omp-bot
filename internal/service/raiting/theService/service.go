package theService

import (
	"fmt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Rating {
	return allEntities
}

func (s *Service) Get(idx int) (*Rating, error) {
	if idx < 0 || idx >= len(allEntities) {
		return nil, fmt.Errorf("Range check error idx: %d ", idx)
	}
	return &allEntities[idx], nil
}

func (s *Service) Delete(idx int) error {
	if idx < 0 || idx >= len(allEntities) {
		return fmt.Errorf("Range check error idx: %d ", idx)
	}
	allEntities = append(allEntities[0:idx], allEntities[idx+1:]...)
	return nil
}

func (s *Service) New() (*Rating, error) {
	allEntities = append(allEntities[:], Rating{})
	idx := len(allEntities) - 1

	return &allEntities[idx], nil
}
