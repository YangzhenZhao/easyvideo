package rds

type Tag struct {
	ID      int32  `gorm:"primaryKey"`
	TagName string `gorm:"index;column:tag_name"`
}

type VideoTag struct {
	ID      int32 `gorm:"primaryKey"`
	VideoId int32 `gorm:"index;column:video_id"`
	TagId   int32 `gorm:"index;column:tag_id"`
}

func (Tag) TableName() string {
	return "tag"
}

func (VideoTag) TableName() string {
	return "video_tag"
}
