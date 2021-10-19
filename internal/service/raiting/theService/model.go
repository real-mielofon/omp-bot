package theService

import (
	"fmt"
	"time"
)

type TheService struct {
	ServiceId    int
	Value        int
	UpdatedTs    time.Time
	ReviewsCount int
}

const (
	DateLayout = "2 January 2006 15:04"
)

func (t TheService) String() string {
	return fmt.Sprintf("ServiceRating ServiceId: %d, Value: %d, updated: %s, reviews count: %d",
		t.ServiceId, t.Value, t.UpdatedTs.Format(DateLayout), t.ReviewsCount)
}
