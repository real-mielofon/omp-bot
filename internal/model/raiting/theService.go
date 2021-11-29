package raiting

import (
	"fmt"
	"time"
)

// TheService is raiting service struct
type TheService struct {
	ID           uint64
	Value        int
	UpdatedTs    time.Time
	ReviewsCount int
}

const (
	dateLayout = "2 January 2006 15:04"
)

func (t TheService) String() string {
	return fmt.Sprintf("ServiceRating ID: %d, Value: %d, updated: %s, reviews count: %d",
		t.ID, t.Value, t.UpdatedTs.Format(dateLayout), t.ReviewsCount)
}
