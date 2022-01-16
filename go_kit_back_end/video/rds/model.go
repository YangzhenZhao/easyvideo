package rds

type Video struct {
	ID               int32  `gorm:"primaryKey"`
	Name             string `gorm:"index"`
	BytesSize        int64  `gorm:"column:bytes_size"`
	VideoPath        string `gorm:"column:video_path"`
	CoverPictruePath string `gorm:"column:cover_picture_path"`
}

func (Video) TableName() string {
	return "video"
}
