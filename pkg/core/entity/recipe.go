package entity

type (
	// Recipe -
	Recipe struct {
		Title       string   `json:"title"`
		Ingredients []string `json:"ingredients"`
		Link        string   `json:"link"`
		GIF         string   `json:"gif"`
	}

	//Recipes -
	Recipes struct {
		KeyWords []string `json:"keywords"`
		Recipes  []Recipe `json:"recipes"`
	}
)
