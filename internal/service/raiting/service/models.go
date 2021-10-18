package service

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	dateLayout = "2 January 2006 15:04"
)

var allEntities []Service

type Service struct {
	ServiceId    int
	Value        int
	UpdatedTs    time.Time
	ReviewsCount int
}

func (s Service) String() string {
	return fmt.Sprintf("Service ServiceId: %d, Value: %d, updated: %s, reviews count: %d",
		s.ServiceId, s.Value, s.UpdatedTs.Format(dateLayout), s.ReviewsCount)
}

func init() {
	allEntities = append(allEntities, Service{ServiceId: 1, Value: 5, UpdatedTs: timeToStr("2 January 2012 15:04")})
	allEntities = append(allEntities, Service{ServiceId: 2, Value: 4, UpdatedTs: timeToStr("3 January 2012 15:04")})

	for i := 0; i < 30; i++ {
		s := Service{
			ServiceId: rand.Intn(10) + 1,
			Value:     rand.Intn(5) + 1,
			UpdatedTs: time.Date(
				2017+rand.Intn(5), time.Month(rand.Intn(12)+1), rand.Intn(29)+1,
				rand.Intn(24), rand.Intn(60), rand.Intn(60), 0, time.UTC),
		}
		allEntities = append(allEntities, s)
	}
	allEntities = append(allEntities, Service{ServiceId: 10, Value: 3})
	allEntities = append(allEntities, Service{ServiceId: 11, Value: 2})
	allEntities = append(allEntities, Service{ServiceId: 12, Value: 1})
}

func timeToStr(s string) (t time.Time) {
	t, _ = time.Parse(dateLayout, s)
	return t
}
