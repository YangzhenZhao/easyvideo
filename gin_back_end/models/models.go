package models

type Video struct {
	ID               uint   `gorm:"primaryKey"`
	Name             string `gorm:"index"`
	BytesSize        uint64 `gorm:"column:bytes_size"`
	VideoPath        string `gorm:"column:video_path"`
	CoverPictruePath string `gorm:"column:cover_picture_path"`
}

type Tag struct {
	ID      uint   `gorm:"primaryKey"`
	TagName string `gorm:"index;column:tag_name"`
}

type VideoTag struct {
	ID      uint `gorm:"primaryKey"`
	VideoId uint `gorm:"index;column:video_id"`
	TagId   uint `gorm:"index;column:tag_id"`
}

func (Video) TableName() string {
	return "video"
}

func (Tag) TableName() string {
	return "tag"
}

func (VideoTag) TableName() string {
	return "video_tag"
}
