package entity

type (
	// PuppyRecipe -
	PuppyRecipe struct {
		Results []struct {
			Title       string `json:"title"`
			Href        string `json:"href"`
			Ingredients string `json:"ingredients"`
		} `json:"results"`
	}
)
