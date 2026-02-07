package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plinyulan/exit-exam/cmd/api/docs"
	"github.com/plinyulan/exit-exam/internal/conf"
	"github.com/plinyulan/exit-exam/internal/services/controller"
	"github.com/plinyulan/exit-exam/internal/services/repository"
	"github.com/plinyulan/exit-exam/internal/services/usecase"
	"github.com/plinyulan/exit-exam/security"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-User-ID")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *Server) Router() (http.Handler, func()) {
	config := conf.NewConfig()
	r := gin.Default()
	r.Use(CORSMiddleware())

	docs.SwaggerInfo.BasePath = "/api/v2"

	catRepository := repository.NewCatRepository()
	catUsecase := usecase.NewCatUsecase(catRepository)
	catController := controller.NewCatController(catUsecase)

	authRepository := repository.NewAuthRepository(s.db)
	authUsecase := usecase.NewAuthUsecase(authRepository)
	authController := controller.NewAuthController(*config, authUsecase)

	promisRository := repository.NewPromisesRepository(s.db)
	promisesUsecase := usecase.NewPromisesUsecase(promisRository)
	promisesController := controller.NewPromisesController(promisesUsecase)

	campaignReposirory := repository.NewCampaignsRepository(s.db)
	campaignUsecase := usecase.NewCampaignsUsecase(campaignReposirory)
	campaignController := controller.NewCampaignsController(campaignUsecase)

	politiciansRepo := repository.NewPoliticiansRepository(s.db)
	politiciansUsecase := usecase.NewPoliticiansUsecase(politiciansRepo)
	politiciansController := controller.NewPoliticiansController(politiciansUsecase, promisesUsecase)

	api := r.Group("/api/v2")
	{
		authGroup := api.Group("/auth")
		{
			authController.AuthRoutes(authGroup)
		}
		catGroup := api.Group("/cat").Use(security.Middleware())
		{
			catController.CatRoutes(catGroup)
		}
		promisesGroup := api.Group("/promises").Use(security.Middleware())
		{
			promisesController.PromisesRoutes(promisesGroup)
		}
		campaignsGroup := api.Group("/campaigns").Use(security.Middleware())
		{
			campaignController.CampaignsRoutes(campaignsGroup)
		}
		politiciansGroup := api.Group("/politicians").Use(security.Middleware())
		{
			politiciansController.PoliticiansRoutes(politiciansGroup)
		}
	}

	if config.ENV == "dev" || config.ENV == "uat" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	return r, func() {
		// sqlDB, err := s.db.DB()
		// if err != nil {
		// 	panic("Failed to get sql.DB from gorm.DB")
		// }
		// sqlDB.Close()
	}
}
