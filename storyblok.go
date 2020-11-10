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
	baseURL    string
	token      string
	httpClient HTTPClient
}

// NewClient returns a pointer to Client
func NewClient(httpClient HTTPClient, token string) *Client {
	return &Client{
		baseURL:    baseURLv1,
		token:      token,
		httpClient: httpClient,
	}
}

// GetStory returns a story object for the full_slug, id or uuid if
// authenticated using a preview or public token.
//
// https://www.storyblok.com/docs/api/content-delivery#core-resources/stories/retrieve-one-story
func (c *Client) GetStory(ctx context.Context,
	id string,
	params ...string) (*StoryResponse, *ResponseError) {

	paramsStr := fmt.Sprintf("?token=%s&", c.token)
	endpoint := fmt.Sprintf("%s/stories/%s%s",
		c.baseURL,
		id,
		paramsStr)

	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, NewResponseError(http.StatusInternalServerError, errCodeRequestSetupFailed)
	}

	res := StoryResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetStories returns a list of stories that are in your Storyblok space.
// The stories are returned in sorted order, depending on the order in
// your space. You can use the query parameter sort_by with any story
// object property and first level of your content type to order the
// response to your needs.
//
// If no entries are found with your filters applied, you will receive
//an empty array. You will not receive a 404 error message, to check
// if you have results go for the array length.
//
// https://www.storyblok.com/docs/api/content-delivery#core-resources/stories/retrieve-multiple-stories
func (c *Client) GetStories(ctx context.Context,
	params ...string) (*Stories, *ResponseError) {

	paramsStr := fmt.Sprintf("?token=%s&", c.token)
	endpoint := fmt.Sprintf("%s/stories%s",
		c.baseURL,
		paramsStr)

	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, NewResponseError(http.StatusInternalServerError, errCodeRequestSetupFailed)
	}

	res := Stories{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) *ResponseError {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.httpClient.Do(req)
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
		return NewResponseError(http.StatusInternalServerError, errCodeRequestDecodeFailed)
	}

	return nil
}
