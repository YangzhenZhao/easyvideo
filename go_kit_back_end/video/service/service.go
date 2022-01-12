package service

type VideoViewItem struct {
	Name      string   `json:"name"`
	BytesSize uint64   `json:"bytesSize"`
	Tags      []string `json:"tags"`
}

type VideoService interface {
	AllVideos() ([]VideoViewItem, error)
}

type videoService struct{}

func (videoService) AllVideos() ([]VideoViewItem, error) {
	resList := []VideoViewItem{}

	return resList, nil
}
