package server

import (
	v1 "github.com/find_property/internal/delivery/v1"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/find_property/docs"
)

func NewRouter(h *v1.Handler) *gin.Engine {

	router := gin.New()

	router.GET("/ping", h.Init)
	router.GET("/subways", h.Subways)
	router.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", h.CheckUserRegistered)
	router.POST("/create", h.RegistrationUser)
	router.POST("/view/add", h.AdViewed)
	router.POST("/user/create", h.CreateUser)
	router.GET("/user/:id", h.GetUser)
	router.POST("/user/subway/lines", h.GetSubwayLines)
	router.POST("/jk/add", h.AddReverseSweepJK)

	router.POST("/search", h.Search)

	router.POST("/subcategory", h.Subcategory)
	router.POST("/price", h.Price)
	router.POST("/area", h.Area)
	router.POST("/jk", h.JK)
	router.POST("/category", h.Category)
	router.POST("/rooms", h.Rooms)
	router.POST("/metro", h.Metro)
	router.POST("/metro/before", h.MetroBefore)
	router.POST("/buildingYear", h.BuildingYear)
	router.POST("/balcony", h.Balcony)
	router.POST("/floors", h.Floors)
	router.POST("/repair", h.RepairType)
	router.POST("/address", h.Address)
	router.POST("/default", h.Default)
	router.POST("/silence", h.Silence)

	return router
}
