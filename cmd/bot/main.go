package main

import (
	"context"
	"github.com/ozonmp/rtg-service-api/pkg/rtg-service-api"
	rtg_service_facade "github.com/ozonmp/rtg-service-facade/pkg/rtg-service-facade"
	"github.com/real-mielofon/omp-bot/internal/pkg/logger"
	"github.com/real-mielofon/omp-bot/internal/service/raiting"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	routerPkg "github.com/real-mielofon/omp-bot/internal/app/router"
)

const (
	rtgServiceAddr = "127.0.0.1:8082"
	rtgFacadeAddr  = "127.0.0.1:8085"
	timeout        = 2 * time.Minute // Timeout
)

func main() {

	ctx := context.Background()

	rtgServiceConn, err := grpc.DialContext(
		ctx,
		rtgServiceAddr,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(),
	)
	if err != nil {
		logger.ErrorKV(ctx, "grpc.DialContext Service", "err", err)
		return
	}

	rtgFacadeConn, err := grpc.DialContext(
		ctx,
		rtgFacadeAddr,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor(),
	)
	if err != nil {
		logger.ErrorKV(ctx, "grpc.DialContext Facade", "err", err)
		return
	}

	rtgServiceAPIClient := rtg_service_api.NewRtgServiceApiServiceClient(rtgServiceConn)
	rtgServiceFacadeClient := rtg_service_facade.NewRtgServiceFacadeServiceClient(rtgFacadeConn)
	grpcClient := raiting.NewClient(rtgServiceAPIClient, rtgServiceFacadeClient)

	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		logger.FatalKV(ctx, "environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		logger.ErrorKV(ctx, "NewBotAPI", "err", err)
		return
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logger.ErrorKV(ctx, "bot.GetUpdatesChan", "err", err)
		return
	}

	routerHandler := routerPkg.NewRouter(bot, grpcClient, timeout)

	for update := range updates {
		routerHandler.HandleUpdate(update)
	}
}
