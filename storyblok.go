package storyblok

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURLv1 = "https://api.storyblok.com/v1/cdn"

// Client is the client struct used to access the Storyblok APIs
type Client struct {
	baseURL      string
	token        string
	SpaceVersion int
	HTTPClient   HTTPClient
}

// NewClient returns a pointer to Client
func NewClient(token string) *Client {
	client := &Client{
		baseURL:      baseURLv1,
		token:        token,
		HTTPClient:   DefaultHTTPClient(),
	}
	// Get the latest space version for all subsequent story calls
    space, err := client.GetLatestSpace(context.Background())
    if err == nil {
		client.SpaceVersion = space.Version
	}
	return client
}

// GetLatestSpace gets the latest space version so that we can always deliver the latest content
//
// https://www.storyblok.com/docs/api/content-delivery#topics/cache-invalidation
func (c *Client) GetLatestSpace(ctx context.Context) (*Space, *ResponseError) {
	endpoint := fmt.Sprintf("%s/spaces/me", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, NewResponseError(http.StatusInternalServerError, errCodeRequestSetupFailed)
	}

	q := req.URL.Query()
	q.Add("token", c.token)
	req.URL.RawQuery = q.Encode()

	res := SpaceResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res.Space, nil
}

// GetStory returns a story object for the full_slug, id or uuid if
// authenticated using a preview or public token.
//
// https://www.storyblok.com/docs/api/content-delivery#core-resources/stories/retrieve-one-story
func (c *Client) GetStory(ctx context.Context,
	id string,
	input *StoryInput) (*StoryResponse, *ResponseError) {

	endpoint := fmt.Sprintf("%s/stories/%s", c.baseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, NewResponseError(http.StatusInternalServerError, errCodeRequestSetupFailed)
	}

	if input != nil {
		q, err := input.QueryParams()
		if err != nil {
			return nil, NewResponseError(http.StatusUnprocessableEntity, err.Error())
		}
		req.URL.RawQuery = q.Encode()
	}

	q := req.URL.Query()
	q.Add("token", c.token)
	if c.SpaceVersion != 0 {
		q.Add("cv", string(c.SpaceVersion))
	}
	req.URL.RawQuery = q.Encode()

	res := StoryResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) *ResponseError {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return NewResponseError(http.StatusInternalServerError, errCodeRequestDoFailed)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		err := buf.String()

		return NewResponseError(res.StatusCode, err)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return NewResponseError(http.StatusInternalServerError,
			fmt.Sprintf("%s: %s", errCodeRequestDecodeFailed, err.Error()))
	}

	return nil
}
