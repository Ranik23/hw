package models

type APIResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Locale string `json:"locale"`
	Seed   string `json:"seed"`
	Total  int    `json:"total"`
	Data   []struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Author      string `json:"author"`
		Genre       string `json:"genre"`
		Description string `json:"description"`
		ISBN        string `json:"isbn"`
		Image       string `json:"image"`
		Published   string `json:"published"`
		Publisher   string `json:"publisher"`
	} `json:"data"`
}
