package sqlite

import (
	"Refined_service/biz/model/comment"
	"Refined_service/biz/model/post"
	"Refined_service/biz/model/user"
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var err error
	dbfile := viper.GetString("database.sqlite.file")
	DB, err = gorm.Open(sqlite.Open(dbfile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	_ = DB.AutoMigrate(&comment.Comment{}, &post.Post{}, &user.User{})
	if err != nil {
		panic(err)
	}
}
