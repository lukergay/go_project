package dbops

import (
	"database/sql"
	"log"
	"time"

	"defs"
	"utils"
	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {
	stemIns, err := dbConn.Prepare("INSERT INTO users(login_name,pwd) VALUES (?,?)")
	if err != nil {
		return err
	}
	defer stemIns.Close()

	_, err = stemIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stemOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name=?")
	defer stemOut.Close()

	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	_, err = stemOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows { /*ErrNoRows 表示查询的列表为空*/
		return nil
	}
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stemDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? and pwd=?")
	defer stemDel.Close()

	if err != nil {
		log.Printf("DeleteUser error %s", err)
		return err
	}

	_, err = stemDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	//creatime --> db
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") // M D y, HH:MM:SS
	stmtIns, err := dbConn.Prepare('INSERT INTO video_info 
			(id,author_id,name,display_ctime) VALUES(?,?,?,?)')

	if err !=nil {
		return nil,err
	}		

	_,err=stmtIns.Exec(vid,aid,name,ctime)
	if err!=nil{
		return nil,err
	}

	res:=&defs.VideoInfo(Id:vid,AuthorId:aid,Name:name,DisplayCtimer:ctime)
	defer stmtIns.Close()
	return res
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stemOut, err := dbConn.Prepare("SELECT author_id,name,display_ctime FROM video_info WHERE id=?")
	defer stemOut.Close()
	if err !=nil {
		return nil,err
	}		

	var aid int
	var dct string
	var name string

	_,err=stmtIns.QueryRow(vid).Scan(&aid,&name,&dct)
	if err != nil && err != sql.ErrNoRows { /*ErrNoRows 表示查询的列表为空*/
		return nil,nil
	}

	res:=&defs.VideoInfo(Id:vid,AuthorId:aid,Name:name,DisplayCtimer:dct)
	return res,nil
}

func DeleteVideoInfo(vid string) error{
	stemDel, err := dbConn.Prepare("DELETE FROM video_server WHERE id=?")
	defer stemDel.Close()

	if err != nil {
		log.Printf("DeleteUser error %s", err)
		return err
	}

	_, err = stemDel.Exec(vid)
	if err != nil {
		return err
	}
	return nil
}