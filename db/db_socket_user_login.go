package db

import (
	"crypto/md5"
	"fmt"
	"github.com/giskook/lorawan/base"
	"io"
)

const (
	SQL_USER_LOGIN string = "select lorawan_user.alias,lorawan_user.passwd, lorawan_role.name ,lorawan_entity.name, lorawan_entity.alias from lorawan_user, lorawan_role, lorawan_entity, lorawan_user_role, lorawan_role_entity where lorawan_user.id=lorawan_user_role.user_id and lorawan_role.id=lorawan_user_role.role_id and lorawan_role.id=lorawan_role_entity.role_id and lorawan_role_entity.entity_id=lorawan_entity.id and lorawan_user.id=$1"
)

func (db_socket *DBSocket) gen_passwd(password , secret string) string{
	m := md5.New()
	io.WriteString(m, password)

	return fmt.Sprintf("%x", m.Sum([]byte(secret)))
}

func (db_socket *DBSocket) valid_passwd(passwd_input string, secret string, passwd_store string) bool {
	if db_socket.gen_passwd(passwd_input, secret) == passwd_store {
		return true
	}

	return false
}

func (db_socket *DBSocket) UserValid(user string, password string) (string, string, error) {
	rows, err := db_socket.db.Query(SQL_USER_LOGIN, user)
	if err != nil {
		return "", "", base.NewErr(err, base.ERR_DB_QUERY_FAIL_CODE, base.ERR_DB_QUERY_FAIL_DESC)
	}
	defer rows.Close()

	var name, passwd, role_name, entity_name, entity_alias, auth string

	if rows.Next() {
		if err := rows.Scan(&name, &passwd, &role_name, &entity_name, &entity_alias); err != nil {
			return "", "", base.NewErr(err, base.ERR_DB_QUERY_FAIL_CODE, base.ERR_DB_QUERY_FAIL_DESC)
		}

		if db_socket.valid_passwd(password, db_socket.conf.Secret, passwd) {
			auth = entity_name
		} else {
			return "", "", base.NewErr(nil, base.ERR_USER_UNVALID_PASSWD_CODE, base.ERR_USER_UNVALID_PASSWD_DESC)
		}
	} else {
		return "", "", base.NewErr(nil, base.ERR_USER_NOT_FOUND_CODE, base.ERR_USER_NOT_FOUND_DESC)
	}

	for rows.Next() {
		if err := rows.Scan(&name, &passwd, &role_name, &entity_name, &entity_alias); err != nil {
			return "", "", base.NewErr(err, base.ERR_DB_QUERY_FAIL_CODE, base.ERR_DB_QUERY_FAIL_DESC)
		}
		auth += "-" + entity_name
	}

	if err := rows.Err(); err != nil {
		return "", "", base.NewErr(err, base.ERR_DB_QUERY_FAIL_CODE, base.ERR_DB_QUERY_FAIL_DESC)
	}

	return auth, user, nil

}
