package db

import (
	"github.com/giskook/lorawan/base"
	"context"
	"log"
)

const (
	SQL_USER_SEARCH string = " select lorawan_user.id, lorawan_user.alias, lorawan_role.name from lorawan_user, lorawan_role ,lorawan_user_role where lorawan_user.id=lorawan_user_role.user_id and lorawan_user_role.role_id=lorawan_role.id  %s limit $1 offset $2"
	SQL_USER_SEARCH_COUNT string = "select count(*) from lorawan_user where 1=1 %s"
)

func (db_socket *DBSocket) UserSearch(ctx context.Context, id , alias string, limit,offset int,c chan *base.DBResult) { 
	var where_clause string
	if id != ""{
		where_clause += db_socket.gen_where_clause_string(SQL_WHERE_CLAUSE_EQ_STRING, COL_LORAWAN_USER_ID, id)
	}

	if alias != ""{
		where_clause += db_socket.gen_where_clause_string(SQL_WHERE_CLAUSE_EQ_STRING, COL_LORAWAN_USER_ALIAS,alias)
	}
	sql_search_sql := db_socket.fmt_where_clause_sql(SQL_USER_SEARCH, where_clause) 
	log.Println(sql_search_sql)
	rows, err := db_socket.db.QueryContext(ctx, sql_search_sql, limit, offset)
	if err != nil {
		c <- &base.DBResult{ 
			Err:base.NewErr(err, base.ERR_USER_QUERY_CODE, base.ERR_USER_QUERY_DESC),
		}

		return
	}
	defer rows.Close()

	var _id, _alias, _role string
	users := make([]*base.User, 0)
	for rows.Next() {
		if err := rows.Scan(&_id, &_alias, &_role); err != nil {
			if err != nil {
				c <- &base.DBResult{ 
					Err:base.NewErr(err, base.ERR_USER_QUERY_CODE, base.ERR_USER_QUERY_DESC),
				}

				return
			}
		}
		users = append(users, &base.User{
			ID:     _id,
			Alias:  _alias,
			Role:_role,
		})
	}
	if err = rows.Err(); err != nil {
		if err != nil {
			c <- &base.DBResult{ 
				Err:base.NewErr(err, base.ERR_USER_QUERY_CODE, base.ERR_USER_QUERY_DESC),
			}

			return
		}
	}

	c<-&base.DBResult{
		Extra:users,
	}
}

func (db_socket *DBSocket)UserSearchCount(ctx context.Context, id, alias string, c chan *base.DBResult){ 
	var where_clause string
	if id != ""{
		where_clause += db_socket.gen_where_clause_string(SQL_WHERE_CLAUSE_EQ_STRING, COL_LORAWAN_USER_ID, id)
	}

	if alias != ""{
		where_clause += db_socket.gen_where_clause_string(SQL_WHERE_CLAUSE_EQ_STRING, COL_LORAWAN_USER_ALIAS,alias)
	}
	sql_search_count_sql := db_socket.fmt_where_clause_sql(SQL_USER_SEARCH_COUNT, where_clause) 
	count := 0
	err := db_socket.db.QueryRowContext(ctx, sql_search_count_sql).Scan(&count)
	if err != nil{
		c <- &base.DBResult{
			Err:base.NewErr(err, base.ERR_USER_QUERY_CODE, base.ERR_USER_QUERY_DESC), 
		}
	}else{
		c <- &base.DBResult{
			Extra:count,
		}
	}
}