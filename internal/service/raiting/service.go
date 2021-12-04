package raiting

import (
	"context"
	"github.com/real-mielofon/omp-bot/internal/model/raiting"
)

// TheServiceService is service of raiting.TheService
type TheServiceService interface {
	Describe(ctx context.Context, serviceID uint64) (*raiting.TheService, error)
	List(ctx context.Context, cursor uint64, limit uint64) ([]raiting.TheService, error)
	//	ListAll() ([]raiting.TheService, error)
	Create(ctx context.Context, service raiting.TheService) (*raiting.TheService, error)
	Update(ctx context.Context, serviceID uint64, service raiting.TheService) error
	Remove(ctx context.Context, serviceID uint64) (bool, error)
}
