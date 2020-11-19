package storyblok

type (
	// Space is a struct representation of the space object
	// https://www.storyblok.com/docs/api/content-delivery#core-resources/spaces/the-space-object
	Space struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Domain  string `json:"domain"`
		Version int    `json:"version"`
	}

	// SpaceResponse represents the structured reponse from Storyblok
	// for a single space
	SpaceResponse struct {
		Space Space `json:"space"`
	}
)
