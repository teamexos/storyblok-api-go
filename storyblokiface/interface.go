package storyblokiface

import (
	"context"

	"github.com/teamexos/storyblok-api-go"
)

// StoryblokClient is an interface for the Storyblok Client
type StoryblokClient interface {
	GetStory(ctx context.Context,
		id string,
		input *storyblok.StoryInput) (*storyblok.StoryResponse, *storyblok.ResponseError)
}

var _ StoryblokClient = (*storyblok.Client)(nil)
