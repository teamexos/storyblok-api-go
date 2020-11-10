package storyblok

// Alternate is a struct of an alternate story object
type Alternate struct {
	ID       int    `json:"id"`
	Name     string `json:"string"`
	Slug     string `json:"slug"`
	FullSlug string `json:"full_slug"`
	IsFolder bool   `json:"is_folder"`
	ParentID int    `json:"parent_id"`
}
