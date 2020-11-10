package storyblok

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

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

	// GetStoryInput defines the valid input parameters for GetStory
	GetStoryInput struct {
		FindBy           string `url:"find_by,omitempty"`
		Version          string `url:"version,omitempty"`
		ResolveLinks     string `url:"find_by,omitempty"`
		ResolveRelations string `url:"resolve_relations,omitempty"`
		FromRelease      int    `url:"from_release,omitempty"`
		CV               string `url:"cv,omitempty"`
		Language         string `url:"language,omitempty"`
		FallbackLang     string `url:"fallback_lang,omitempty"`
	}

	// StoryResponse represents the structured reponse from Storyblok
	// for a single story
	// https://www.storyblok.com/docs/api/content-delivery#core-resources/stories/the-story-object
	StoryResponse struct {
		Story Story `json:"story"`
	}
)

// QueryParams returns a GetStoryInput struct as a query param string
func (i GetStoryInput) QueryParams() (url.Values, error) {
	v, err := query.Values(i)
	if err != nil {
		return nil, err
	}

	return v, nil
}
