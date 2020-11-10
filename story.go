package storyblok

type (
	// Stories represents the structured response from Storyblok
	// for multiple stories
	// https://www.storyblok.com/docs/api/content-delivery#core-resources/stories/the-story-object
	Stories struct {
		Stories []Story `json:"stories"`
	}

	// Story is a struct representation of the main story object
	Story struct {
		ID               int                    `json:"id"`
		UUID             string                 `json:"uuid"`
		Name             string                 `json:"name"`
		Slug             string                 `json:"slug"`
		FullSlug         string                 `json:"full_slug"`
		DefaultFullSlug  string                 `json:"default_full_slug"`
		CreatedAt        string                 `json:"created_at"`
		PublishedAt      string                 `json:"published_at"`
		FirstPublishedAt string                 `json:"first_published_at"`
		ReleaseID        string                 `json:"release_id"`
		Lang             string                 `json:"lang"`
		Content          map[string]interface{} `json:"content"`
		Position         int                    `json:"position"`
		IsStartPage      bool                   `json:"is_startpage"`
		ParentID         int                    `json:"parent_id"`
		GroupID          string                 `json:"group_id"`
		Alternates       []*Alternate           `json:"alternates"`
		TranslatedSlugs  []*TranslatedSlug      `json:"translated_slugs"`
	}

	// StoryResponse represents the structured reponse from Storyblok
	// for a single story
	// https://www.storyblok.com/docs/api/content-delivery#core-resources/stories/the-story-object
	StoryResponse struct {
		Story Story `json:"story"`
	}
)
