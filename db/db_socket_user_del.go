package db

import (
	"github.com/giskook/lorawan/base"
	"context"
	"log"
)

const (
	SQL_USER_DEL string = "delete from lorawan_user where 1=1 %s"
)

func (db_socket *DBSocket) UserDel(ctx context.Context, id string, c chan *base.DBResult) { 
	var where_clause string
	if id != ""{
		where_clause += db_socket.gen_where_clause_string(SQL_WHERE_CLAUSE_EQ_STRING, COL_LORAWAN_USER_ID, id)
	}else{
		c <- &base.DBResult{ 
			Err:base.NewErr(nil, base.ERR_USER_DEL_NO_ID_CODE, base.ERR_USER_DEL_NO_ID_DESC),
		}

		return 
	}


	sql_del_sql := db_socket.fmt_where_clause_sql(SQL_USER_DEL, where_clause) 
	log.Println(sql_del_sql)
	_, err := db_socket.db.ExecContext(ctx, sql_del_sql)
	if err != nil {
		log.Println(err)
		c <- &base.DBResult{ 
			Err:base.NewErr(err, base.ERR_USER_DEL_CODE, base.ERR_USER_DEL_DESC),
		}

		return
	}

	c<-&base.DBResult{
		Err:base.ERROR_NONE,
	}
}