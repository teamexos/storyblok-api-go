package storyblok_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamexos/storyblok-api-go"
	"github.com/teamexos/storyblok-api-go/testdata/fixtures"
	"github.com/teamexos/storyblok-api-go/testdata/mocks"
)

var (
	ctx              context.Context
	httpClient       *mocks.MockHTTPClient
	storyblockClient *storyblok.Client
)

func init() {
	ctx = context.Background()
	httpClient = &mocks.MockHTTPClient{}
	storyblockClient = storyblok.NewClient(httpClient, "access_token")
}

func TestGetStory(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       fixtures.ResponseGetStory(),
		}, nil
	}

	story, err := storyblockClient.GetStory(ctx, "/fake/slug", nil)

	assert.Nil(t, err)
	assert.NotNil(t, story.Story)
	assert.Equal(t, "Perseverance", story.Story.Name)
	assert.Equal(t, 21231306, story.Story.ID)
}

func TestStoryNotFound(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       fixtures.ResponseStoryNotFound(),
		}, nil
	}

	_, err := storyblockClient.GetStory(ctx, "/fake/slug", nil)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, err.Message, `{"error":["This record could not be found"]}`)
}

func TestUnauthorized(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       fixtures.ResponseUnauthorized(),
		}, nil
	}

	_, err := storyblockClient.GetStory(ctx, "/fake/slug", nil)

	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusUnauthorized)
	assert.EqualValues(t, err.Message, `{"error":"Unauthorized"}`)
}
