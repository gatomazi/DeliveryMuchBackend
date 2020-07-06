package api

import (
	"deliverymuch/api/handlers/recipes"
	"deliverymuch/pkg/core/recipe"
	"deliverymuch/pkg/middleware"
	"deliverymuch/pkg/router"
	"os"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

// Start -
func Start() {

	root := "api"
	rout := router.Setup()
	rout.Use(middleware.Cors())

	public := rout.Group(root)
	// Prepare Recipe Service
	recipeService := recipe.NewService()
	recipes.EnableHandlers(public, *recipeService)

	port := os.Getenv("API_PORT")
	rout.Run(port)
}
