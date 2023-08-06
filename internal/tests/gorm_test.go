package tests

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (*User) TableName() string {
	return "user"
}

func TestGormFindOne(t *testing.T) {

	dsn := "root:root@tcp(120.77.176.211:53306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("connect DB error, ", err)
	}

	user := &User{
		Username: "tom",
	}

	result := db.First(user)
	if result.Error != nil {
		t.Fatal("query error: ", result.Error)
	}

	// 显示查询结果
	if result.RowsAffected > 0 {
		t.Logf("\nUser info: %#v\n", user)
	} else {
		t.Log("User not found")
	}
}
