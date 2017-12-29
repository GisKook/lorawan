package db

import (
	"fmt"
)

const (
	TABLE_LORAWAN_USER string = "lorawan_user"
	COL_LORAWAN_USER_ID string = "lorawan_user.id"
	COL_LORAWAN_USER_ALIAS string = "lorawan_user.alias"
	SQL_WHERE_CLAUSE_EQ_STRING string = "and %s='%s'"
)

func (db_socket *DBSocket) gen_where_clause_string(op, col ,value string) string {
	return fmt.Sprintf(op, col, value)
}

func (db_socket *DBSocket) fmt_where_clause_sql(dst_sql, where_clause string) string{
	return fmt.Sprintf(dst_sql, where_clause)
}