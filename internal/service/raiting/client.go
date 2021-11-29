package raiting

import (
	"context"
	"github.com/real-mielofon/omp-bot/internal/model/raiting"
	"github.com/real-mielofon/omp-bot/internal/pkg/logger"

	rtg_service_api "github.com/ozonmp/rtg-service-api/pkg/rtg-service-api"
	rtg_service_facade "github.com/ozonmp/rtg-service-facade/pkg/rtg-service-facade"
)

type raitingSertviceClient struct {
	grpcClient rtg_service_api.RtgServiceApiServiceClient
	grpcFacade rtg_service_facade.RtgServiceFacadeServiceClient
}

func (r raitingSertviceClient) Describe(ctx context.Context, serviceID uint64) (*raiting.TheService, error) {
	resp, err := r.grpcClient.DescribeService(ctx, &rtg_service_api.DescribeServiceRequest{
		Id: serviceID,
	})
	if err != nil {
		logger.ErrorKV(ctx, "grpcClient.DescribeService", "err", err)
		return nil, err
	}
	return &raiting.TheService{
		ID:           resp.Service.Id,
		Value:        int(resp.Service.Value),
		UpdatedTs:    resp.Service.UpdatedTs.AsTime(),
		ReviewsCount: int(resp.Service.ReviewsCount),
	}, nil
}

func (r raitingSertviceClient) Create(ctx context.Context, service raiting.TheService) (*raiting.TheService, error) {
	resp, err := r.grpcClient.CreateService(ctx, &rtg_service_api.CreateServiceRequest{
		Value:        uint64(service.Value),
		ReviewsCount: uint64(service.ReviewsCount),
	})
	if err != nil {
		logger.ErrorKV(ctx, "grpcClient.CreateService", "err", err)
		return nil, err
	}
	return &raiting.TheService{
		ID:           resp.Service.Id,
		Value:        int(resp.Service.Value),
		UpdatedTs:    resp.Service.UpdatedTs.AsTime(),
		ReviewsCount: int(resp.Service.ReviewsCount),
	}, nil
}

func (r raitingSertviceClient) Update(ctx context.Context, serviceID uint64, service raiting.TheService) error {
	_, err := r.grpcClient.UpdateService(ctx, &rtg_service_api.UpdateServiceRequest{
		Id: serviceID,
		Service: &rtg_service_api.Service{
			Value:        uint64(service.Value),
			ReviewsCount: uint64(service.ReviewsCount),
		},
	})
	if err != nil {
		logger.ErrorKV(ctx, "grpcClient.UpdateService", "err", err)
		return err
	}
	return nil
}

func (r raitingSertviceClient) Remove(ctx context.Context, serviceID uint64) (bool, error) {
	resp, err := r.grpcClient.RemoveService(ctx, &rtg_service_api.RemoveServiceRequest{
		Id: serviceID,
	})
	if err != nil {
		logger.ErrorKV(ctx, "grpcClient.RemoveService", "err", err)
		return false, err
	}
	return resp.Found, nil
}

func (r raitingSertviceClient) List(ctx context.Context, cursor uint64, limit uint64) ([]raiting.TheService, error) {
	resp, err := r.grpcFacade.ListServices(ctx, &rtg_service_facade.ListServicesRequest{
		Limit:  limit,
		Cursor: cursor,
	})
	if err != nil {
		logger.ErrorKV(ctx, "grpcFacade.ListServices", "err", err)
		return nil, err
	}

	services := make([]raiting.TheService, len(resp.Services))
	for i, s := range resp.Services {
		services[i] = convertPBService(s)
	}

	return services, nil
}

func convertPBService(s *rtg_service_facade.Service) raiting.TheService {
	return raiting.TheService{
		s.Id,
		int(s.Value),
		s.UpdatedTs.AsTime(),
		int(s.ReviewsCount),
	}
}

// NewClient for GPRS raiting service
func NewClient(grpcClient rtg_service_api.RtgServiceApiServiceClient, grpcFacade rtg_service_facade.RtgServiceFacadeServiceClient) *raitingSertviceClient {
	return &raitingSertviceClient{
		grpcClient: grpcClient,
		grpcFacade: grpcFacade,
	}
}
