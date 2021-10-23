package theService

import (
	"fmt"
	"math/rand"
	"time"
)

type ServiceService interface {
	Describe(serviceID uint64) (*TheService, error)
	List(cursor uint64, limit uint64) ([]TheService, error)
	Create(TheService) (uint64, error)
	Update(serviceID uint64, service TheService) error
	Remove(serviceID uint64) (bool, error)
}

type DummyServiceService struct {
	allEntities []TheService
}

func NewDummyTheServiceService() *DummyServiceService {
	dummyService := DummyServiceService{}
	dummyService.init()
	return &dummyService
}

func (s *DummyServiceService) List(cursor uint64, limit uint64) ([]TheService, error) {
	if int(cursor) >= len(s.allEntities) {
		return nil, fmt.Errorf("Range check error idx: %d ", cursor)
	}
	last := int(cursor + limit)
	if last > len(s.allEntities) {
		last = len(s.allEntities)
	}
	return s.allEntities[cursor:last], nil
}

func (s *DummyServiceService) Describe(serviceID uint64) (*TheService, error) {
	if int(serviceID) > len(s.allEntities)-1 {
		return nil, fmt.Errorf("Range check error idx: %d ", serviceID)
	}
	return &s.allEntities[serviceID], nil
}

func (s *DummyServiceService) Update(serviceID uint64, service TheService) error {
	if int(serviceID) >= len(s.allEntities) {
		return fmt.Errorf("Range check error idx: %d ", serviceID)
	}
	s.allEntities[serviceID] = service
	return nil
}

func (s *DummyServiceService) Remove(serviceID uint64) (bool, error) {
	if int(serviceID) >= len(s.allEntities) {
		return false, fmt.Errorf("Range check error idx: %d ", serviceID)
	}
	s.allEntities = append(s.allEntities[0:serviceID], s.allEntities[serviceID+1:]...)
	return true, nil
}

func (s *DummyServiceService) Create(r TheService) (uint64, error) {
	s.allEntities = append(s.allEntities[:], r)
	idx := len(s.allEntities) - 1

	return uint64(idx), nil
}

func (s *DummyServiceService) init() {
	_, _ = s.Create(TheService{ServiceId: 1, Value: 5, UpdatedTs: strToTime("2 January 2012 15:04")})
	_, _ = s.Create(TheService{ServiceId: 1, Value: 5, UpdatedTs: strToTime("2 January 2012 15:04")})
	_, _ = s.Create(TheService{ServiceId: 2, Value: 4, UpdatedTs: strToTime("3 January 2012 15:04")})

	for i := 0; i < 30; i++ {
		r := TheService{
			ServiceId: rand.Intn(10) + 1,
			Value:     rand.Intn(5) + 1,
			UpdatedTs: time.Date(
				2017+rand.Intn(5), time.Month(rand.Intn(12)+1), rand.Intn(29)+1,
				rand.Intn(24), rand.Intn(60), rand.Intn(60), 0, time.UTC),
		}
		_, _ = s.Create(r)
	}
}

func strToTime(s string) (t time.Time) {
	t, _ = time.Parse(DateLayout, s)
	return t
}
