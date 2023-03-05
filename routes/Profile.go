package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/middleware"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func ProfileRoutes(e *echo.Group) {
	ProfileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(ProfileRepository)

	e.GET("/profile/:id", middleware.Auth(h.GetProfile))
	e.POST("/profile", middleware.Auth(middleware.UploadFile(h.CreateProfile)))
	e.PATCH("/profile", middleware.Auth(middleware.UploadFile(h.UpdateProfile)))
	e.DELETE("/profile", middleware.Auth(h.DeleteProfile))
}
