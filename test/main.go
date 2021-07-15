package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

type User struct {
	Name      string    `xorm:"varchar(25) 'name'"`
	Id        int       `xorm:"pk 'id' autoincr"`
	CreatedAt time.Time `xorm:"created"`
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/test")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = engine.CreateTables(User{})
	if err != nil {
		log.Fatal(err)
		return
	}
	//下面插入语句第一次运行程序时运行，后面注释掉
	/*
		u := make([]User, 5)
		u[0].Name = "abcd"
		u[1].Name = "acbd"
		u[2].Name = "dbac"
		u[3].Name = "cbda"
		u[4].Name = "bdca"

		_, err = engine.Insert(u)
		if err != nil {
			log.Fatal(err)
			return
		}
	*/

	/*	1) 传入Slice用于返回数据
		everyone := make([]Userinfo, 0)
		err := engine.Find(&everyone)

		pEveryOne := make([]*Userinfo, 0)
		err := engine.Find(&pEveryOne)*/

	users := make([]User, 0)
	err = engine.Find(&users)
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, u := range users {
		fmt.Println("只传第一个参数，传入slice用于返回数据，index:", i, "user", u.Name)
	}
	fmt.Println()

	users1 := make([]User, 0)
	err = engine.Find(&users1)
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, u := range users1 {
		fmt.Println("只传第一参数，传入slice用于返回数据，index:", i, "user", u.Name)
	}
	fmt.Println()

	users2 := make([]User, 0)
	s2 := new(User)
	s2.Name = "abcd"

	err = engine.Find(&users2, s2)
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, u := range users2 {
		fmt.Println("传两个参数，传入slice用于返回数据，index:", i, "user", u.Name)
	}

	fmt.Println()

	users3 := make([]User, 0)
	s3 := new(User)
	s3.Name = "abcd"
	err = engine.Find(&users3, s3)
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, u := range users3 {
		fmt.Println("只传第一参数，传入slice用于返回数据，index:", i, "user", u.Name)
	}

	fmt.Println()
	fmt.Println()

	/*2) 传入Map用户返回数据，map必须为map[int64]Userinfo的形式，map的key为id，因此对于复合主键无法使用这种方式。

	users := make(map[int64]Userinfo)
	err := engine.Find(&users)

	pUsers := make(map[int64]*Userinfo)
	err := engine.Find(&pUsers)*/

	users4 := make(map[int64]User, 0)
	err = engine.Find(&users4)
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, u := range users4 {
		fmt.Println("只传第一个参数，传入map用于返回数据，index:", i, "user", u.Name)
	}
	fmt.Println()

	users5 := make(map[int64]User, 0)
	s5 := new(User)
	s5.Name = "abcd"

	err = engine.Find(&users5, s5)
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, u := range users5 {
		fmt.Println("传两个参数，传入map用于返回数据，index:", i, "user", u.Name)
	}

	fmt.Println()
	fmt.Println()

	/*
		3) 也可以加入各种条件
		users := make([]Userinfo, 0)
		err := engine.Where("age > ? or name = ?", 30, "xlw").Limit(20, 10).Find(&users)
	*/

	users6 := make(map[int64]User, 0)
	err = engine.Where("name > ?", "baaa").Find(&users6)
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, u := range users6 {
		fmt.Println("只传第一个参数，传入map用于返回数据，index:", i, "user", u.Name)
	}
	fmt.Println()

	/*4) 如果只选择单个字段，也可使用非结构体的Slice
	var ints []int64
	err := engine.Table("user").Cols("id").Find(&ints)
	*/

	var ints []int64
	err = engine.Table("user").Cols("id").Find(&ints)

	if err != nil {
		log.Fatal(err)
		return
	}

	for i := range ints {
		fmt.Println("如果只选择单个字段，也可使用非结构体的Slice，index:", i)
	}
}

/*
输出：
只传第一个参数，传入slice用于返回数据，index: 0 user abcd
只传第一个参数，传入slice用于返回数据，index: 1 user acbd
只传第一个参数，传入slice用于返回数据，index: 2 user dbac
只传第一个参数，传入slice用于返回数据，index: 3 user cbda
只传第一个参数，传入slice用于返回数据，index: 4 user bdca

只传第一参数，传入slice用于返回数据，index: 0 user abcd
只传第一参数，传入slice用于返回数据，index: 1 user acbd
只传第一参数，传入slice用于返回数据，index: 2 user dbac
只传第一参数，传入slice用于返回数据，index: 3 user cbda
只传第一参数，传入slice用于返回数据，index: 4 user bdca

传两个参数，传入slice用于返回数据，index: 0 user abcd

只传第一参数，传入slice用于返回数据，index: 0 user abcd


只传第一个参数，传入map用于返回数据，index: 1 user abcd
只传第一个参数，传入map用于返回数据，index: 2 user acbd
只传第一个参数，传入map用于返回数据，index: 3 user dbac
只传第一个参数，传入map用于返回数据，index: 4 user cbda
只传第一个参数，传入map用于返回数据，index: 5 user bdca

传两个参数，传入map用于返回数据，index: 1 user abcd


只传第一个参数，传入map用于返回数据，index: 3 user dbac
只传第一个参数，传入map用于返回数据，index: 4 user cbda
只传第一个参数，传入map用于返回数据，index: 5 user bdca

如果只选择单个字段，也可使用非结构体的Slice，index: 0
如果只选择单个字段，也可使用非结构体的Slice，index: 1
如果只选择单个字段，也可使用非结构体的Slice，index: 2
如果只选择单个字段，也可使用非结构体的Slice，index: 3
如果只选择单个字段，也可使用非结构体的Slice，index: 4

*/
