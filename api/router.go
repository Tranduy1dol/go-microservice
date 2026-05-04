package api

import (
	"github.com/Tranduy1dol/learning-japanese/api/handler"
	"github.com/Tranduy1dol/learning-japanese/api/middleware"
	_ "github.com/Tranduy1dol/learning-japanese/docs"
	"github.com/Tranduy1dol/learning-japanese/internal/auth"
	"github.com/Tranduy1dol/learning-japanese/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(
	enableSwagger bool,
	authSvc *auth.GoogleOAuthService,
	jwtSvc *auth.JWTService,
	lookupSvc *usecase.LookupService,
	testGenSvc *usecase.TestGeneratorService,
	userSvc *usecase.UserService,
	adminSvc *usecase.AdminService,
	srsSvc *usecase.SRSService,
) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	if enableSwagger {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

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
				return
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
		userHandler := handler.NewUserHandler(userSvc)
		srsHandler := handler.NewSRSHandler(srsSvc)

		v1.GET("/words/:id", wordHandler.GetWord)
		v1.GET("/words/search", wordHandler.SearchWords)
		v1.GET("/words/jlpt/:level", wordHandler.BrowseWordsByJLPT)

		v1.GET("/grammar/:id", grammarHandler.GetGrammar)
		v1.GET("/grammar", grammarHandler.ListGrammar)

		v1.POST("/tests/generate/:level", testHandler.GenerateTest)
		v1.POST("/tests/:id/submit", testHandler.SubmitTest)

		v1.GET("/users/me", userHandler.GetMe)

		v1.POST("/srs/deck", srsHandler.AddWordToDeck)
		v1.GET("/srs/due", srsHandler.GetDueCards)
		v1.POST("/srs/review/:id", srsHandler.ReviewCard)

		admin := v1.Group("/admin")
		admin.Use(middleware.AdminMiddleware())
		{
			adminHandler := handler.NewAdminHandler(adminSvc)
			admin.POST("/words", adminHandler.CreateWord)
			admin.DELETE("/words/:id", adminHandler.DeleteWord)
			admin.POST("/questions", adminHandler.CreateQuestion)
			admin.DELETE("/questions/:id", adminHandler.DeleteQuestion)
			admin.POST("/paragraphs", adminHandler.CreateParagraph)
			admin.DELETE("/paragraphs/:id", adminHandler.DeleteParagraph)
			admin.POST("/grammars", adminHandler.CreateGrammar)
			admin.DELETE("/grammars/:id", adminHandler.DeleteGrammar)
		}
	}

	return r
}
