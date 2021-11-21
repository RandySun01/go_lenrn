package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UsersSelect struct {
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
	db.AutoMigrate(&UsersSelect{})

	// 添加数据
	//db.Debug().Create(&UsersSelect{Name: sql.NullString{String: "bary", Valid: true}, Age: 18})
	//db.Debug().Create(&UsersSelect{Name: sql.NullString{String: "RandySun", Valid: true}, Age: 18})
	//db.Debug().Create(&UsersSelect{Name: sql.NullString{String: "Jack", Valid: true}, Age: 18})
	//db.Debug().Create(&UsersSelect{Name: sql.NullString{String: "", Valid: true}, Age: 19}) // 添加记录name默认为Null

	// 查询
	var user UsersSelect
	db.Debug().First(&user)
	fmt.Printf("%#v", user)

}
