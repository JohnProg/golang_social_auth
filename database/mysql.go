package database

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/golang_social_auth/settings"
	"github.com/jinzhu/gorm"
	"fmt"
)

var DB *gorm.DB

func ConnectMysql() error {
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		settings.Config.Database.User, settings.Config.Database.Password, settings.Config.Database.Dbname)

	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}
