package youtube

type youtubeCached struct {
	youtube *youtube
	videos  []video
}

func NewYoutubeCached() *youtubeCached {
	return &youtubeCached{
		youtube: &youtube{},
		videos:  make([]video, 0),
	}
}

func (y *youtubeCached) ListVideos() []video {
	if len(y.videos) == 0 {
		y.videos = y.youtube.ListVideos()
	}
	return y.videos
}
