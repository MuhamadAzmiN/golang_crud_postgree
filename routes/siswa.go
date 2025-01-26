package routes

import (
	"myapp/handler"
	"myapp/middleware"

	"github.com/labstack/echo/v4"
)


func InitSiswaRoutes(e *echo.Echo) {
    api := e.Group("/api")
    api.Use(middleware.AuthMiddleware)
    api.GET("/siswa", handler.GetAllSiswa)
    api.POST("/store", handler.CreateSiswa)
    api.DELETE("/siswa/:id", handler.DeleteSiswa)
    api.GET("/siswa/:id", handler.DetailSiswa)
    api.PUT("/siswa/:id", handler.UpdateSiswa)
    
}

