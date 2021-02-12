package dbops

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn  *sql.DB
	err     error
	tempvid int
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(8.129.188.19:3306)/video_server_mysql?charset=utf8")
	if err != nil {
		panic(err.Error()) //无法完成正常逻辑
	}
	fmt.Println("database success !!!")

}
