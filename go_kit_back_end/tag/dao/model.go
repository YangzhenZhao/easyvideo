package dao

type Tag struct {
	ID      uint   `gorm:"primaryKey"`
	TagName string `gorm:"index;column:tag_name"`
}

type VideoTag struct {
	ID      uint `gorm:"primaryKey"`
	VideoId uint `gorm:"index;column:video_id"`
	TagId   uint `gorm:"index;column:tag_id"`
}

func (Tag) TableName() string {
	return "tag"
}

func (VideoTag) TableName() string {
	return "video_tag"
}
