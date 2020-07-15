package main

/*
	Go实战--go语言操作sqlite数据库(The way to go)
	https://blog.csdn.net/wangshubo1989/article/details/70313667

	https://www.thepolyglotdeveloper.com/2017/04/using-sqlite-database-golang-application/
*/

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//打开数据库, 如果不存子， 则创建
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	//创建表
	sql_table := `
	CREATE TABLE IF NOT EXISTS userinfo(
	uid INTEGER PRIMARY KEY AUTOINCREMENT,
	username VARCHAR(64) NULL,
	departname VARCHAR(64) NULL,
	created DATE NULL
	);`

	//stmt, _ := db.Prepare(sql_table)
	//stmt.Exec()

	_, err = db.Exec(sql_table)
	checkErr(err)

	//insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("wangshubo", "国务院", "2017-04-21")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("wangshubo_new", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err := rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close()

	//delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
