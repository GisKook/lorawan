package db

import (
	"github.com/giskook/lorawan/base"
)

const (
	SQL_USER_INFO_GET string = "select alias from lorawan_user where id=$1"
)

func (db_socket *DBSocket) UserInfoGet(id string, c chan *base.DBResult) {
	var username string
	err := db_socket.db.QueryRow(SQL_USER_INFO_GET, id).Scan(&username)
	c<-&base.DBResult{ 
		Err:err,
		Extra:username,
	}
}