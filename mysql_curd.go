package main

import (
	"database/sql"
	"fmt"

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
	res, err := stmt.Exec("bb", "12", "23")
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

	//user_add()
	//user_list()
	//user_edit()
	//user_del()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
