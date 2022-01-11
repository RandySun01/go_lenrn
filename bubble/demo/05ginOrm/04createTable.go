package main

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
type CreateUsers struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreditCard   CreditCard `gorm:"foreignKey:ID"`
}

func (u *CreateUsers) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("创建之前触发钩子")
	if u.Name == "RandySun" {
		return errors.New("invalid role")
	}
	return
}
func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 迁移表创建对应关系
	db.AutoMigrate(&CreateUsers{}, &CreditCard{})

	// 创建数据
	//timeNow := time.Now()
	//user := CreateUsers{Name: "RandySun", Age: 18, Birthday: &timeNow}

	// 创建用户
	//result := db.Debug().Create(&user) // 通过数据的指针来创建
	//
	//fmt.Println(user.ID)             // 返回插入数据的主键
	//fmt.Println(result.Error)        // 返回 error
	//fmt.Println(result.RowsAffected) // 返回插入记录的条数
	// 插入指定字段
	//db.Debug().Select("Name", "Age", "CreatedAt").Create(&user)

	// 插入排除字段
	//db.Debug().Omit("Name", "Age", "CreatedAt").Create(&user)

	// 批量插入数据
	//var users = []CreateUsers{{Name: "RandySun01"}, {Name: "RandySun02"}, {Name: "RandySun03"}}
	//db.Debug().Create(&users)
	//for _, user := range users {
	//	fmt.Println(user.ID) // 4,5,6
	//}

	// 分批创建
	//var users = []CreateUsers{{Name: "RandySun01"}, {Name: "RandySun02"},{Name: "RandySun04"},{Name: "RandySun0...."}, {Name: "RandySun200000"}}
	//
	//// 数量为 2
	//db.Debug().CreateInBatches(users, 2)
	//
	//for _, user := range users {
	//	fmt.Println(user.ID) // 4,5,6
	//}
	// 根据 Map 创建

	//db.Debug().Model(&CreateUsers{}).Create(map[string]interface{}{
	//	"Name": "RandySun", "Age": 18,
	//})
	//
	//// batch insert from `[]map[string]interface{}{}`
	//db.Debug().Model(&CreateUsers{}).Create([]map[string]interface{}{
	//	{"Name": "RandySunMap01", "Age": 18},
	//	{"Name": "RandySunMap02", "Age": 20},
	//})

	// 关联创建
	//
	//db.Debug().Create(&CreateUsers{
	//	Name:       "Randy",
	//	CreditCard: CreditCard{Number: "34353435"},
	//})

	//// 查询
	//var user CreateUsers
	//db.Debug().First(&user)
	//fmt.Printf(user.Name)

	//// 获取一条记录，没有指定排序字段
	//db.Debug().Take(&user)
	//fmt.Printf(user.Name)

	//// 获取最后一条记录（主键降序）
	//db.Debug().Last(&user)
	//fmt.Println(user.ID, user.Name)

	//result := db.First(&user)
	//// 返回找到的记录数
	//fmt.Println(result.RowsAffected)
	//fmt.Println(result.Error) // returns error or nil
	//
	//// 检查 ErrRecordNotFound 错误
	//isErrors := errors.Is(result.Error, gorm.ErrRecordNotFound)
	//fmt.Println(isErrors)

	//result := map[string]interface{}{}
	//db.Debug().Model(&CreateUsers{}).First(&result)
	//fmt.Printf("%#v\n", result)
	//fmt.Printf("%#v\n", result["name"])

	// 无效
	//result := map[string]interface{}{}
	//db.Table("create_users").First(&result)
	//fmt.Printf("%#v\n", result)
	//fmt.Printf("%#v\n", result["name"])

	// 配合 Take 有效
	//result := map[string]interface{}{}
	//db.Table("create_users").Take(&result) // 表名
	//fmt.Printf("%#v\n", result)
	//fmt.Printf("%#v\n", result["name"])

	// 根据主键查询第一条记录
	//var user CreateUsers
	//db.Debug().First(&user, 5)
	//fmt.Println(user.ID)

	// 根据主键查询第一条记录
	//var user CreateUsers
	//db.Debug().First(&user, "10")
	//fmt.Println(user.ID)

	// 根据主键查询第多条记录
	//var users []CreateUsers
	//db.Debug().Find(&users, []int{1,2,3})
	//
	//for _, u :=range users{
	//	fmt.Println(u.ID)
	//}

	//// 获取全部记录
	//var users []CreateUsers
	//result := db.Debug().Find(&users)
	//
	//fmt.Println(result.RowsAffected) // 返回找到的记录数，相当于 `len(users)`
	//fmt.Println(result.Error)        // returns error
	// 获取满足条件第一条匹配的记录
	//var user CreateUsers
	//db.Debug().Where("name = ?", "RandySun").First(&user)
	//fmt.Println(user.ID)

	// 获取全部匹配的记录
	//var users []CreateUsers
	//db.Debug().Where("name <> ?", "RandySun").Find(&users)
	//for _, u := range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	// IN
	//var users []CreateUsers
	//
	//db.Debug().Where("name IN ?", []string{"RandySun", "RandySun03"}).Find(&users)
	//for _, u := range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	// LIKE
	//var users []CreateUsers
	//
	//db.Debug().Where("name LIKE ?", "%sun%").Find(&users)
	//for _, u := range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	// AND
	//var users []CreateUsers
	//
	//db.Debug().Where("name = ? AND age >= ?", "RandySun", "18").Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	//for _, u := range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	// Time
	//var users []CreateUsers
	//lasterWeek := time.Now()
	//
	//db.Debug().Where("updated_at < ?", lasterWeek).Find(&users)
	//for _, u := range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	//// BETWEEN
	//var users []CreateUsers
	//
	//
	//db.Debug().Where("age BETWEEN ? AND ?", 1, 18).Find(&users)
	//for _, u := range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	// Struct
	//var user CreateUsers
	//db.Debug().Where(&CreateUsers{Name: "RandySun", Age: 18}).First(&user)
	//fmt.Println(user.ID, user.Name)

	//// Map
	//var users []CreateUsers
	//db.Debug().Where(map[string]interface{}{"name": "RandySun", "age": 18}).Find(&users)
	//
	//for _, u := range users{
	//	fmt.Println(u.ID, u.Name)
	//}
	//

	//
	//SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 主键切片条件
	//var users []CreateUsers
	//
	//db.Debug().Where([]int64{18, 8, 1}).Find(&users)
	//for _, u := range users{
	//	fmt.Println(u.ID, u.Name)
	//}
	//var users []CreateUsers
	//db.Debug().Where(&CreateUsers{Name: "RandySun"}, "name", "Age").Find(&users)
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	//var users []CreateUsers
	//db.Debug().Where(&CreateUsers{Name: "RandySun"}, "Age").Find(&users)
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}
	//SELECT * FROM users WHERE age = 0;

	// 内联
	//var user CreateUsers
	//db.Debug().First(&user, 18)
	//fmt.Println(user.ID, user.Name)
	// 根据主键获取记录, 如果它是一个非整形主键

	//var user CreateUsers
	//db.Debug().First(&user, "name = ?", "RandySun01")
	//fmt.Println(user.ID, user.Name)
	//
	//var user1 CreateUsers
	//db.Debug().First(&user1, "id = ?", "2")
	//fmt.Println(user1.ID, user1.Name)

	//var users []CreateUsers
	//// find可群查和单查
	//db.Debug().Find(&users, "name = ?", "RandySun")
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}
	//
	//var user CreateUsers
	//// find可群查和单查
	//db.Debug().Find(&user, "name = ?", "RandySun")
	//fmt.Println(user.ID, user.Name)

	//var users []CreateUsers
	//db.Debug().Find(&users, "name <> ? AND age >= ?", "RandySun", 18)
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}
	//var users []CreateUsers
	//db.Debug().Find(&users, CreateUsers{Age: 18})
	//
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	//var users []CreateUsers
	//db.Debug().Find(&users, map[string]interface{}{"age": 20})
	//
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}
	//var user CreateUsers
	//db.Debug().Not("name", "RandySun").First(&user)
	//fmt.Println(user.ID, user.Name)

	//var users []CreateUsers
	//db.Debug().Not("name", []string{"RandySun", "RandySun02"}).Find(&users)
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	// Not In slice of primary keys

	//var users []CreateUsers
	//db.Debug().Not([]int64{1,2,3}).Find(&users)
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	//var users []CreateUsers
	//db.Debug().Not([]int64{}).Find(&users)
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	// Plain SQL
	//var users []CreateUsers
	//// SELECT * FROM `create_users` WHERE NOT name = 'RandySun' AND `create_users`.`deleted_at` IS NULL
	//db.Debug().Not("name = ?", "RandySun").Find(&users)
	//
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	// Struct
	//var users []CreateUsers
	//// SELECT * FROM `create_users` WHERE `create_users`.`name` <> 'RandySun' AND `create_users`.`deleted_at` IS NULL
	//db.Debug().Not(CreateUsers{Name: "RandySun"}).Find(&users)
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	//var users []CreateUsers
	//
	//// SELECT * FROM `create_users` WHERE (name = 'RandySun' OR age = 20) AND `create_users`.`deleted_at` IS NULL
	//db.Debug().Where("name = ?", "RandySun").Or("age = ?", 20).Find(&users)
	//
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	//var users []CreateUsers
	//// SELECT * FROM `create_users` WHERE ((name = 'RandySun') OR `create_users`.`name` = 'RandySun02') AND `create_users`.`deleted_at` IS NULL
	//db.Debug().Where("name = 'RandySun'").Or(CreateUsers{Name: "RandySun02"}).Find(&users)
	//
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	//var users []CreateUsers
	//// SELECT * FROM `create_users` WHERE ((name = 'RandySun') OR `name` = 'RandySun02') AND `create_users`.`deleted_at` IS NULL
	//db.Debug().Where("name = 'RandySun'").Or(map[string]interface{}{"name": "RandySun02"}).Find(&users)
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	// 多个 order
	//var users []CreateUsers
	//// SELECT * FROM `create_users` WHERE `create_users`.`deleted_at` IS NULL ORDER BY age desc,name
	//db.Debug().Order("age desc").Order("name").Find(&users)
	//
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	// 多个 order
	//var users []CreateUsers
	//// SELECT * FROM `create_users` WHERE `create_users`.`deleted_at` IS NULL ORDER BY age desc,name
	//db.Debug().Order("age desc").Order("name").Find(&users)
	//for _, u := range users {
	//	fmt.Println(u.ID, u.Name)
	//}

	// 自定义排序ORDER BY FIELD()

	//var users []CreateUsers
	//db.Debug().Clauses(clause.OrderBy{
	//	Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	//}).Find(&users)
	//
	//for _, u := range users {
	//	fmt.Println(u.ID, u.Name)
	//}

	//var users []CreateUsers
	////  SELECT * FROM `create_users` WHERE `create_users`.`deleted_at` IS NULL LIMIT 3
	//db.Debug().Limit(3).Find(&users)
	//
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	//// 通过 -1 消除 Limit 条件
	//var users1 []CreateUsers
	//var users2 []CreateUsers
	//
	////  SELECT * FROM `create_users` WHERE `create_users`.`deleted_at` IS NULL LIMIT 10
	////  SELECT * FROM `create_users` WHERE `create_users`.`deleted_at` IS NULL
	//db.Debug().Limit(10).Find(&users1).Limit(-1).Find(&users2)
	//
	//for _, u :=range users1{
	//	fmt.Println(u.ID, u.Name)
	//}
	//fmt.Println("*********************************************************")
	//for _, u :=range users2{
	//	fmt.Println(u.ID, u.Name)
	//}

	//	fmt.Println(u.ID, u.Name)

	//var users []CreateUsers
	//db.Debug().Offset(3).Find(&users)
	//
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	//var users []CreateUsers
	//// SELECT * FROM `create_users` WHERE `create_users`.`deleted_at` IS NULL LIMIT 10 OFFSET 5
	//db.Debug().Limit(10).Offset(5).Find(&users)
	//
	//for _, u :=range users{
	//	fmt.Println(u.ID, u.Name)
	//}

	//var users1 []CreateUsers
	//var users2 []CreateUsers
	//db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	//
	//for _, u :=range users1{
	//	fmt.Println(u.ID, u.Name)
	//}
	//fmt.Println("*********************************************************")
	//for _, u :=range users2{
	//	fmt.Println(u.ID, u.Name)
	//}
	//type Result struct {
	//	Date  time.Time
	//	Total int
	//	name string
	//}
	//
	//var result Result
	//// SELECT name, sum(age) as total FROM `create_users` WHERE name LIKE 'Randy%' AND `create_users`.`deleted_at` IS NULL GROUP BY `name` ORDER BY `create_users`.`id` LIMIT 1
	//db.Debug().Model(&CreateUsers{}).Select("name, sum(age) as total").Where("name LIKE ?", "Randy%").Group("name").First(&result)
	//
	//fmt.Println(result.Date, result.Total, result.name)

	//
	//type Result struct {
	//	Date  time.Time
	//	Total int
	//	name string
	//}
	//
	//var result []Result
	//db.Debug().Model(&CreateUsers{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "RandySun").Find(&result)
	//for _, u :=range result{
	//	fmt.Println(u.name, u.Total, u.Date)
	//}

	//type Result struct {
	//	Date  *time.Time
	//	Total int
	//	name string
	//}
	//
	//// SELECT name , sum(age) as total FROM `create_users` GROUP BY `name`
	//rows, err := db.Debug().Table("create_users").Select("name , sum(age) as total").Group("name").Rows()
	//for rows.Next() {
	//	var r Result
	//	err := rows.Scan(&r.name, &r.Total)
	//	if err != nil {
	//		fmt.Printf("scan failed, err:%v\n", err)
	//		return
	//	}
	//	fmt.Printf("name: %s, Total: %d\n", r.name, r.Total)
	//}

	//type Result struct {
	//	Date  *time.Time
	//	Total int
	//	name string
	//}
	//rows, err := db.Debug().Table("create_users").Select("name, sum(age) as total").Group("name").Having("sum(age) > ?", 10).Rows()
	//
	//for rows.Next() {
	//	var r Result
	//	err := rows.Scan(&r.name, &r.Total)
	//	if err != nil {
	//		fmt.Printf("scan failed, err:%v\n", err)
	//		return
	//	}
	//	fmt.Printf("name: %s, Total: %d\n", r.name, r.Total)
	//}

	//type Result struct {
	//	names string
	//	Total int
	//	Date  *time.Time
	//}
	//var r []Result
	//db.Debug().Table("create_users").Select("name, sum(age) as total").Group("name").Having("sum(age) > ?", 10).Scan(&r)
	//for _, r := range r{
	//	fmt.Println(r.names, r.Total)
	//}

	//type APIUser struct {
	//	ID   uint
	//	Name string
	//}
	//// 查询时会自动选择 `id`, `name` 字段
	//var apiUser []APIUser
	////  SELECT `create_users`.`id`,`create_users`.`name` FROM `create_users` WHERE `create_users`.`deleted_at` IS NULL LIMIT 10
	//db.Debug().Model(&CreateUsers{}).Limit(10).Find(&apiUser)
	//for _, u := range apiUser{
	//	fmt.Println(u.ID, u.Name)
	//}

	//var users []CreateUsers
	//db.Debug().Session(&gorm.Session{QueryFields: true}).Find(&users)
	//for _, u := range users{
	//	fmt.Println(u.Name, u.Age, u.Birthday, u.CreatedAt)
	//}

	//var users []CreateUsers
	//// SELECT * FROM `create_users` WHERE age > (SELECT AVG(age) FROM `create_users`) AND `create_users`.`deleted_at` IS NULL
	//db.Debug().Where("age > (?)", db.Debug().Table("create_users").Select("AVG(age)")).Find(&users)
	//for _, u := range users {
	//	fmt.Println(u.Name, u.Age, u.Birthday, u.CreatedAt)
	//}



	//type Result struct {
	//	names string
	//	Total float64
	//	Date  *time.Time
	//}
	//var result []Result
	//subQuery := db.Debug().Select("AVG(age)").Where("name LIKE ?", "Randy%").Table("create_users")
	//db.Debug().Select("AVG(age) as Total").Group("name").Having("AVG(age) > (?)", subQuery).Table("create_users").Find(&result)
	//for _, r := range result{
	//	fmt.Println(r.Total)
	//}
	//
	//rows, err := db.Select("AVG(age) as Total").Group("name").Having("AVG(age) > (?)", subQuery).Table("create_users").Rows()
	//
	//for rows.Next() {
	//	var r Result
	//	err := rows.Scan(&r.Total)
	//	if err != nil {
	//		fmt.Printf("scan failed, err:%v\n", err)
	//		return
	//	}
	//	fmt.Printf("Total: %d\n", r.Total)
	//}


	//var users []CreateUsers
	//// SELECT * FROM (SELECT `name`,`age`,`deleted_at` FROM `create_users` WHERE `create_users`.`deleted_at` IS NULL) as u WHERE age = 18 AND `u`.`deleted_at` IS NULL
	//db.Debug().Table("(?) as u", db.Model(&CreateUsers{}).Select("name", "age", "deleted_at")).Where("age = ?", 18).Find(&users)
	//
	//for _, u := range users {
	//	fmt.Println(u.Name, u.Age, u.Birthday, u.CreatedAt)
	//}


	var users []CreateUsers
	subQuery1 := db.Model(&CreateUsers{}).Select("name", "deleted_at")
	subQuery2 := db.Model(&CreateUsers{}).Select("name", "deleted_at")
	// 	SELECT * FROM (SELECT `name`,`deleted_at` FROM `create_users` WHERE `create_users`.`deleted_at` IS NULL) as u, (SELECT `name`,`deleted_at` FROM `create_users` WHERE `create_users`.`deleted_at` IS NULL) as p WHERE `u`.`deleted_at` IS NULL
	db.Debug().Table("(?) as u, (?) as p", subQuery1, subQuery2).Find(&users)

	for _, u := range users {
		fmt.Println(u.Name, u.Age, u.Birthday, u.CreatedAt)
	}

}
