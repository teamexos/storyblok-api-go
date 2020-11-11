package storyblok

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

// StoryInput defines the valid input parameters for GetStory
type StoryInput struct {
	FindBy           string `url:"find_by,omitempty"`
	Version          string `url:"version,omitempty"`
	ResolveLinks     string `url:"resolve_links,omitempty"`
	ResolveRelations string `url:"resolve_relations,omitempty"`
	FromRelease      int    `url:"from_release,omitempty"`
	CV               string `url:"cv,omitempty"`
	Language         string `url:"language,omitempty"`
	FallbackLang     string `url:"fallback_lang,omitempty"`
}

// QueryParams returns a GetStoryInput struct as a query param string
func (i StoryInput) QueryParams() (url.Values, error) {
	v, err := query.Values(i)
	if err != nil {
		return nil, err
	}

	return v, nil
}
