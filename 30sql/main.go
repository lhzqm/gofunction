package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面
}
