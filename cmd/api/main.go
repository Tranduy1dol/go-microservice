package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Tranduy1dol/kotoba-press-core/api"
	"github.com/Tranduy1dol/kotoba-press-core/config"
	"github.com/Tranduy1dol/kotoba-press-core/internal/adapter/grpc"
	"github.com/Tranduy1dol/kotoba-press-core/internal/adapter/mongo"
	"github.com/Tranduy1dol/kotoba-press-core/internal/auth"
	"github.com/Tranduy1dol/kotoba-press-core/internal/logger"
	"github.com/Tranduy1dol/kotoba-press-core/internal/usecase"
)

// @title           Learning Japanese API
// @version         1.0
// @description     A Japanese learning application API

// @BasePath        /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter "Bearer {token}"
func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger.Init(cfg.Server.Env, cfg.Server.LogLevel)

	_, db, err := mongo.NewClient(context.Background(), cfg.MongoDB.URI, cfg.MongoDB.Database)
	if err != nil {
		slog.Error("failed to connect mongodb", "error", err)
		os.Exit(1)
	}

	var searchClient *grpc.SearchClient
	if cfg.GRPC.SearchEngineAddr != "" {
		sc, err := grpc.NewSearchClient(cfg.GRPC.SearchEngineAddr)
		if err != nil {
			slog.Warn("could not connect to search engine", "error", err)
		} else {
			searchClient = sc
			defer func() { _ = searchClient.Close() }()
		}
	}

	wordRepo := mongo.NewWordRepository(db)
	userRepo := mongo.NewUserRepository(db)
	questionRepo := mongo.NewQuestionRepository(db)
	grammarRepo := mongo.NewGrammarRepository(db)
	paragraphRepo := mongo.NewParagraphRepository(db)
	testRepo := mongo.NewTestRepository(db)
	srsRepo := mongo.NewSRSRepository(db)

	jwtSvc := auth.NewJWTService(cfg.OAuth.JWTSecret)
	googleOAuthService := auth.NewGoogleOAuthService(cfg.OAuth, jwtSvc, userRepo)

	lookupSvc := usecase.NewLookupService(wordRepo, grammarRepo, searchClient)
	testGenSvc := usecase.NewTestGeneratorService(questionRepo, paragraphRepo, testRepo)
	userSvc := usecase.NewUserService(userRepo)
	adminSvc := usecase.NewAdminService(wordRepo, questionRepo, paragraphRepo, grammarRepo, searchClient)
	srsSvc := usecase.NewSRSService(srsRepo, wordRepo)

	router := api.SetupRouter(cfg.Server.EnableSwagger, cfg.Server.UIBaseURL, googleOAuthService, jwtSvc, lookupSvc, testGenSvc, userSvc, adminSvc, srsSvc)

	slog.Info("server starting", "port", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		slog.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
