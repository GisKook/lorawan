package db

import (
	"github.com/giskook/lorawan/base"
	"github.com/lib/pq"
)

const (
	SQL_USER_ADD string = "insert into lorawan_user(id, alias, passwd) values ($1, $2,$3)"
	SQL_USER_ROLE_ADD string = "insert into lorawan_user_role(user_id, role_id) values($1, $2)"
)

func (db_socket *DBSocket) UserAdd(id string, alias string, passwd string, role_id string,c chan *base.DBResult) {
	tx, err := db_socket.db.Begin()
	if err != nil{
		tx.Rollback()
		c <- &base.DBResult{
			Err:base.NewErr(err, base.ERR_DB_BEGIN_TRANSCATION_CODE, base.ERR_DB_BEGIN_TRANSCATION_DESC),
		}
		return
	}

	passwd_secret := db_socket.gen_passwd(passwd, db_socket.conf.Secret)
	_, err = tx.Exec(SQL_USER_ADD, id, alias, passwd_secret)
	if err != nil{ 
		tx.Rollback()
		if err.(*pq.Error).Code == "23505"{ 
			c<-&base.DBResult{
				Err:base.NewErr(err, base.ERR_USER_ALREADY_EXIST_CODE, base.ERR_USER_ALREADY_EXIST_DESC),
			}
		}else{
			c<-&base.DBResult{
				Err:base.ERROR_NOT_CAPTURE,
			}
		}
		return
	}

	_, err = tx.Exec(SQL_USER_ROLE_ADD, id, role_id)
	if err != nil{ 
		tx.Rollback()
		c <- &base.DBResult{
			Err:base.NewErr(err, base.ERR_USER_ROLE_ADD_CODE, base.ERR_USER_ROLE_ADD_DESC),
		}
		return
	}

	err = tx.Commit()
	if err != nil{
		c <- &base.DBResult{
			Err:base.NewErr(err, base.ERR_DB_COMMIT_TRANSCATION_CODE, base.ERR_DB_COMMIT_TRANSCATION_DESC),
		}
		return
	}
	c <- &base.DBResult{
		Err:err,
	}
}