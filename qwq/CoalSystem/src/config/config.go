package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func init() {

	var err error

	db, err := gorm.Open("mysql", "qys:admin@tcp(192.168.1.4:3306)/mydb?charset=utf8")
	//db, err := gorm.Open("mysql", "root:123456@/mydb?charset=utf8")
	db.SingularTable(true)

	if err != nil {
		//panic(err)
		fmt.Println("mysql connect error")
	}
	fmt.Println(1111)

}
