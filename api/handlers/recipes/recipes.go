package recipes

import (
	"deliverymuch/pkg/core/recipe"
	"deliverymuch/pkg/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

var service recipe.Service

var messages = map[string]string{
	"unexpected": "Não foi possível fazer esta operação no momento.",
}

func getRecipes(ctx *router.Context) (int, *router.Response) {
	ingredients := ctx.Queries["i"].(string)
	u, errCode, err := service.ReadAll(ingredients)
	if err != nil {
		return http.StatusNotFound, router.NewResposeError(err.Error(), errCode)
	}

	return http.StatusOK, router.NewResponseSuccess(u)
}

// EnableHandlers -
func EnableHandlers(public *gin.RouterGroup, serv recipe.Service) {
	service = serv
	mapEnd := []*router.EndPoint{

		{
			Name:    "recipes",
			Method:  "GET",
			Handler: getRecipes,
			Group:   public,
		},
	}

	for _, point := range mapEnd {
		router.EnableHandlers(point)
	}
}
