package rds

import (
	"github.com/YangzhenZhao/easyvideo/go_kit_back_end/video/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(config config.RDSConfig) error {
	DB, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	VideoDao.InitConf(DB, config.StorageDir)
	return nil
}
