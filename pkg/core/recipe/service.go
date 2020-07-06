package recipe

import (
	"deliverymuch/pkg/core/entity"
	"deliverymuch/pkg/utils"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/subosito/gotenv"
)

type (
	// Service service interface
	Service struct {
		PuppyURL string
		GiphyURL string
		GiphyKEY string
		Messages map[string]string
	}
)

//NewService -
func NewService() *Service {
	gotenv.Load()
	return &Service{
		PuppyURL: os.Getenv("PUPPY_URL"),
		GiphyURL: os.Getenv("GIPHY_URL"),
		GiphyKEY: os.Getenv("GIPHY_KEY"),
		Messages: map[string]string{
			"not-found":           "Receita não encontrada",
			"unexpected":          "Algo inesperado aconteceu.",
			"unexpected-request":  "Algo inesperado aconteceu na requisição.",
			"unexpected-read-all": "Não foi possível acessar este recurso",
		},
	}
}

//ReadAll -
func (s *Service) ReadAll(ingredients string) (data entity.Recipes, errCode string, err error) {
	var giphy entity.Giphy

	data.KeyWords = s.SortIngredients(ingredients)
	puppyRecipes, err := s.GetPuppyRecipes(data.KeyWords)
	if err != nil {
		errCode = s.Messages["unexpected-request"]
		return
	}

	data.Recipes = make([]entity.Recipe, len(puppyRecipes.Results))
	for index, puppyRecipe := range puppyRecipes.Results {
		data.Recipes[index].Title = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(puppyRecipe.Title, "\r", ""), "\t", ""), "\n", "")
		data.Recipes[index].Ingredients = s.SortIngredients(puppyRecipe.Ingredients)
		data.Recipes[index].Link = puppyRecipe.Href
		giphy, err = s.GetGiphy(data.Recipes[index].Title)
		if err != nil {
			errCode = s.Messages["unexpected-request"]
			return
		}
		if len(giphy.Data) > 0 {
			data.Recipes[index].GIF = giphy.Data[0].URL
		}
	}

	return
}

//SortIngredients -
func (s *Service) SortIngredients(ingredients string) []string {
	explodeIngredients := strings.Split(ingredients, ",")

	fmt.Println(explodeIngredients)
	for i := range explodeIngredients {
		explodeIngredients[i] = strings.TrimSpace(explodeIngredients[i])
	}
	sort.Strings(explodeIngredients)
	return explodeIngredients
}

//GetGiphy -
func (s *Service) GetGiphy(title string) (structGiphy entity.Giphy, err error) {

	giphyTarget := fmt.Sprintf("%s/gifs/search?api_key=%s&q=%s&limit=1", s.GiphyURL, s.GiphyKEY, title)

	giphy, err := utils.RequestGenerator(giphyTarget, "GET", nil)
	if err != nil {
		return
	}
	mapstructure.Decode(giphy, &structGiphy)
	return
}

//GetPuppyRecipes -
func (s *Service) GetPuppyRecipes(ingredients []string) (structPuppy entity.PuppyRecipe, err error) {
	puppyTarget := fmt.Sprintf("%s/?i=%s", s.PuppyURL, strings.Join(ingredients, ","))
	puppyRecipes, err := utils.RequestGenerator(puppyTarget, "GET", nil)
	if err != nil {
		return
	}
	mapstructure.Decode(puppyRecipes, &structPuppy)
	return
}
