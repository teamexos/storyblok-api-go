# storyblok-api-go

[Storyblock](https://www.storyblok.com/docs/api/content-delivery) Go client

## Install

```
go get -u github.com/teamexos/storyblok-api-go
```

## Usage

```go
package main

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	storyblok "github.com/teamexos/storyblok-api-go"
)

func main() {
	storyblokClient := storyblok.NewClient("TOKEN")
	ctx := context.Background()

	input := &storyblok.StoryInput{
        // Put query params here
        // See story_input.go for fields
	}
	story, err := storyblokClient.GetStory(ctx, "/content/slug", input)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	spew.Dump(story.Story.Content)
}

```

## Testing

```go
go test -v -tags=unit
```