package main

import (
	"context"
	"log"

	"github.com/Tranduy1dol/learning-japanese/api"
	"github.com/Tranduy1dol/learning-japanese/config"
	"github.com/Tranduy1dol/learning-japanese/internal/adapter/mongo"
	"github.com/Tranduy1dol/learning-japanese/internal/auth"
	"github.com/Tranduy1dol/learning-japanese/internal/usecase"
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
		log.Fatalf("failed to connect mongodb")
	}

	wordRepo := mongo.NewWordRepository(db)
	userRepo := mongo.NewUserRepository(db)
	questionRepo := mongo.NewQuestionRepository(db)
	grammarRepo := mongo.NewGrammarRepository(db)
	paragraphRepo := mongo.NewParagraphRepository(db)

	jwtSvc := auth.NewJWTService(cfg.OAuth.JWTSecret)
	googleOAuthService := auth.NewGoogleOAuthService(cfg.OAuth, jwtSvc, userRepo)

	lookupSvc := usecase.NewLookupService(wordRepo, grammarRepo)
	testGenSvc := usecase.NewTestGeneratorService(questionRepo, paragraphRepo)
	userSvc := usecase.NewUserService(userRepo)
	adminSvc := usecase.NewAdminService(wordRepo, questionRepo, paragraphRepo, grammarRepo)

	router := api.SetupRouter(cfg.Server.EnableSwagger, googleOAuthService, jwtSvc, lookupSvc, testGenSvc, userSvc, adminSvc)

	log.Printf("server starting on port %s", cfg.Server.Port)
	if err := router.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
