package routes

import (
	"apidal/controllers"
	"apidal/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	dalam := r.Group("/api")
	
	dalam.Get("/", middlewares.AuthMiddleware, controllers.Index)
	dalam.Get("/:id", middlewares.AuthMiddleware, controllers.Show)
	dalam.Post("/", middlewares.AuthMiddleware, controllers.Create)
	dalam.Put("/:id", middlewares.AuthMiddleware, controllers.Update)
	dalam.Delete("/:id", middlewares.AuthMiddleware, controllers.Delete)
}