package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-tbc/internal/config"
)

func StartServer(handlers *config.Handlers) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:5173", "http://103.26.13.166:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12,
	}))

	public := router.Group("api")
	{
		public.POST("/verify-user", handlers.AuthHandler.VerifyUser)
		public.GET("/health-centers", handlers.FetchHandler.GetHealthCenters)
		public.GET("/files/health-centers/:filename", handlers.FetchHandler.GetHealthCenterImage)
	}

	admin := router.Group("api").Use(IsValidJWT()).Use(IsRole("ADMIN")).Use(SetUserUuid())
	{
		admin.GET("/users/detail", handlers.FetchHandler.GetUserByUuid)
		admin.GET("/health-centers/:uuid", handlers.FetchHandler.GetHealthCenterByUuid)
		admin.POST("/health-centers", handlers.ManagementHandler.CreateHealthCenter)
		admin.PUT("/health-centers/:uuid", handlers.ManagementHandler.UpdateHealthCenter)

		admin.POST("/cases", handlers.ManagementHandler.CreateCase)
		admin.PUT("/cases/:uuid", handlers.ManagementHandler.UpdateCase)
		admin.DELETE("/cases/:uuid", handlers.ManagementHandler.DeleteCase)
		admin.GET("/cases", handlers.FetchHandler.GetCases)
		admin.GET("/cases/:uuid", handlers.FetchHandler.GetCaseByUuid)

		admin.GET("/years/cases", handlers.FetchHandler.GetCaseYears)
		admin.POST("/clustering", handlers.ProcessingHandler.KmeansClustering)
		admin.GET("/results/years/:year", handlers.FetchHandler.GetResultByYear)
	}

	return router
}
