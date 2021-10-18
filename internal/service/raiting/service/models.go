package service

import (
	"fmt"
	"syscall"
)

var allEntities = []Service{
	{service_id: 1, value: 5},
	{service_id: 2, value: 5},
	{service_id: 1, value: 3},
	{service_id: 2, value: 2},
}

type Service struct {
	service_id    int
	value         int
	updated_ts    syscall.Timespec
	reviews_count int
}

func (s Service) String() string {
	return fmt.Sprintf("Service service_id: %d, value: %d, updated: %t, reviews count: %d", s.service_id, s.value, s.updated_ts, s.reviews_count)
}
