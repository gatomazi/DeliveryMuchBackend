package entity

type (
	// Giphy -
	Giphy struct {
		Data []struct {
			URL string `json:"url"`
		} `json:"data"`
	}
)
