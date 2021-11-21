package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	//Name *string `gorm:"type:varchar(100);default:RandySun;comment:姓名"`
	Name sql.NullString `gorm:"type:varchar(100);default:RandySun;comment:姓名"`
	Age  int64
	// 设置默认值
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 迁移表创建对应关系
	db.AutoMigrate(&Users{})
	//db.Debug().Create(&Users{Name: "randy", Age: 18})
	//db.Debug().Create(&Users{Age: 18}) // 添加记录name默认为Null

	db.Debug().Create(&Users{Name: sql.NullString{String: "yruy", Valid: true}, Age: 19}) // 添加记录name默认为Null

}
