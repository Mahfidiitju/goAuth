package router

import (
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger) // Apply the logging middleware
	chiRouter.Use(middlewares.RateLimitMiddleware)
	UserRouter.Register(chiRouter)
	return chiRouter
}
