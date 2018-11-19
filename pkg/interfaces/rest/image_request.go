package rest

type PostProcessImageRequest struct {
	ImageURL    string   `json:"image_url"`
	SaveToCloud bool     `json:"save_to_cloud"`
	Transforms  []string `json "transforms"`
}
