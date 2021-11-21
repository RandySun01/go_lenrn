package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);comment:姓名"`
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// TableName 将 User 的表名设置为 `profiles`
func (User) TableName() string {
	return "user_randy"
}

// Animal 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64  `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(100);comment:姓名"`
	Age      int64
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 迁移表创建对应关系
	db.AutoMigrate(&User{}, &Animal{})

}
