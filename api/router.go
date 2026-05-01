package api

import (
	"github.com/Tranduy1dol/learning-japanese/api/handler"
	"github.com/Tranduy1dol/learning-japanese/api/middleware"
	"github.com/Tranduy1dol/learning-japanese/internal/auth"
	"github.com/Tranduy1dol/learning-japanese/internal/port"
	"github.com/Tranduy1dol/learning-japanese/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	lookupSvc *usecase.LookupService,
	testGenSvc *usecase.TestGeneratorService,
	authSvc *auth.GoogleOAuthService,
	jwtSvc *auth.JWTService,
	userRepo port.UserRepository,
) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	authGroup := r.Group("/auth")
	{
		authGroup.GET("/google", func(ctx *gin.Context) {
			url, _ := authSvc.GetAuthUrl()
			ctx.Redirect(302, url)
		})
		authGroup.GET("/google/callback", func(ctx *gin.Context) {
			code := ctx.Query("code")
			token, user, err := authSvc.HandleCallback(ctx, code)
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
			}
			ctx.JSON(200, gin.H{"token": token, "user": user})
		})
	}

	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware(jwtSvc))
	{
		wordHandler := handler.NewWordHandler(lookupSvc)
		grammarHandler := handler.NewGrammarHandler(lookupSvc)
		testHandler := handler.NewTestHandler(testGenSvc)
		userHandler := handler.NewUserHandler(userRepo)

		v1.GET("/words/:id", wordHandler.GetWord)
		v1.GET("/words/search", wordHandler.SearchWords)
		v1.GET("/words/jlpt/:level", wordHandler.BrowseWordsByJLPT)

		v1.GET("/grammar/:id", grammarHandler.GetGrammar)
		v1.GET("/grammar", grammarHandler.ListGrammar)

		v1.POST("/tests/generate", testHandler.GenerateTest)

		v1.GET("/users/me", userHandler.GetMe)
	}

	return r
}
