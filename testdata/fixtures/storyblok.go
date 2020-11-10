package fixtures

import (
	"bytes"
	"io"
	"io/ioutil"
)

// ResponseGetStory returns an example response for a story query
func ResponseGetStory() io.ReadCloser {
	json := `{"story":{"name":"Perseverance","created_at":"2020-09-21T21:14:17.767Z","published_at":"2020-09-22T17:04:18.000Z","alternates":[],"id":21231306,"uuid":"6a121842-7acd-41c1-a4df-ca6abae0bca4","content":{"_uid":"88b77448-364f-4935-9895-c90e1e606f37","trailer":{"id":null,"alt":null,"name":"","focus":null,"title":null,"filename":"","copyright":null,"fieldtype":"asset"},"benefits":[{"_uid":"d6bf904b-c1e3-4207-85c3-b6dd190d6eac","text":"This is a desc","image":{"id":1442746,"alt":null,"name":"","focus":null,"title":null,"filename":"https://a.storyblok.com/f/86546/1380x960/8efa149fce/benefit1-2x.jpg","copyright":null,"fieldtype":"asset"},"number":"Benefit 1","component":"benefit"}],"component":"page-program","equipment":{"type":"doc","content":[{"type":"paragraph","content":[{"text":"Equipment","type":"text"}]}]},"unique_id":{"_uid":"44527235-b0a4-4658-ab9b-65d1941b5a1b","plugin":"uuiddisplay"},"categories":[{"_uid":"86a9f011-2326-4fad-b1ea-6d91524eda9a","title":"Week 1","component":"program-category","description":"","subcategories":[{"_uid":"a0a5fada-199f-4d5d-acdd-0897b5f3fe5c","title":"Day 1","modules":["3008164d-500d-4758-a96a-964099584582","0a45369e-cec2-49c9-8e4b-7c2fb87d075f","a79c64a9-0be4-4e21-b338-461fdbbb3472","65e4cbf4-6b20-4947-9622-097e80b69085","0ad431ab-a8c1-4474-9809-106b32c68f2d"],"component":"program-subcategory","unique_id":{"_uid":"06ee2627-c3fb-4d54-8bcb-3d53af11c326","plugin":"uuiddisplay"},"description":""}]}],"hero_image":{"id":1413700,"alt":null,"name":"","focus":null,"title":null,"filename":"https://a.storyblok.com/f/86546/2736x992/2fa86e4d39/gabrielle-elleirbag.jpg","copyright":null,"fieldtype":"asset"},"preview_title":"Previews title","program_guide":{"id":null,"alt":null,"name":"","focus":null,"title":null,"filename":"","copyright":null,"fieldtype":"asset"},"hours_per_week":"24","preview_modules":[{"_uid":"9b7db633-3701-48de-8f58-c76a57a7f9dd","image":{"id":1431436,"alt":null,"name":"","focus":null,"title":null,"filename":"https://a.storyblok.com/f/86546/480x270/4f76054fd2/core-2x.jpg","copyright":null,"fieldtype":"asset"},"title":"Title ","component":"preview","description":"kjhkjh hkj kj hkh kjh k"}],"long_description":{"type":"doc","content":[{"type":"paragraph","content":[{"text":"This is a long description","type":"text"}]}]},"preview_subtitle":"Preview subtitle","short_description":"This is a short description"},"slug":"perseverance","full_slug":"programs/perseverance","default_full_slug":null,"sort_by_date":null,"position":30,"tag_list":[],"is_startpage":false,"parent_id":20223993,"meta_data":null,"group_id":"b84aefe1-e6b9-4529-90a4-240164510054","first_published_at":"2020-09-22T17:04:18.490Z","release_id":null,"lang":"default","path":"","translated_slugs":[]}}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}

// ResponseStoryNotFound returns an example response for an invalid story slug/id/uuid
func ResponseStoryNotFound() io.ReadCloser {
	json := `{"error":["This record could not be found"]}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}

// ResponseUnauthorized returns an example Unauthorized response
func ResponseUnauthorized() io.ReadCloser {
	json := `{"error":"Unauthorized"}`
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}
