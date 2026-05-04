package main

import (
	"context"
	"log"

	"github.com/Tranduy1dol/kotoba-press-core/api"
	"github.com/Tranduy1dol/kotoba-press-core/config"
	"github.com/Tranduy1dol/kotoba-press-core/internal/adapter/mongo"
	"github.com/Tranduy1dol/kotoba-press-core/internal/auth"
	"github.com/Tranduy1dol/kotoba-press-core/internal/usecase"
)

// @title           Learning Japanese API
// @version         1.0
// @description     A Japanese learning application API

// @host            learning-japanese.onrender.com
// @BasePath        /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter "Bearer {token}"
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config")
	}

	_, db, err := mongo.NewClient(context.Background(), cfg.MongoDB.URI, cfg.MongoDB.Database)
	if err != nil {
		log.Fatalf("failed to connect mongodb: %v", err)
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

	lookupSvc := usecase.NewLookupService(wordRepo, grammarRepo)
	testGenSvc := usecase.NewTestGeneratorService(questionRepo, paragraphRepo, testRepo)
	userSvc := usecase.NewUserService(userRepo)
	adminSvc := usecase.NewAdminService(wordRepo, questionRepo, paragraphRepo, grammarRepo)
	srsSvc := usecase.NewSRSService(srsRepo, wordRepo)

	router := api.SetupRouter(cfg.Server.EnableSwagger, cfg.Server.UIBaseURL, googleOAuthService, jwtSvc, lookupSvc, testGenSvc, userSvc, adminSvc, srsSvc)

	log.Printf("server starting on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
