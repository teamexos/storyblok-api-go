# storyblok-api-go

[Storyblock](https://www.storyblok.com/docs/api/content-delivery) Go client

## Install

```
go get -u github.com/teamexos/storyblok-api-go
```

## Usage

```go
import "github.com/teamexos/storyblok-api-go"

// Replace TOKEN with your real API token

httpClient := storyblok.DefaultHTTPClient()
storyblokClient := storyblok.NewClient(httpClient, "TOKEN")
```

## Testing

```go
go test -v -tags=unit
```