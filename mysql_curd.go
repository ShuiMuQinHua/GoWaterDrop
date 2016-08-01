package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type UserModel struct {
}

func init() {
	GetDbCon()
}

func GetDbCon() {
	db, _ = sql.Open("mysql", "root:root@/ci_cms?charset=utf8")
}

func user_list() {
	//查询数据
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	for rows.Next() {
		var id int
		var user_name string
		var pwd string
		var roles string
		err = rows.Scan(&id, &user_name, &pwd, &roles)
		checkErr(err)

		fmt.Println(id)
	}
}

func user_add() {
	//插入数据
	stmt, err := db.Prepare("INSERT user SET user_name=?,pwd=?,roles=?")
	checkErr(err)
	res, err := stmt.Exec("bbbb", "12", "23")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func user_edit() {
	//更新数据
	stmt, _ := db.Prepare("update user set user_name=?,pwd=? where id=?")
	res, _ := stmt.Exec("dd", "245", "1")
	affect, _ := res.RowsAffected()
	if affect > 0 {
		fmt.Println("更新成功")
	}
}

func user_del() {
	rslt, _ := db.Exec("delete from `user` where id = ?", "1")
	affected_row, _ := rslt.RowsAffected()
	if affected_row > 0 {
		fmt.Println("删除成功")
	}
}

func main() {
	//根据命令行参数调取方法
	args := os.Args[1]
	switch {
	case args == "1":
		user_add()
		break
	case args == "2":
		user_del()
		break
	case args == "3":
		user_list()
		break
	case args == "4":
		user_edit()
		break
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
