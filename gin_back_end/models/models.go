package models

type Video struct {
	ID                 uint   `gorm:"primaryKey"`
	name               string `gorm:"index"`
	bytes_size         uint64
	video_path         string
	cover_pictrue_path string
}

type Tag struct {
	ID       uint   `gorm:"primaryKey"`
	tag_name string `gorm:"index"`
}

type VideoTag struct {
	ID       uint `gorm:"primaryKey"`
	video_id uint `gorm:"index"`
	tag_id   uint `gorm:"index"`
}
