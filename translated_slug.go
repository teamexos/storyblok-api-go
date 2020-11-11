package storyblok

// TranslatedSlug defines the struct storing a translated slug
type TranslatedSlug struct {
	Path string `json:"path"`
	Name string `json:"name"`
	Lang string `json:"lang"`
}
