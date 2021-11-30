package main

import (
	"context"
	"github.com/ozonmp/rtg-service-api/pkg/rtg-service-api"
	rtg_service_facade "github.com/ozonmp/rtg-service-facade/pkg/rtg-service-facade"
	"github.com/real-mielofon/omp-bot/internal/config"
	"github.com/real-mielofon/omp-bot/internal/pkg/logger"
	"github.com/real-mielofon/omp-bot/internal/service/raiting"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	routerPkg "github.com/real-mielofon/omp-bot/internal/app/router"
	"os"
)

func main() {

	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.FatalKV(ctx, "Failed init configuration", "err", err)
	}
	cfg := config.GetConfigInstance()

	logger.SetLogger(logger.CloneWithLevel(ctx, zapcore.DebugLevel))

	var err error
	var rtgServiceConn *grpc.ClientConn
	for i := 0; i < cfg.Retry.Count; i++ {
		logger.DebugKV(ctx, "rtg-service-api connect...", "i", i)
		ctxDial, cancel := context.WithTimeout(ctx, cfg.Timeout)
		defer cancel()
		rtgServiceConn, err = grpc.DialContext(
			ctxDial,
			cfg.RtgService.Address,
			grpc.WithInsecure(),
			grpc.WithChainUnaryInterceptor(),
			grpc.WithBlock(),
		)
		if err == nil {
			break
		}
		time.Sleep(cfg.Retry.Delay)
	}

	if err != nil {
		logger.ErrorKV(ctx, "grpc.DialContext rtg-service-api", "err", err)
		return
	}
	logger.InfoKV(ctx, "rtg-service-api connected", "err", err)

	var rtgFacadeConn *grpc.ClientConn
	for i := 0; i < cfg.Retry.Count; i++ {
		logger.DebugKV(ctx, "rtg-service-facade connect...", "i", i)
		ctxDial, cancel := context.WithTimeout(ctx, cfg.Timeout)
		defer cancel()
		rtgFacadeConn, err = grpc.DialContext(
			ctxDial,
			cfg.RtgFacade.Address,
			grpc.WithInsecure(),
			grpc.WithChainUnaryInterceptor(),
			grpc.WithBlock(),
		)
		if err == nil {
			break
		}
		time.Sleep(cfg.Retry.Delay)
	}
	if err != nil {
		logger.ErrorKV(ctx, "grpc.DialContext rtg-service-facade", "err", err)
		return
	}
	logger.InfoKV(ctx, "rtg-service-facade connected", "err", err)

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

	logger.InfoKV(ctx, "Authorized on account", "bot username", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		logger.ErrorKV(ctx, "bot.GetUpdatesChan", "err", err)
		return
	}

	routerHandler := routerPkg.NewRouter(bot, grpcClient, cfg.Timeout)

	for update := range updates {
		routerHandler.HandleUpdate(ctx, update)
	}
}
