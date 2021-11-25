package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UsersUpdate struct {
	gorm.Model
	//Name *string `gorm:"type:varchar(100);default:RandySun;comment:姓名"`
	Name string `gorm:"type:varchar(100);default:RandySun;comment:姓名"`
	Age  int64
	// 设置默认值
	Active bool
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 迁移表创建对应关系
	db.AutoMigrate(&UsersUpdate{})

	// 添加数据
	//db.Debug().Create(&UsersUpdate{Name: "bary", Age: 18})
	//db.Debug().Create(&UsersUpdate{Name: "RandySun", Age: 18})
	//db.Debug().Create(&UsersUpdate{Name: "Jack", Age: 18})
	//db.Debug().Create(&UsersUpdate{Name: "", Age: 19}) // 添加记录name默认为Null
	//
	//// 查询
	var user UsersUpdate
	db.Debug().First(&user)
	fmt.Printf("%#v", user)
	// 更新
	//user.Name = "xiaoSun"
	//user.Age = 18
	//// 保存 save默认更新所有字段
	//db.Debug().Save(&user)
	//
	//// 更新指定字段
	//db.Debug().Model(&user).Update("name","Randy")

	//db.Debug().Model(&user).Where("active = ?", false).Update("name", "bak")
	// 批量操作
	//m1 := map[string]interface{}{
	//	"name":"Randy",
	//	"age": 188,
	//	"active": false,
	//
	//}
	//// m1列出来的所有字段都会更新
	//db.Debug().Model(&user).Updates(m1)
	//
	//db.Debug().Model(&user).Updates(UsersUpdate{Name: "hello", Age: 0, Active: false})
	//// 只更新age字段
	//db.Debug().Model(&user).Select("age").Updates(m1)

	//// 排除m1中active更新其他的字段
	//db.Debug().Model(&user).Omit("active").Updates(m1)

	// 批量更新

	//update_res := db.Debug().Table("users_updates").Where("id IN (?)", []int{3, 4}).Updates(map[string]interface{}{"name": "xiao_zhi", "age": 18})
	//fmt.Printf("%#v\n", update_res)
	//rows := update_res.RowsAffected
	//fmt.Printf("%#v\n", rows)

	//res := db.Debug().Model(&user).Update("age", gorm.Expr("age * ? + ?", 2, 100))
	//fmt.Println(res.Row())
	//fmt.Println(res.RowsAffected)
	//db.Debug().Model(&user).Updates(map[string]interface{}{"age": gorm.Expr("age * ? + ?", 2, 100)})

	//db.Debug().Model(&user).UpdateColumn("age", gorm.Expr("age - ?", 1))

	//db.Debug().Model(&user).Where("age > 10").UpdateColumn("age", gorm.Expr("age - ?", 1))
	//db.Debug().Model(&user).Set("gorm:update_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Update("name", "hello")
	//db.Debug().Delete(&user)

	//db.Debug().Where("name LIKE ?", "%he%").Delete(&user)
	//db.Debug().Delete(UsersUpdate{}, "name LIKE ?", "%Randy%")

	db.Debug().Unscoped().Delete(&user)
}
