package youtube

import "fmt"

type YoutubeLib interface {
	ListVideos() []video
}

type video struct {
	id    int
	title string
	url   string
}

type youtube struct {
}

func (y *youtube) ListVideos() []video {
	fmt.Println("get video list from source")
	return []video{
		{id: 1, title: "bird", url: "http://example.com"},
		{id: 2, title: "cat", url: "http://example.com"},
		{id: 3, title: "dog", url: "http://example.com"},
		{id: 4, title: "human", url: "http://example.com"},
	}
}
